// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceDiscoverSchemaRequestBody struct {
	ConnectionID       *string `json:"connectionId,omitempty"`
	DisableCache       *bool   `json:"disable_cache,omitempty"`
	NotifySchemaChange *bool   `json:"notifySchemaChange,omitempty"`
	SourceID           string  `json:"sourceId"`
}
