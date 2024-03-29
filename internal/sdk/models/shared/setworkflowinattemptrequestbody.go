// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type SetWorkflowInAttemptRequestBody struct {
	AttemptNumber       int     `json:"attemptNumber"`
	JobID               int64   `json:"jobId"`
	ProcessingTaskQueue *string `default:"" json:"processingTaskQueue"`
	WorkflowID          string  `json:"workflowId"`
}

func (s SetWorkflowInAttemptRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetWorkflowInAttemptRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetWorkflowInAttemptRequestBody) GetAttemptNumber() int {
	if o == nil {
		return 0
	}
	return o.AttemptNumber
}

func (o *SetWorkflowInAttemptRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *SetWorkflowInAttemptRequestBody) GetProcessingTaskQueue() *string {
	if o == nil {
		return nil
	}
	return o.ProcessingTaskQueue
}

func (o *SetWorkflowInAttemptRequestBody) GetWorkflowID() string {
	if o == nil {
		return ""
	}
	return o.WorkflowID
}
