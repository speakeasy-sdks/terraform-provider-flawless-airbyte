// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AuthFlowType string

const (
	AuthFlowTypeOauth20 AuthFlowType = "oauth2.0"
	AuthFlowTypeOauth10 AuthFlowType = "oauth1.0"
)

func (e AuthFlowType) ToPointer() *AuthFlowType {
	return &e
}

func (e *AuthFlowType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		fallthrough
	case "oauth1.0":
		*e = AuthFlowType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AuthFlowType: %v", v)
	}
}

type AdvancedAuth struct {
	AuthFlowType             *AuthFlowType             `json:"authFlowType,omitempty"`
	OauthConfigSpecification *OAuthConfigSpecification `json:"oauthConfigSpecification,omitempty"`
	// Json Path to a field in the connectorSpecification that should exist for the advanced auth to be applicable.
	PredicateKey []string `json:"predicateKey,omitempty"`
	// Value of the predicate_key fields for the advanced auth to be applicable.
	PredicateValue *string `json:"predicateValue,omitempty"`
}

func (o *AdvancedAuth) GetAuthFlowType() *AuthFlowType {
	if o == nil {
		return nil
	}
	return o.AuthFlowType
}

func (o *AdvancedAuth) GetOauthConfigSpecification() *OAuthConfigSpecification {
	if o == nil {
		return nil
	}
	return o.OauthConfigSpecification
}

func (o *AdvancedAuth) GetPredicateKey() []string {
	if o == nil {
		return nil
	}
	return o.PredicateKey
}

func (o *AdvancedAuth) GetPredicateValue() *string {
	if o == nil {
		return nil
	}
	return o.PredicateValue
}