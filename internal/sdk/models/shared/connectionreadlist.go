// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ConnectionReadList struct {
	Connections []ConnectionRead `json:"connections"`
}

func (o *ConnectionReadList) GetConnections() []ConnectionRead {
	if o == nil {
		return []ConnectionRead{}
	}
	return o.Connections
}