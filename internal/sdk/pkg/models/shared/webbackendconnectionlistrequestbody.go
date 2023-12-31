// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WebBackendConnectionListRequestBody struct {
	DestinationID []string `json:"destinationId,omitempty"`
	SourceID      []string `json:"sourceId,omitempty"`
	WorkspaceID   string   `json:"workspaceId"`
}

func (o *WebBackendConnectionListRequestBody) GetDestinationID() []string {
	if o == nil {
		return nil
	}
	return o.DestinationID
}

func (o *WebBackendConnectionListRequestBody) GetSourceID() []string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

func (o *WebBackendConnectionListRequestBody) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
