// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/models/shared"
	"net/http"
)

type GetSourceOAuthConsentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Input failed validation
	InvalidInputExceptionInfo *shared.InvalidInputExceptionInfo
	// Object with given id was not found.
	NotFoundKnownExceptionInfo *shared.NotFoundKnownExceptionInfo
	// Successful operation
	OAuthConsentRead *shared.OAuthConsentRead
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSourceOAuthConsentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourceOAuthConsentResponse) GetInvalidInputExceptionInfo() *shared.InvalidInputExceptionInfo {
	if o == nil {
		return nil
	}
	return o.InvalidInputExceptionInfo
}

func (o *GetSourceOAuthConsentResponse) GetNotFoundKnownExceptionInfo() *shared.NotFoundKnownExceptionInfo {
	if o == nil {
		return nil
	}
	return o.NotFoundKnownExceptionInfo
}

func (o *GetSourceOAuthConsentResponse) GetOAuthConsentRead() *shared.OAuthConsentRead {
	if o == nil {
		return nil
	}
	return o.OAuthConsentRead
}

func (o *GetSourceOAuthConsentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourceOAuthConsentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}