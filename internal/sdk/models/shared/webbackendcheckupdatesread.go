// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WebBackendCheckUpdatesRead - Summary of source and destination definitions that could be updated
type WebBackendCheckUpdatesRead struct {
	DestinationDefinitions int64 `json:"destinationDefinitions"`
	SourceDefinitions      int64 `json:"sourceDefinitions"`
}

func (o *WebBackendCheckUpdatesRead) GetDestinationDefinitions() int64 {
	if o == nil {
		return 0
	}
	return o.DestinationDefinitions
}

func (o *WebBackendCheckUpdatesRead) GetSourceDefinitions() int64 {
	if o == nil {
		return 0
	}
	return o.SourceDefinitions
}