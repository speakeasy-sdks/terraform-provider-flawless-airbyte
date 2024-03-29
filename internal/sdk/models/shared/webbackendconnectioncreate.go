// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/internal/utils"
)

type WebBackendConnectionCreate struct {
	DestinationID string     `json:"destinationId"`
	Geography     *Geography `json:"geography,omitempty"`
	// Optional name of the connection
	Name *string `json:"name,omitempty"`
	// Method used for computing final namespace in destination
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat              *string                       `default:"null" json:"namespaceFormat"`
	NonBreakingChangesPreference *NonBreakingChangesPreference `json:"nonBreakingChangesPreference,omitempty"`
	OperationIds                 []string                      `json:"operationIds,omitempty"`
	Operations                   []OperationCreate             `json:"operations,omitempty"`
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

func (w WebBackendConnectionCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebBackendConnectionCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebBackendConnectionCreate) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *WebBackendConnectionCreate) GetGeography() *Geography {
	if o == nil {
		return nil
	}
	return o.Geography
}

func (o *WebBackendConnectionCreate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *WebBackendConnectionCreate) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *WebBackendConnectionCreate) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *WebBackendConnectionCreate) GetNonBreakingChangesPreference() *NonBreakingChangesPreference {
	if o == nil {
		return nil
	}
	return o.NonBreakingChangesPreference
}

func (o *WebBackendConnectionCreate) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *WebBackendConnectionCreate) GetOperations() []OperationCreate {
	if o == nil {
		return nil
	}
	return o.Operations
}

func (o *WebBackendConnectionCreate) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *WebBackendConnectionCreate) GetResourceRequirements() *ResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *WebBackendConnectionCreate) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *WebBackendConnectionCreate) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *WebBackendConnectionCreate) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *WebBackendConnectionCreate) GetSourceCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.SourceCatalogID
}

func (o *WebBackendConnectionCreate) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *WebBackendConnectionCreate) GetStatus() ConnectionStatus {
	if o == nil {
		return ConnectionStatus("")
	}
	return o.Status
}

func (o *WebBackendConnectionCreate) GetSyncCatalog() *AirbyteCatalog {
	if o == nil {
		return nil
	}
	return o.SyncCatalog
}
