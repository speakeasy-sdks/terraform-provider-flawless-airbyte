// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateActiveManifestRequestBody struct {
	SourceDefinitionID string `json:"sourceDefinitionId"`
	Version            int64  `json:"version"`
	WorkspaceID        string `json:"workspaceId"`
}

func (o *UpdateActiveManifestRequestBody) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}

func (o *UpdateActiveManifestRequestBody) GetVersion() int64 {
	if o == nil {
		return 0
	}
	return o.Version
}

func (o *UpdateActiveManifestRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
