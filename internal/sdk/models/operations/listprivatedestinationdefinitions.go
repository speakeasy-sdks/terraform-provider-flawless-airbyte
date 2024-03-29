// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type ListPrivateDestinationDefinitionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	PrivateDestinationDefinitionReadList *shared.PrivateDestinationDefinitionReadList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListPrivateDestinationDefinitionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPrivateDestinationDefinitionsResponse) GetPrivateDestinationDefinitionReadList() *shared.PrivateDestinationDefinitionReadList {
	if o == nil {
		return nil
	}
	return o.PrivateDestinationDefinitionReadList
}

func (o *ListPrivateDestinationDefinitionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPrivateDestinationDefinitionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
