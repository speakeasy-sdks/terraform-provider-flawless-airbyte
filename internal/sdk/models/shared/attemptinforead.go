// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AttemptInfoRead struct {
	Attempt AttemptRead `json:"attempt"`
	Logs    LogRead     `json:"logs"`
}

func (o *AttemptInfoRead) GetAttempt() AttemptRead {
	if o == nil {
		return AttemptRead{}
	}
	return o.Attempt
}

func (o *AttemptInfoRead) GetLogs() LogRead {
	if o == nil {
		return LogRead{}
	}
	return o.Logs
}