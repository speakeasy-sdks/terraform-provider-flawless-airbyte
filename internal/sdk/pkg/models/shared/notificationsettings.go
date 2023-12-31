// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type NotificationSettings struct {
	SendOnConnectionUpdate               *NotificationItem `json:"sendOnConnectionUpdate,omitempty"`
	SendOnConnectionUpdateActionRequired *NotificationItem `json:"sendOnConnectionUpdateActionRequired,omitempty"`
	SendOnFailure                        *NotificationItem `json:"sendOnFailure,omitempty"`
	SendOnSuccess                        *NotificationItem `json:"sendOnSuccess,omitempty"`
	SendOnSyncDisabled                   *NotificationItem `json:"sendOnSyncDisabled,omitempty"`
	SendOnSyncDisabledWarning            *NotificationItem `json:"sendOnSyncDisabledWarning,omitempty"`
}

func (o *NotificationSettings) GetSendOnConnectionUpdate() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnConnectionUpdate
}

func (o *NotificationSettings) GetSendOnConnectionUpdateActionRequired() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnConnectionUpdateActionRequired
}

func (o *NotificationSettings) GetSendOnFailure() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnFailure
}

func (o *NotificationSettings) GetSendOnSuccess() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnSuccess
}

func (o *NotificationSettings) GetSendOnSyncDisabled() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnSyncDisabled
}

func (o *NotificationSettings) GetSendOnSyncDisabledWarning() *NotificationItem {
	if o == nil {
		return nil
	}
	return o.SendOnSyncDisabledWarning
}
