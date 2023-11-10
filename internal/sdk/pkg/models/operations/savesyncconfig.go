// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type SaveSyncConfigResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful Operation
	InternalOperationResult *shared.InternalOperationResult
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SaveSyncConfigResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SaveSyncConfigResponse) GetInternalOperationResult() *shared.InternalOperationResult {
	if o == nil {
		return nil
	}
	return o.InternalOperationResult
}

func (o *SaveSyncConfigResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SaveSyncConfigResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
