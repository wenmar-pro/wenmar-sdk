package wenmar

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DownloadResult holds a streaming response body for binary download
// operations (exports, files). The caller must close Body when done.
type DownloadResult struct {
	Body          io.ReadCloser
	ContentType   string
	Filename      string
	ContentLength int64
	HTTPResponse  *http.Response
}

// closeOnError closes the body if err is non-nil, preventing leaks on
// error paths.
func (d *DownloadResult) closeOnError(err error) {
	if err != nil && d.Body != nil {
		d.Body.Close()
	}
}

// parseFilename extracts the filename from a Content-Disposition header.
// Returns "" if not present.
func parseFilename(contentDisposition string) string {
	// format: attachment; filename="customers.csv"
	parts := strings.Split(contentDisposition, ";")
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if strings.HasPrefix(p, "filename=") {
			val := strings.TrimPrefix(p, "filename=")
			val = strings.Trim(val, "\"")
			return val
		}
	}
	return ""
}

// downloadURL performs a GET and returns a streaming DownloadResult.
// The caller must close result.Body.
func (c *Client) downloadURL(ctx context.Context, url string) (*DownloadResult, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if err := c.requestEditor(ctx, req); err != nil {
		return nil, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, ParseErrorBodyWithRequest(body, resp.StatusCode, "GET", url)
	}

	result := &DownloadResult{
		Body:          resp.Body,
		ContentType:   resp.Header.Get("Content-Type"),
		Filename:      parseFilename(resp.Header.Get("Content-Disposition")),
		ContentLength: resp.ContentLength,
		HTTPResponse:  resp,
	}
	return result, nil
}

// downloadPath performs a GET against a path (relative to BaseURL) and
// returns a streaming DownloadResult.
func (c *Client) downloadPath(ctx context.Context, path string) (*DownloadResult, error) {
	url := c.BaseURL + path
	return c.downloadURL(ctx, url)
}

// downloadPathWithQuery performs a GET against a path with query params
// and returns a streaming DownloadResult.
func (c *Client) downloadPathWithQuery(ctx context.Context, path string, q url.Values) (*DownloadResult, error) {
	fullURL := c.BaseURL + path
	encoded := q.Encode()
	if encoded != "" {
		fullURL += "?" + encoded
	}
	return c.downloadURL(ctx, fullURL)
}

// downloadPathWithParams performs a GET against a path, substituting path
// parameters, and returns a streaming DownloadResult.
func (c *Client) downloadPathWithParams(ctx context.Context, path string, pathArgs ...any) (*DownloadResult, error) {
	for _, arg := range pathArgs {
		idx := strings.Index(path, "{")
		if idx == -1 {
			break
		}
		end := strings.Index(path[idx:], "}")
		if end == -1 {
			break
		}
		replacement := fmt.Sprintf("%v", arg)
		path = path[:idx] + replacement + path[idx+end+1:]
	}
	return c.downloadPath(ctx, path)
}
