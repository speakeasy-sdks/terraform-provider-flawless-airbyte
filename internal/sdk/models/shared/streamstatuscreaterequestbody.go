// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StreamStatusCreateRequestBody struct {
	AttemptNumber int    `json:"attemptNumber"`
	ConnectionID  string `json:"connectionId"`
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

func (o *StreamStatusCreateRequestBody) GetAttemptNumber() int {
	if o == nil {
		return 0
	}
	return o.AttemptNumber
}

func (o *StreamStatusCreateRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *StreamStatusCreateRequestBody) GetIncompleteRunCause() *StreamStatusIncompleteRunCause {
	if o == nil {
		return nil
	}
	return o.IncompleteRunCause
}

func (o *StreamStatusCreateRequestBody) GetJobID() int64 {
	if o == nil {
		return 0
	}
	return o.JobID
}

func (o *StreamStatusCreateRequestBody) GetJobType() StreamStatusJobType {
	if o == nil {
		return StreamStatusJobType("")
	}
	return o.JobType
}

func (o *StreamStatusCreateRequestBody) GetRunState() StreamStatusRunState {
	if o == nil {
		return StreamStatusRunState("")
	}
	return o.RunState
}

func (o *StreamStatusCreateRequestBody) GetStreamName() string {
	if o == nil {
		return ""
	}
	return o.StreamName
}

func (o *StreamStatusCreateRequestBody) GetStreamNamespace() *string {
	if o == nil {
		return nil
	}
	return o.StreamNamespace
}

func (o *StreamStatusCreateRequestBody) GetTransitionedAt() int64 {
	if o == nil {
		return 0
	}
	return o.TransitionedAt
}

func (o *StreamStatusCreateRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}