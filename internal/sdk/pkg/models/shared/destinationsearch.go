// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationSearch struct {
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	ConnectionConfiguration interface{} `json:"connectionConfiguration,omitempty"`
	DestinationDefinitionID *string     `json:"destinationDefinitionId,omitempty"`
	DestinationID           *string     `json:"destinationId,omitempty"`
	DestinationName         *string     `json:"destinationName,omitempty"`
	Name                    *string     `json:"name,omitempty"`
	WorkspaceID             *string     `json:"workspaceId,omitempty"`
}

func (o *DestinationSearch) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *DestinationSearch) GetDestinationDefinitionID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationDefinitionID
}

func (o *DestinationSearch) GetDestinationID() *string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *DestinationSearch) GetDestinationName() *string {
	if o == nil {
		return nil
	}
	return o.DestinationName
}

func (o *DestinationSearch) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *DestinationSearch) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}
