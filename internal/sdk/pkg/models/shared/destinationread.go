// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationRead struct {
	// The values required to configure the destination. The schema for this must match the schema return by destination_definition_specifications/get for the destinationDefinition.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	DestinationDefinitionID string      `json:"destinationDefinitionId"`
	DestinationID           string      `json:"destinationId"`
	DestinationName         string      `json:"destinationName"`
	Icon                    *string     `json:"icon,omitempty"`
	Name                    string      `json:"name"`
	WorkspaceID             string      `json:"workspaceId"`
}

func (o *DestinationRead) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *DestinationRead) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationRead) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *DestinationRead) GetDestinationName() string {
	if o == nil {
		return ""
	}
	return o.DestinationName
}

func (o *DestinationRead) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *DestinationRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
