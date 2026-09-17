package wenmar

import (
	"fmt"
	"io"
	"net/http"
	"sync"
)

// cacheEntry holds the validators, headers, and body for a conditional-GET
// cache slot. Header is the full response header set so a 304 can reproduce
// headers (e.g. Link, X-Total-Count, X-Per-Page) that pagination relies on.
type cacheEntry struct {
	ETag         string
	LastModified string
	Header       http.Header
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
	order     []string // insertion order, for FIFO eviction
}

func newCachingTransport(transport http.RoundTripper) *cachingTransport {
	return &cachingTransport{
		transport: transport,
		cache:     make(map[string]*cacheEntry),
		order:     make([]string, 0, maxCacheEntries),
	}
}

// maxCacheEntries bounds the conditional-GET cache so long-running
// processes (watch pollers, TUIs) cannot grow it without limit.
const maxCacheEntries = 128

// store writes a cache entry, evicting the oldest entries (FIFO) beyond
// maxCacheEntries. Re-writing an existing key keeps its position.
func (t *cachingTransport) store(key string, entry *cacheEntry) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if _, exists := t.cache[key]; !exists {
		t.order = append(t.order, key)
	}
	t.cache[key] = entry
	for len(t.order) > maxCacheEntries {
		oldest := t.order[0]
		t.order = t.order[1:]
		delete(t.cache, oldest)
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
		cachedResp := cloneResponseWithBody(req, entry.Body, entry.Header)
		return cachedResp, nil
	}

	if isGet && resp.StatusCode == http.StatusOK {
		etag := resp.Header.Get("ETag")
		lastModified := resp.Header.Get("Last-Modified")
		// Cache any 200 that carries a validator (ETag or Last-Modified) so a
		// conditional revalidation can later return 304. A mid-body read
		// error aborts the request: partial bytes must never be cached
		// or served.
		if etag != "" || lastModified != "" {
			body, readErr := io.ReadAll(resp.Body)
			resp.Body.Close()
			if readErr != nil {
				return nil, fmt.Errorf("read response body: %w", readErr)
			}
			if len(body) > 0 {
				t.store(cacheKey(req), &cacheEntry{
					ETag:         etag,
					LastModified: lastModified,
					Header:       resp.Header.Clone(),
					Body:         body,
				})
			}
			// Restore a readable body so downstream parsers still work.
			resp.Body = &bodyReadCloser{data: body}
		}
	}

	return resp, nil
}

// cloneResponseWithBody builds a 200 response around the cached body, carrying
// the cached headers (so Link/X-Total-Count/X-Per-Page survive a 304). The
// headers are deep-cloned to avoid mutating the shared cache entry, and the
// original request is attached so resp.Request works for callers.
func cloneResponseWithBody(req *http.Request, body []byte, headers http.Header) *http.Response {
	cloned := make(http.Header)
	for k, vv := range headers {
		cloned[k] = append([]string(nil), vv...)
	}
	if cloned.Get("Content-Type") == "" {
		cloned.Set("Content-Type", "application/json")
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Status:     "200 OK",
		Header:     cloned,
		Body:       &bodyReadCloser{data: body},
		Request:    req,
	}
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
