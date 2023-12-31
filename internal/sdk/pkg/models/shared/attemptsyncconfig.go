// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AttemptSyncConfig struct {
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	DestinationConfiguration interface{} `json:"destinationConfiguration"`
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	SourceConfiguration interface{} `json:"sourceConfiguration"`
	// Contains the state for a connection. The stateType field identifies what type of state it is. Only the field corresponding to that type will be set, the rest will be null. If stateType=not_set, then none of the fields will be set.
	State *ConnectionState `json:"state,omitempty"`
}

func (o *AttemptSyncConfig) GetDestinationConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.DestinationConfiguration
}

func (o *AttemptSyncConfig) GetSourceConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.SourceConfiguration
}

func (o *AttemptSyncConfig) GetState() *ConnectionState {
	if o == nil {
		return nil
	}
	return o.State
}
