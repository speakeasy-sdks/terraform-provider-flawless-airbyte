// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type UpdateConnectionResponse struct {
	// Successful operation
	ConnectionRead *shared.ConnectionRead
	// HTTP response content type for this operation
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateConnectionResponse) GetConnectionRead() *shared.ConnectionRead {
	if o == nil {
		return nil
	}
	return o.ConnectionRead
}

func (o *UpdateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConnectionResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}

func (o *UpdateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
