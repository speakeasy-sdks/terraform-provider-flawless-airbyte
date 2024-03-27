// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationSnippetRead struct {
	DestinationDefinitionID string  `json:"destinationDefinitionId"`
	DestinationID           string  `json:"destinationId"`
	DestinationName         string  `json:"destinationName"`
	Icon                    *string `json:"icon,omitempty"`
	Name                    string  `json:"name"`
}

func (o *DestinationSnippetRead) GetDestinationDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.DestinationDefinitionID
}

func (o *DestinationSnippetRead) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *DestinationSnippetRead) GetDestinationName() string {
	if o == nil {
		return ""
	}
	return o.DestinationName
}

func (o *DestinationSnippetRead) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *DestinationSnippetRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}