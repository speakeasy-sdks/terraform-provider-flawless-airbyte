// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SlackNotificationConfiguration struct {
	Webhook string `json:"webhook"`
}

func (o *SlackNotificationConfiguration) GetWebhook() string {
	if o == nil {
		return ""
	}
	return o.Webhook
}
