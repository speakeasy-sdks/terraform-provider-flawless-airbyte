// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PrivateDestinationDefinitionRead struct {
	DestinationDefinition DestinationDefinitionRead `json:"destinationDefinition"`
	Granted               bool                      `json:"granted"`
}

func (o *PrivateDestinationDefinitionRead) GetDestinationDefinition() DestinationDefinitionRead {
	if o == nil {
		return DestinationDefinitionRead{}
	}
	return o.DestinationDefinition
}

func (o *PrivateDestinationDefinitionRead) GetGranted() bool {
	if o == nil {
		return false
	}
	return o.Granted
}
