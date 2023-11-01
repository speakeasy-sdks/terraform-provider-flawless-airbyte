// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CheckConnectionReadStatus string

const (
	CheckConnectionReadStatusSucceeded CheckConnectionReadStatus = "succeeded"
	CheckConnectionReadStatusFailed    CheckConnectionReadStatus = "failed"
)

func (e CheckConnectionReadStatus) ToPointer() *CheckConnectionReadStatus {
	return &e
}

func (e *CheckConnectionReadStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "succeeded":
		fallthrough
	case "failed":
		*e = CheckConnectionReadStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CheckConnectionReadStatus: %v", v)
	}
}

type CheckConnectionRead struct {
	JobInfo SynchronousJobRead         `json:"jobInfo"`
	Message *string                    `json:"message,omitempty"`
	Status  *CheckConnectionReadStatus `json:"status,omitempty"`
}

func (o *CheckConnectionRead) GetJobInfo() SynchronousJobRead {
	if o == nil {
		return SynchronousJobRead{}
	}
	return o.JobInfo
}

func (o *CheckConnectionRead) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CheckConnectionRead) GetStatus() *CheckConnectionReadStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
