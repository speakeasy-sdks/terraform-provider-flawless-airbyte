// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDefinitionIDRequestBody struct {
	SourceDefinitionID string `json:"sourceDefinitionId"`
}

func (o *SourceDefinitionIDRequestBody) GetSourceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.SourceDefinitionID
}