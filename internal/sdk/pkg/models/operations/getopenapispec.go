// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetOpenAPISpecResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Returns the openapi specification file
	GetOpenAPISpec200TextPlainBinaryString []byte
}

func (o *GetOpenAPISpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOpenAPISpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOpenAPISpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetOpenAPISpecResponse) GetGetOpenAPISpec200TextPlainBinaryString() []byte {
	if o == nil {
		return nil
	}
	return o.GetOpenAPISpec200TextPlainBinaryString
}
