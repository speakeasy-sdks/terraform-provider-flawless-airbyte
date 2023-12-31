// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectorBuilderProjectIDWithWorkspaceID struct {
	BuilderProjectID string `json:"builderProjectId"`
	Version          *int64 `json:"version,omitempty"`
	WorkspaceID      string `json:"workspaceId"`
}

func (o *ConnectorBuilderProjectIDWithWorkspaceID) GetBuilderProjectID() string {
	if o == nil {
		return ""
	}
	return o.BuilderProjectID
}

func (o *ConnectorBuilderProjectIDWithWorkspaceID) GetVersion() *int64 {
	if o == nil {
		return nil
	}
	return o.Version
}

func (o *ConnectorBuilderProjectIDWithWorkspaceID) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
