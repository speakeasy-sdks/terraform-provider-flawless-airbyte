// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationDefinitionIDRequestBody struct {
	DestinationDefinitionID string `json:"destinationDefinitionId"`
}

func (o *DestinationDefinitionIDRequestBody) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}
