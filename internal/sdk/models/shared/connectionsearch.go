// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type ConnectionSearch struct {
	ConnectionID  *string            `json:"connectionId,omitempty"`
	Destination   *DestinationSearch `json:"destination,omitempty"`
	DestinationID *string            `json:"destinationId,omitempty"`
	Name          *string            `json:"name,omitempty"`
	// Method used for computing final namespace in destination
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat *string `default:"null" json:"namespaceFormat"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	// if null, then no schedule is set.
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	ScheduleData *ConnectionScheduleData `json:"scheduleData,omitempty"`
	// determine how the schedule data should be interpreted
	ScheduleType *ConnectionScheduleType `json:"scheduleType,omitempty"`
	Source       *SourceSearch           `json:"source,omitempty"`
	SourceID     *string                 `json:"sourceId,omitempty"`
	// Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
	Status *ConnectionStatus `json:"status,omitempty"`
}

func (c ConnectionSearch) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionSearch) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionSearch) GetConnectionID() *string {
	if o == nil {
		return nil
	}
	return o.ConnectionID
}

func (o *ConnectionSearch) GetDestination() *DestinationSearch {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *ConnectionSearch) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *ConnectionSearch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionSearch) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionSearch) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionSearch) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionSearch) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionSearch) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *ConnectionSearch) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *ConnectionSearch) GetSource() *SourceSearch {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *ConnectionSearch) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

func (o *ConnectionSearch) GetStatus() *ConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}
