// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LogRead struct {
	LogLines []string `json:"logLines"`
}

func (o *LogRead) GetLogLines() []string {
	if o == nil {
		return []string{}
	}
	return o.LogLines
}
