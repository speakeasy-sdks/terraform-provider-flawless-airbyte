// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type JobInfoLightRead struct {
	Job JobRead `json:"job"`
}

func (o *JobInfoLightRead) GetJob() JobRead {
	if o == nil {
		return JobRead{}
	}
	return o.Job
}