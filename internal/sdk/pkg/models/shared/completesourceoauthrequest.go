// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CompleteSourceOauthRequest struct {
	// The values required to configure OAuth flows. The schema for this must match the `OAuthConfigSpecification.oauthUserInputFromConnectorConfigSpecification` schema.
	OAuthInputConfiguration interface{} `json:"oAuthInputConfiguration,omitempty"`
	// The query parameters present in the redirect URL after a user granted consent e.g auth code
	QueryParams map[string]interface{} `json:"queryParams,omitempty"`
	// When completing OAuth flow to gain an access token, some API sometimes requires to verify that the app re-send the redirectUrl that was used when consent was given.
	RedirectURL *string `json:"redirectUrl,omitempty"`
	// If set to true, returns a secret coordinate which references the stored tokens. By default, returns raw tokens.
	ReturnSecretCoordinate *bool   `json:"returnSecretCoordinate,omitempty"`
	SourceDefinitionID     string  `json:"sourceDefinitionId"`
	SourceID               *string `json:"sourceId,omitempty"`
	WorkspaceID            string  `json:"workspaceId"`
}
