// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobIDRequestBody struct {
	ID int64 `json:"id"`
}

func (o *JobIDRequestBody) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}