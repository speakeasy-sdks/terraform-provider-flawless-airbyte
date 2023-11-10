// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/flawless/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type WebBackendConnectionRead struct {
	// Describes the difference between two Airbyte catalogs.
	CatalogDiff   *CatalogDiff    `json:"catalogDiff,omitempty"`
	CatalogID     *string         `json:"catalogId,omitempty"`
	ConnectionID  string          `json:"connectionId"`
	Destination   DestinationRead `json:"destination"`
	DestinationID string          `json:"destinationId"`
	Geography     *Geography      `json:"geography,omitempty"`
	IsSyncing     bool            `json:"isSyncing"`
	// epoch time of the latest sync job. null if no sync job has taken place.
	LatestSyncJobCreatedAt *int64     `json:"latestSyncJobCreatedAt,omitempty"`
	LatestSyncJobStatus    *JobStatus `json:"latestSyncJobStatus,omitempty"`
	Name                   string     `json:"name"`
	// Method used for computing final namespace in destination
	NamespaceDefinition *NamespaceDefinitionType `json:"namespaceDefinition,omitempty"`
	// Used when namespaceDefinition is 'customformat'. If blank then behaves like namespaceDefinition = 'destination'. If "${SOURCE_NAMESPACE}" then behaves like namespaceDefinition = 'source'.
	NamespaceFormat              *string                      `default:"null" json:"namespaceFormat"`
	NonBreakingChangesPreference NonBreakingChangesPreference `json:"nonBreakingChangesPreference"`
	NotifySchemaChanges          bool                         `json:"notifySchemaChanges"`
	NotifySchemaChangesByEmail   bool                         `json:"notifySchemaChangesByEmail"`
	OperationIds                 []string                     `json:"operationIds,omitempty"`
	Operations                   []OperationRead              `json:"operations,omitempty"`
	// Prefix that will be prepended to the name of each stream when it is written to the destination.
	Prefix *string `json:"prefix,omitempty"`
	// optional resource requirements to run workers (blank for unbounded allocations)
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
	// if null, then no schedule is set.
	Schedule *ConnectionSchedule `json:"schedule,omitempty"`
	// schedule for when the the connection should run, per the schedule type
	ScheduleData *ConnectionScheduleData `json:"scheduleData,omitempty"`
	// determine how the schedule data should be interpreted
	ScheduleType *ConnectionScheduleType `json:"scheduleType,omitempty"`
	SchemaChange SchemaChange            `json:"schemaChange"`
	Source       SourceRead              `json:"source"`
	SourceID     string                  `json:"sourceId"`
	// Active means that data is flowing through the connection. Inactive means it is not. Deprecated means the connection is off and cannot be re-activated. the schema field describes the elements of the schema that will be synced.
	Status ConnectionStatus `json:"status"`
	// describes the available schema (catalog).
	SyncCatalog AirbyteCatalog `json:"syncCatalog"`
}

func (w WebBackendConnectionRead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebBackendConnectionRead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebBackendConnectionRead) GetCatalogDiff() *CatalogDiff {
	if o == nil {
		return nil
	}
	return o.CatalogDiff
}

func (o *WebBackendConnectionRead) GetCatalogID() *string {
	if o == nil {
		return nil
	}
	return o.CatalogID
}

func (o *WebBackendConnectionRead) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

func (o *WebBackendConnectionRead) GetDestination() DestinationRead {
	if o == nil {
		return DestinationRead{}
	}
	return o.Destination
}

func (o *WebBackendConnectionRead) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *WebBackendConnectionRead) GetGeography() *Geography {
	if o == nil {
		return nil
	}
	return o.Geography
}

func (o *WebBackendConnectionRead) GetIsSyncing() bool {
	if o == nil {
		return false
	}
	return o.IsSyncing
}

func (o *WebBackendConnectionRead) GetLatestSyncJobCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.LatestSyncJobCreatedAt
}

func (o *WebBackendConnectionRead) GetLatestSyncJobStatus() *JobStatus {
	if o == nil {
		return nil
	}
	return o.LatestSyncJobStatus
}

func (o *WebBackendConnectionRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WebBackendConnectionRead) GetNamespaceDefinition() *NamespaceDefinitionType {
	if o == nil {
		return nil
	}
	return o.NamespaceDefinition
}

func (o *WebBackendConnectionRead) GetNamespaceFormat() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceFormat
}

func (o *WebBackendConnectionRead) GetNonBreakingChangesPreference() NonBreakingChangesPreference {
	if o == nil {
		return NonBreakingChangesPreference("")
	}
	return o.NonBreakingChangesPreference
}

func (o *WebBackendConnectionRead) GetNotifySchemaChanges() bool {
	if o == nil {
		return false
	}
	return o.NotifySchemaChanges
}

func (o *WebBackendConnectionRead) GetNotifySchemaChangesByEmail() bool {
	if o == nil {
		return false
	}
	return o.NotifySchemaChangesByEmail
}

func (o *WebBackendConnectionRead) GetOperationIds() []string {
	if o == nil {
		return nil
	}
	return o.OperationIds
}

func (o *WebBackendConnectionRead) GetOperations() []OperationRead {
	if o == nil {
		return nil
	}
	return o.Operations
}

func (o *WebBackendConnectionRead) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *WebBackendConnectionRead) GetResourceRequirements() *ResourceRequirements {
	if o == nil {
		return nil
	}
	return o.ResourceRequirements
}

func (o *WebBackendConnectionRead) GetSchedule() *ConnectionSchedule {
	if o == nil {
		return nil
	}
	return o.Schedule
}

func (o *WebBackendConnectionRead) GetScheduleData() *ConnectionScheduleData {
	if o == nil {
		return nil
	}
	return o.ScheduleData
}

func (o *WebBackendConnectionRead) GetScheduleType() *ConnectionScheduleType {
	if o == nil {
		return nil
	}
	return o.ScheduleType
}

func (o *WebBackendConnectionRead) GetSchemaChange() SchemaChange {
	if o == nil {
		return SchemaChange("")
	}
	return o.SchemaChange
}

func (o *WebBackendConnectionRead) GetSource() SourceRead {
	if o == nil {
		return SourceRead{}
	}
	return o.Source
}

func (o *WebBackendConnectionRead) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *WebBackendConnectionRead) GetStatus() ConnectionStatus {
	if o == nil {
		return ConnectionStatus("")
	}
	return o.Status
}

func (o *WebBackendConnectionRead) GetSyncCatalog() AirbyteCatalog {
	if o == nil {
		return AirbyteCatalog{}
	}
	return o.SyncCatalog
}
