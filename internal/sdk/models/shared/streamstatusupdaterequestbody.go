// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StreamStatusUpdateRequestBody struct {
	AttemptNumber int    `json:"attemptNumber"`
	ConnectionID  string `json:"connectionId"`
	ID            string `json:"id"`
	// Values:
	//   * `FAILED` - A failure has occurred
	//   * `CANCELED` - The run has been canceled
	//
	IncompleteRunCause *StreamStatusIncompleteRunCause `json:"incompleteRunCause,omitempty"`
	JobID              int64                           `json:"jobId"`
	JobType            StreamStatusJobType             `json:"jobType"`
	// Values:
	//   * `PENDING` - The stream operation has been selected to run
	//   * `RUNNING` - The stream operation is running
	//   * `COMPLETE` - The stream operation ran successfully to completion
	//   * `INCOMPLETE` - The stream operation has terminated in an incomplete state.
	//   See StreamStatusIncompleteRunCause for more details.
	//
	RunState        StreamStatusRunState `json:"runState"`
	StreamName      string               `json:"streamName"`
	StreamNamespace *string              `json:"streamNamespace,omitempty"`
	TransitionedAt  int64                `json:"transitionedAt"`
	WorkspaceID     string               `json:"workspaceId"`
}

func (o *StreamStatusUpdateRequestBody) GetAttemptNumber() int {
	if o == nil {
		return 0
	}
	return o.AttemptNumber
}

func (o *StreamStatusUpdateRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *StreamStatusUpdateRequestBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *StreamStatusUpdateRequestBody) GetIncompleteRunCause() *StreamStatusIncompleteRunCause {
	if o == nil {
		return nil
	}
	return o.IncompleteRunCause
}

func (o *StreamStatusUpdateRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *StreamStatusUpdateRequestBody) GetJobType() StreamStatusJobType {
	if o == nil {
		return StreamStatusJobType("")
	}
	return o.JobType
}

func (o *StreamStatusUpdateRequestBody) GetRunState() StreamStatusRunState {
	if o == nil {
		return StreamStatusRunState("")
	}
	return o.RunState
}

func (o *StreamStatusUpdateRequestBody) GetStreamName() string {
	if o == nil {
		return ""
	}
	return o.StreamName
}

func (o *StreamStatusUpdateRequestBody) GetStreamNamespace() *string {
	if o == nil {
		return nil
	}
	return o.StreamNamespace
}

func (o *StreamStatusUpdateRequestBody) GetTransitionedAt() int64 {
	if o == nil {
		return 0
	}
	return o.TransitionedAt
}

func (o *StreamStatusUpdateRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}