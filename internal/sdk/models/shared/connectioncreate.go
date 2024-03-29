// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type ConnectionCreate struct {
	DestinationID string     `json:"destinationId"`
	Geography     *Geography `json:"geography,omitempty"`
	// Optional name of the connection
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
	SourceID        string                  `json:"sourceId"`
	// Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
	Status ConnectionStatus `json:"status"`
	// describes the available schema (catalog).
	SyncCatalog *AirbyteCatalog `json:"syncCatalog,omitempty"`
}

func (c ConnectionCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConnectionCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionCreate) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *ConnectionCreate) GetGeography() *Geography {
	if o == nil {
		return nil
	}
	return o.Geography
}

func (o *ConnectionCreate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ConnectionCreate) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *ConnectionCreate) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *ConnectionCreate) GetNonBreakingChangesPreference() *NonBreakingChangesPreference {
	if o == nil {
		return nil
	}
	return o.NonBreakingChangesPreference
}

func (o *ConnectionCreate) GetNotifySchemaChanges() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChanges
}

func (o *ConnectionCreate) GetNotifySchemaChangesByEmail() *bool {
	if o == nil {
		return nil
	}
	return o.NotifySchemaChangesByEmail
}

func (o *ConnectionCreate) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *ConnectionCreate) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *ConnectionCreate) GetResourceRequirements() *ResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *ConnectionCreate) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *ConnectionCreate) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *ConnectionCreate) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *ConnectionCreate) GetSourceCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.SourceCatalogID
}

func (o *ConnectionCreate) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *ConnectionCreate) GetStatus() ConnectionStatus {
	if o == nil {
		return ConnectionStatus("")
	}
	return o.Status
}

func (o *ConnectionCreate) GetSyncCatalog() *AirbyteCatalog {
	if o == nil {
		return nil
	}
	return o.SyncCatalog
}
