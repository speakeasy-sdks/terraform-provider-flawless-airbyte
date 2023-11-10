// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceCreate struct {
	// The values required to configure the source. The schema for this must match the schema return by source_definition_specifications/get for the source.
	ConnectionConfiguration interface{} `json:"connectionConfiguration"`
	Name                    string      `json:"name"`
	SecretID                *string     `json:"secretId,omitempty"`
	SourceDefinitionID      string      `json:"sourceDefinitionId"`
	WorkspaceID             string      `json:"workspaceId"`
}

func (o *SourceCreate) GetConnectionConfiguration() interface{} {
	if o == nil {
		return nil
	}
	return o.ConnectionConfiguration
}

func (o *SourceCreate) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SourceCreate) GetSecretID() *string {
	if o == nil {
		return nil
	}
	return o.SecretID
}

func (o *SourceCreate) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}

func (o *SourceCreate) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
