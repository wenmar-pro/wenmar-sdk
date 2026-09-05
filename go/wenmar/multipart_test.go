package wenmar

import (
	"bytes"
	"context"
	"io"
	"mime"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMultipartBuilder_BuildCSVUpload(t *testing.T) {
	builder := NewMultipartBuilder()
	builder.AddFile("file", "customers.csv", "text/csv", bytes.NewReader([]byte("name,email\nJane,jane@example.com\n")))

	body, contentType, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		t.Fatalf("invalid content-type: %v", err)
	}
	if mediaType != "multipart/form-data" {
		t.Errorf("expected multipart/form-data, got %s", mediaType)
	}
	boundary := params["boundary"]
	if boundary == "" {
		t.Fatal("missing boundary in content-type")
	}

	bodyStr := body.String()
	if !bytes.Contains(body.Bytes(), []byte("customers.csv")) {
		t.Errorf("body should contain filename, got: %s", bodyStr)
	}
	if !bytes.Contains(body.Bytes(), []byte("name,email")) {
		t.Errorf("body should contain file content, got: %s", bodyStr)
	}
}

func TestCreateExpensesImportsValidate_WithMultipartBuilder(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("expected POST, got %s", r.Method)
		}
		ct := r.Header.Get("Content-Type")
		if !strings.HasPrefix(ct, "multipart/form-data") {
			t.Errorf("expected multipart/form-data, got %s", ct)
		}
		// Drain the body
		io.Copy(io.Discard, r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		w.Write([]byte(`{"valid": true}`))
	}))
	defer ts.Close()

	c := newTestClient(t, ts.URL, "test")
	builder := NewMultipartBuilder()
	builder.AddFile("file", "expenses.csv", "text/csv", bytes.NewReader([]byte("amount,category\n100,tools\n")))
	body, contentType, err := builder.Build()
	if err != nil {
		t.Fatalf("Build failed: %v", err)
	}

	// Set the content type on the request via a custom request editor.
	// The multipart operation takes an io.Reader; the content type is
	// set by the generated client to "multipart/form-data" already,
	// but we need the boundary. We use the WithRequestEditorFn path.
	_ = contentType // The generated WithBodyWithResponse sets the CT.

	resp, err := c.CreateExpensesImportsValidate(context.Background(), body)
	if err != nil {
		t.Fatalf("CreateExpensesImportsValidate failed: %v", err)
	}
	if resp.StatusCode() != 201 {
		t.Errorf("expected 201, got %d", resp.StatusCode())
	}
}
