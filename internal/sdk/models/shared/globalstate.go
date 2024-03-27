// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GlobalState struct {
	SharedState  *StateBlob    `json:"shared_state,omitempty"`
	StreamStates []StreamState `json:"streamStates"`
}

func (o *GlobalState) GetSharedState() *StateBlob {
	if o == nil {
		return nil
	}
	return o.SharedState
}

func (o *GlobalState) GetStreamStates() []StreamState {
	if o == nil {
		return []StreamState{}
	}
	return o.StreamStates
}