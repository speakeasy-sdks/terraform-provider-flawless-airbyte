// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"airbyte/internal/sdk/pkg/models/operations"
	"airbyte/internal/sdk/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type openapi struct {
	sdkConfiguration sdkConfiguration
}

func newOpenapi(sdkConfig sdkConfiguration) *openapi {
	return &openapi{
		sdkConfiguration: sdkConfig,
	}
}

// GetOpenAPISpec - Returns the openapi specification
func (s *openapi) GetOpenAPISpec(ctx context.Context) (*operations.GetOpenAPISpecResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/openapi"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetOpenAPISpecResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			res.GetOpenAPISpec200TextPlainBinaryString = rawBody
		}
	}

	return res, nil
}
