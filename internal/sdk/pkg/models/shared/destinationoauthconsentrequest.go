// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationOauthConsentRequest struct {
	DestinationDefinitionID string  `json:"destinationDefinitionId"`
	DestinationID           *string `json:"destinationId,omitempty"`
	// The values required to configure OAuth flows. The schema for this must match the `OAuthConfigSpecification.oauthUserInputFromConnectorConfigSpecification` schema.
	OAuthInputConfiguration interface{} `json:"oAuthInputConfiguration,omitempty"`
	// The url to redirect to after getting the user consent
	RedirectURL string `json:"redirectUrl"`
	WorkspaceID string `json:"workspaceId"`
}

func (o *DestinationOauthConsentRequest) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationOauthConsentRequest) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *DestinationOauthConsentRequest) GetOAuthInputConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.OAuthInputConfiguration
}

func (o *DestinationOauthConsentRequest) GetRedirectURL() string {
	if o == nil {
		return ""
	}
	return o.RedirectURL
}

func (o *DestinationOauthConsentRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
