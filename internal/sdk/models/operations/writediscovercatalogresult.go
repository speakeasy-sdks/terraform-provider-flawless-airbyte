// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type WriteDiscoverCatalogResultResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful Operation
	DiscoverCatalogResult *shared.DiscoverCatalogResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *WriteDiscoverCatalogResultResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WriteDiscoverCatalogResultResponse) GetDiscoverCatalogResult() *shared.DiscoverCatalogResult {
	if o == nil {
		return nil
	}
	return o.DiscoverCatalogResult
}

func (o *WriteDiscoverCatalogResultResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WriteDiscoverCatalogResultResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}