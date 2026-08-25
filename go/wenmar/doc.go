// Package wenmar provides a hand-written client layer over the generated
// oapi-codegen client. It adds retry, pagination, error mapping, and
// bearer-token auth.
//
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config ../oapi-codegen.yaml ../../spec/openapi.enriched.yaml
package wenmar
