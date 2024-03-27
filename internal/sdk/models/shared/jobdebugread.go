// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobDebugRead struct {
	AirbyteVersion        string                    `json:"airbyteVersion"`
	ConfigID              string                    `json:"configId"`
	ConfigType            JobConfigType             `json:"configType"`
	DestinationDefinition DestinationDefinitionRead `json:"destinationDefinition"`
	ID                    int64                     `json:"id"`
	SourceDefinition      SourceDefinitionRead      `json:"sourceDefinition"`
	Status                JobStatus                 `json:"status"`
}

func (o *JobDebugRead) GetAirbyteVersion() string {
	if o == nil {
		return ""
	}
	return o.AirbyteVersion
}

func (o *JobDebugRead) GetConfigID() string {
	if o == nil {
		return ""
	}
	return o.ConfigID
}

func (o *JobDebugRead) GetConfigType() JobConfigType {
	if o == nil {
		return JobConfigType("")
	}
	return o.ConfigType
}

func (o *JobDebugRead) GetDestinationDefinition() DestinationDefinitionRead {
	if o == nil {
		return DestinationDefinitionRead{}
	}
	return o.DestinationDefinition
}

func (o *JobDebugRead) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *JobDebugRead) GetSourceDefinition() SourceDefinitionRead {
	if o == nil {
		return SourceDefinitionRead{}
	}
	return o.SourceDefinition
}

func (o *JobDebugRead) GetStatus() JobStatus {
	if o == nil {
		return JobStatus("")
	}
	return o.Status
}