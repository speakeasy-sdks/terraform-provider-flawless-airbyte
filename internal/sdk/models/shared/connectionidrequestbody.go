// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionIDRequestBody struct {
	ConnectionID string `json:"connectionId"`
}

func (o *ConnectionIDRequestBody) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}