// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DbtCloud struct {
	// The account id associated with the job
	AccountID int64 `json:"accountId"`
	// The job id associated with the job
	JobID int64 `json:"jobId"`
}

func (o *DbtCloud) GetAccountID() int64 {
	if o == nil {
		return 0
	}
	return o.AccountID
}

func (o *DbtCloud) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

type WebhookType string

const (
	WebhookTypeDbtCloud WebhookType = "dbtCloud"
)

func (e WebhookType) ToPointer() *WebhookType {
	return &e
}

func (e *WebhookType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "dbtCloud":
		*e = WebhookType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebhookType: %v", v)
	}
}

type OperatorWebhook struct {
	DbtCloud *DbtCloud `json:"dbtCloud,omitempty"`
	// DEPRECATED. Populate dbtCloud instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ExecutionBody *string `json:"executionBody,omitempty"`
	// DEPRECATED. Populate dbtCloud instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	ExecutionURL *string `json:"executionUrl,omitempty"`
	// The id of the webhook configs to use from the workspace.
	WebhookConfigID *string      `json:"webhookConfigId,omitempty"`
	WebhookType     *WebhookType `json:"webhookType,omitempty"`
}

func (o *OperatorWebhook) GetDbtCloud() *DbtCloud {
	if o == nil {
		return nil
	}
	return o.DbtCloud
}

func (o *OperatorWebhook) GetExecutionBody() *string {
	if o == nil {
		return nil
	}
	return o.ExecutionBody
}

func (o *OperatorWebhook) GetExecutionURL() *string {
	if o == nil {
		return nil
	}
	return o.ExecutionURL
}

func (o *OperatorWebhook) GetWebhookConfigID() *string {
	if o == nil {
		return nil
	}
	return o.WebhookConfigID
}

func (o *OperatorWebhook) GetWebhookType() *WebhookType {
	if o == nil {
		return nil
	}
	return o.WebhookType
}