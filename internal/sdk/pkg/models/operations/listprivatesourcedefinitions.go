// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListPrivateSourceDefinitionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	PrivateSourceDefinitionReadList *shared.PrivateSourceDefinitionReadList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
