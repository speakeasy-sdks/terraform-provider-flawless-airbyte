// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceOauthConsentRequest struct {
	// The values required to configure OAuth flows. The schema for this must match the `OAuthConfigSpecification.oauthUserInputFromConnectorConfigSpecification` schema.
	OAuthInputConfiguration interface{} `json:"oAuthInputConfiguration,omitempty"`
	// The url to redirect to after getting the user consent
	RedirectURL        string  `json:"redirectUrl"`
	SourceDefinitionID string  `json:"sourceDefinitionId"`
	SourceID           *string `json:"sourceId,omitempty"`
	WorkspaceID        string  `json:"workspaceId"`
}

func (o *SourceOauthConsentRequest) GetOAuthInputConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.OAuthInputConfiguration
}

func (o *SourceOauthConsentRequest) GetRedirectURL() string {
	if o == nil {
		return ""
	}
	return o.RedirectURL
}

func (o *SourceOauthConsentRequest) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}

func (o *SourceOauthConsentRequest) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

func (o *SourceOauthConsentRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}