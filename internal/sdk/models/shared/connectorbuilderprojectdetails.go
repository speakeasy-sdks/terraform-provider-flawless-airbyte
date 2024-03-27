// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectorBuilderProjectDetails struct {
	// Low code CDK manifest JSON object
	DraftManifest *DeclarativeManifest `json:"draftManifest,omitempty"`
	Name          string               `json:"name"`
}

func (o *ConnectorBuilderProjectDetails) GetDraftManifest() *DeclarativeManifest {
	if o == nil {
		return nil
	}
	return o.DraftManifest
}

func (o *ConnectorBuilderProjectDetails) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}