// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperationRead struct {
	Name                  string                `json:"name"`
	OperationID           string                `json:"operationId"`
	OperatorConfiguration OperatorConfiguration `json:"operatorConfiguration"`
	WorkspaceID           string                `json:"workspaceId"`
}

func (o *OperationRead) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *OperationRead) GetOperationID() string {
	if o == nil {
		return ""
	}
	return o.OperationID
}

func (o *OperationRead) GetOperatorConfiguration() OperatorConfiguration {
	if o == nil {
		return OperatorConfiguration{}
	}
	return o.OperatorConfiguration
}

func (o *OperationRead) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
