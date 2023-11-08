// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type WebBackendUpdateConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful operation
	WebBackendConnectionRead *shared.WebBackendConnectionRead
}

func (o *WebBackendUpdateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *WebBackendUpdateConnectionResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}

func (o *WebBackendUpdateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *WebBackendUpdateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *WebBackendUpdateConnectionResponse) GetWebBackendConnectionRead() *shared.WebBackendConnectionRead {
	if o == nil {
		return nil
	}
	return o.WebBackendConnectionRead
}
