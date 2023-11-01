// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AirbyteStream - the immutable schema defined by the source
type AirbyteStream struct {
	// Path to the field that will be used to determine if a record is new or modified since the last sync. If not provided by the source, the end user will have to specify the comparable themselves.
	DefaultCursorField []string `json:"defaultCursorField,omitempty"`
	// Stream schema using Json Schema specs.
	JSONSchema *StreamJSONSchema `json:"jsonSchema,omitempty"`
	// Stream's name.
	Name string `json:"name"`
	// Optional Source-defined namespace. Airbyte streams from the same sources should have the same namespace. Currently only used by JDBC destinations to determine what schema to write to.
	Namespace *string `json:"namespace,omitempty"`
	// If the source defines the cursor field, then any other cursor field inputs will be ignored. If it does not, either the user_provided one is used, or the default one is used as a backup.
	SourceDefinedCursor *bool `json:"sourceDefinedCursor,omitempty"`
	// If the source defines the primary key, paths to the fields that will be used as a primary key. If not provided by the source, the end user will have to specify the primary key themselves.
	SourceDefinedPrimaryKey [][]string `json:"sourceDefinedPrimaryKey,omitempty"`
	SupportedSyncModes      []SyncMode `json:"supportedSyncModes,omitempty"`
}

func (o *AirbyteStream) GetDefaultCursorField() []string {
	if o == nil {
		return nil
	}
	return o.DefaultCursorField
}

func (o *AirbyteStream) GetJSONSchema() *StreamJSONSchema {
	if o == nil {
		return nil
	}
	return o.JSONSchema
}

func (o *AirbyteStream) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *AirbyteStream) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *AirbyteStream) GetSourceDefinedCursor() *bool {
	if o == nil {
		return nil
	}
	return o.SourceDefinedCursor
}

func (o *AirbyteStream) GetSourceDefinedPrimaryKey() [][]string {
	if o == nil {
		return nil
	}
	return o.SourceDefinedPrimaryKey
}

func (o *AirbyteStream) GetSupportedSyncModes() []SyncMode {
	if o == nil {
		return nil
	}
	return o.SupportedSyncModes
}
