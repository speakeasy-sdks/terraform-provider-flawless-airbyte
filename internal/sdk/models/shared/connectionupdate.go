// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/internal/utils"
)

// ConnectionUpdate - Used to apply a patch-style update to a connection, which means that null properties remain unchanged
type ConnectionUpdate struct {
	BreakingChange *bool      `json:"breakingChange,omitempty"`
	ConnectionID   string     `json:"connectionId"`
	Geography      *Geography `json:"geography,omitempty"`
	// Name that will be set to this connection
	Name *string `json:"name,omitempty"`
	// Method used for computing final namespace in destination
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat              *string                       `default:"null" json:"namespaceFormat"`
	NonBreakingChangesPreference *NonBreakingChangesPreference `json:"nonBreakingChangesPreference,omitempty"`
	NotifySchemaChanges          *bool                         `json:"notifySchemaChanges,omitempty"`
	NotifySchemaChangesByEmail   *bool                         `json:"notifySchemaChangesByEmail,omitempty"`
	OperationIds                 []string                      `json:"operationIds,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	// optional resource requirements to run workers (blank for unbounded allocations)
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
	// if null, then no schedule is set.
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	ScheduleData *ConnectionScheduleData `json:"scheduleData,omitempty"`
	// determine how the schedule data should be interpreted
	ScheduleType    *ConnectionScheduleType `json:"scheduleType,omitempty"`
	SourceCatalogID *string                 `json:"sourceCatalogId,omitempty"`
	// Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
	Status *ConnectionStatus `json:"status,omitempty"`
	// describes the available schema (catalog).
	SyncCatalog *AirbyteCatalog `json:"syncCatalog,omitempty"`
}

func (c ConnectionUpdate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionUpdate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionUpdate) GetBreakingChange() *bool {
	if o == nil {
		return nil
	}
	return o.BreakingChange
}

func (o *ConnectionUpdate) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *ConnectionUpdate) GetGeography() *Geography {
	if o == nil {
		return nil
	}
	return o.Geography
}

func (o *ConnectionUpdate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionUpdate) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionUpdate) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionUpdate) GetNonBreakingChangesPreference() *NonBreakingChangesPreference {
	if o == nil {
		return nil
	}
	return o.NonBreakingChangesPreference
}

func (o *ConnectionUpdate) GetNotifySchemaChanges() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChanges
}

func (o *ConnectionUpdate) GetNotifySchemaChangesByEmail() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChangesByEmail
}

func (o *ConnectionUpdate) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *ConnectionUpdate) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionUpdate) GetResourceRequirements() *ResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *ConnectionUpdate) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionUpdate) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *ConnectionUpdate) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *ConnectionUpdate) GetSourceCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.SourceCatalogID
}

func (o *ConnectionUpdate) GetStatus() *ConnectionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ConnectionUpdate) GetSyncCatalog() *AirbyteCatalog {
	if o == nil {
		return nil
	}
	return o.SyncCatalog
}