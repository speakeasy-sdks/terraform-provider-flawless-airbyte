// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/sdk/pkg/models/shared"
	"airbyte/internal/sdk/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// sourceDefinitionSpecification - SourceDefinition specification related resources.
type sourceDefinitionSpecification struct {
	sdkConfiguration sdkConfiguration
}

func newSourceDefinitionSpecification(sdkConfig sdkConfiguration) *sourceDefinitionSpecification {
	return &sourceDefinitionSpecification{
		sdkConfiguration: sdkConfig,
	}
}

// GetSourceDefinitionSpecification - Get specification for a SourceDefinition.
func (s *sourceDefinitionSpecification) GetSourceDefinitionSpecification(ctx context.Context, request shared.SourceDefinitionIDWithWorkspaceID) (*operations.GetSourceDefinitionSpecificationResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/source_definition_specifications/get"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSourceDefinitionSpecificationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SourceDefinitionSpecificationRead
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.SourceDefinitionSpecificationRead = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.NotFoundKnownExceptionInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.NotFoundKnownExceptionInfo = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InvalidInputExceptionInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.InvalidInputExceptionInfo = out
		}
	}

	return res, nil
}

// GetSpecificationForSourceID - Get specification for a source.
func (s *sourceDefinitionSpecification) GetSpecificationForSourceID(ctx context.Context, request shared.SourceIDRequestBody) (*operations.GetSpecificationForSourceIDResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/source_definition_specifications/get_for_source"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetSpecificationForSourceIDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.SourceDefinitionSpecificationRead
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.SourceDefinitionSpecificationRead = out
		}
	case httpRes.StatusCode == 404:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.NotFoundKnownExceptionInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.NotFoundKnownExceptionInfo = out
		}
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.InvalidInputExceptionInfo
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.InvalidInputExceptionInfo = out
		}
	}

	return res, nil
}