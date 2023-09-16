// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type JobConfigType string

const (
	JobConfigTypeCheckConnectionSource      JobConfigType = "check_connection_source"
	JobConfigTypeCheckConnectionDestination JobConfigType = "check_connection_destination"
	JobConfigTypeDiscoverSchema             JobConfigType = "discover_schema"
	JobConfigTypeGetSpec                    JobConfigType = "get_spec"
	JobConfigTypeSync                       JobConfigType = "sync"
	JobConfigTypeResetConnection            JobConfigType = "reset_connection"
)

func (e JobConfigType) ToPointer() *JobConfigType {
	return &e
}

func (e *JobConfigType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "check_connection_source":
		fallthrough
	case "check_connection_destination":
		fallthrough
	case "discover_schema":
		fallthrough
	case "get_spec":
		fallthrough
	case "sync":
		fallthrough
	case "reset_connection":
		*e = JobConfigType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JobConfigType: %v", v)
	}
}