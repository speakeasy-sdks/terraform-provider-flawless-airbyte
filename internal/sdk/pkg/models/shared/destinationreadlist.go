// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationReadList struct {
	Destinations []DestinationRead `json:"destinations"`
}

func (o *DestinationReadList) GetDestinations() []DestinationRead {
	if o == nil {
		return []DestinationRead{}
	}
	return o.Destinations
}
