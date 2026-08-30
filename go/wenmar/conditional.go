package wenmar

import (
	"io"
	"net/http"
	"sync"
)

// cacheEntry holds the validators and body for a conditional-GET cache slot.
type cacheEntry struct {
	ETag         string
	LastModified string
	Body         []byte
}

// cachingTransport adds conditional-GET support (ETag / Last-Modified) on top
// of the underlying transport. On repeat GET requests it sends
// If-None-Match / If-Modified-Since and, on 304 Not Modified, returns the
// cached body so callers get data without re-downloading the full payload.
//
// Only GET requests are cached. Non-2xx responses are not cached.
type cachingTransport struct {
	transport http.RoundTripper
	mu        sync.Mutex
	cache     map[string]*cacheEntry
}

func newCachingTransport(transport http.RoundTripper) *cachingTransport {
	return &cachingTransport{
		transport: transport,
		cache:     make(map[string]*cacheEntry),
	}
}

// cacheKey derives a cache slot key from the request, including the
// X-Wenmar-Location header so two location-scoped clients never serve each
// other's cached bodies.
func cacheKey(req *http.Request) string {
	key := req.Method + " " + req.URL.String()
	if loc := req.Header.Get("X-Wenmar-Location"); loc != "" {
		key += " loc=" + loc
	}
	if accept := req.Header.Get("Accept"); accept != "" {
		key += " accept=" + accept
	}
	return key
}

func (t *cachingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	isGet := req.Method == http.MethodGet

	t.mu.Lock()
	entry := t.cache[cacheKey(req)]
	t.mu.Unlock()

	reqCopy := req
	if isGet && entry != nil {
		// Shallow-copy so we don't mutate the caller's request headers.
		r := *req
		reqCopy = &r
		reqCopy.Header = req.Header.Clone()
		if entry.ETag != "" {
			reqCopy.Header.Set("If-None-Match", entry.ETag)
		}
		if entry.LastModified != "" {
			reqCopy.Header.Set("If-Modified-Since", entry.LastModified)
		}
	}

	resp, err := t.transport.RoundTrip(reqCopy)
	if err != nil {
		return resp, err
	}

	if isGet && resp.StatusCode == http.StatusNotModified && entry != nil {
		resp.Body.Close()
		cachedResp := cloneResponseWithBody(entry.Body)
		return cachedResp, nil
	}

	if isGet && resp.StatusCode == http.StatusOK {
		etag := resp.Header.Get("ETag")
		if etag != "" {
			body := readResponseBody(resp)
			if body != nil {
				t.mu.Lock()
				t.cache[cacheKey(req)] = &cacheEntry{
					ETag:         etag,
					LastModified: resp.Header.Get("Last-Modified"),
					Body:         body,
				}
				t.mu.Unlock()
				// Restore a readable body so downstream parsers still work.
				resp.Body = &bodyReadCloser{data: body}
			}
		}
	}

	return resp, nil
}

func cloneResponseWithBody(body []byte) *http.Response {
	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Header:     headers,
		Body:       &bodyReadCloser{data: body},
		Request:    nil,
	}
}

func readResponseBody(resp *http.Response) []byte {
	body := make([]byte, 0, 1024)
	buf := make([]byte, 4096)
	for {
		n, err := resp.Body.Read(buf)
		body = append(body, buf[:n]...)
		if err != nil {
			break
		}
	}
	resp.Body.Close()
	if len(body) == 0 {
		return nil
	}
	return body
}

type bodyReadCloser struct {
	data []byte
	pos  int
}

func (b *bodyReadCloser) Read(p []byte) (int, error) {
	if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	n := copy(p, b.data[b.pos:])
	b.pos += n
	return n, nil
}

func (b *bodyReadCloser) Close() error { return nil }
