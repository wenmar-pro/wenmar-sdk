package wenmar

import (
	"bytes"
	"mime"
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