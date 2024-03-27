// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type WebBackendCheckUpdatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	WebBackendCheckUpdatesRead *shared.WebBackendCheckUpdatesRead
}

func (o *WebBackendCheckUpdatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebBackendCheckUpdatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebBackendCheckUpdatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebBackendCheckUpdatesResponse) GetWebBackendCheckUpdatesRead() *shared.WebBackendCheckUpdatesRead {
	if o == nil {
		return nil
	}
	return o.WebBackendCheckUpdatesRead
}