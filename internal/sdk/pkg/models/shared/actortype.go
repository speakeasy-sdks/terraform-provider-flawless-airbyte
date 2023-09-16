// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ActorType string

const (
	ActorTypeSource      ActorType = "source"
	ActorTypeDestination ActorType = "destination"
)

func (e ActorType) ToPointer() *ActorType {
	return &e
}

func (e *ActorType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "source":
		fallthrough
	case "destination":
		*e = ActorType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ActorType: %v", v)
	}
}