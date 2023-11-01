// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetHealthCheckResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful operation
	HealthCheckRead *shared.HealthCheckRead
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetHealthCheckResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetHealthCheckResponse) GetHealthCheckRead() *shared.HealthCheckRead {
	if o == nil {
		return nil
	}
	return o.HealthCheckRead
}

func (o *GetHealthCheckResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetHealthCheckResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
