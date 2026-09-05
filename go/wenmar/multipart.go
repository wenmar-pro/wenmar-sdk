package wenmar

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
)

// MultipartBuilder constructs a multipart/form-data body for file upload
// operations. It is the recommended way to build request bodies for
// CreateExpensesImportsValidate and CreateTiresImportsValidate.
type MultipartBuilder struct {
	fields map[string]multipartField
}

type multipartField struct {
	filename    string
	contentType string
	reader      io.Reader
}

// NewMultipartBuilder creates a new MultipartBuilder.
func NewMultipartBuilder() *MultipartBuilder {
	return &MultipartBuilder{fields: make(map[string]multipartField)}
}

// AddFile adds a file field to the multipart body.
// fieldName is the form field name (e.g. "file").
// filename is the name sent to the server (e.g. "import.csv").
// contentType is the MIME type of the file (e.g. "text/csv").
// reader is the file content.
func (b *MultipartBuilder) AddFile(fieldName, filename, contentType string, reader io.Reader) *MultipartBuilder {
	b.fields[fieldName] = multipartField{
		filename:    filename,
		contentType: contentType,
		reader:      reader,
	}
	return b
}

// Build returns the multipart body as a reader and the Content-Type header
// value (including the boundary). Pass these to the multipart upload
// operations.
func (b *MultipartBuilder) Build() (*bytes.Buffer, string, error) {
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	for fieldName, field := range b.fields {
		header := make(textproto.MIMEHeader)
		header.Set("Content-Disposition",
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`, fieldName, field.filename))
		if field.contentType != "" {
			header.Set("Content-Type", field.contentType)
		}
		part, err := writer.CreatePart(header)
		if err != nil {
			return nil, "", fmt.Errorf("multipart: creating part for %s: %w", fieldName, err)
		}
		if _, err := io.Copy(part, field.reader); err != nil {
			return nil, "", fmt.Errorf("multipart: copying content for %s: %w", fieldName, err)
		}
	}

	if err := writer.Close(); err != nil {
		return nil, "", fmt.Errorf("multipart: closing writer: %w", err)
	}

	return &buf, writer.FormDataContentType(), nil
}
