// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLatestSourceDefinitionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	SourceDefinitionReadList *shared.SourceDefinitionReadList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLatestSourceDefinitionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLatestSourceDefinitionsResponse) GetSourceDefinitionReadList() *shared.SourceDefinitionReadList {
	if o == nil {
		return nil
	}
	return o.SourceDefinitionReadList
}

func (o *ListLatestSourceDefinitionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLatestSourceDefinitionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
