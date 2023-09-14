// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SchemaChange string

const (
	SchemaChangeNoChange    SchemaChange = "no_change"
	SchemaChangeNonBreaking SchemaChange = "non_breaking"
	SchemaChangeBreaking    SchemaChange = "breaking"
)

func (e SchemaChange) ToPointer() *SchemaChange {
	return &e
}

func (e *SchemaChange) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "no_change":
		fallthrough
	case "non_breaking":
		fallthrough
	case "breaking":
		*e = SchemaChange(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemaChange: %v", v)
	}
}
