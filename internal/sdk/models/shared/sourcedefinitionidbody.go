// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDefinitionIDBody struct {
	SourceDefinitionID string `json:"sourceDefinitionId"`
}

func (o *SourceDefinitionIDBody) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}