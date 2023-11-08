// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListLatestDestinationDefinitionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	DestinationDefinitionReadList *shared.DestinationDefinitionReadList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListLatestDestinationDefinitionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListLatestDestinationDefinitionsResponse) GetDestinationDefinitionReadList() *shared.DestinationDefinitionReadList {
	if o == nil {
		return nil
	}
	return o.DestinationDefinitionReadList
}

func (o *ListLatestDestinationDefinitionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListLatestDestinationDefinitionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
