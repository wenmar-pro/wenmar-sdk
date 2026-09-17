package wenmar

import "net/http"

// hooksTransport wraps an http.RoundTripper to fire request-level
// observability hooks around each logical HTTP request. It sits OUTERMOST in
// the transport stack (wrapping retry/caching), so a single SDK operation
// yields exactly one OnRequestStart/OnRequestEnd pair, while internal retry
// attempts are reported separately via OnRetry from the retryTransport.
//
// Every client gets this wrapper — including caller-supplied
// cfg.HTTPClient values, which NewClient shallow-copies before wrapping so
// the caller's client is never mutated. Retry/caching transports exist only
// in the SDK-built stack; custom clients keep their own transport below the
// hooks wrapper.
type hooksTransport struct {
	transport http.RoundTripper
	hooks     Hooks
}

func (t *hooksTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	info := RequestInfo{Method: req.Method, URL: req.URL.String()}
	t.hooks.OnRequestStart(req.Context(), info)
	resp, err := t.transport.RoundTrip(req)
	result := RequestResult{Method: info.Method, URL: info.URL, Err: err}
	if resp != nil {
		result.StatusCode = resp.StatusCode
	}
	t.hooks.OnRequestEnd(req.Context(), info, result)
	return resp, err
}
