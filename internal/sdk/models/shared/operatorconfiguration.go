// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OperatorConfiguration struct {
	Dbt           *OperatorDbt           `json:"dbt,omitempty"`
	Normalization *OperatorNormalization `json:"normalization,omitempty"`
	OperatorType  OperatorType           `json:"operatorType"`
	Webhook       *OperatorWebhook       `json:"webhook,omitempty"`
}

func (o *OperatorConfiguration) GetDbt() *OperatorDbt {
	if o == nil {
		return nil
	}
	return o.Dbt
}

func (o *OperatorConfiguration) GetNormalization() *OperatorNormalization {
	if o == nil {
		return nil
	}
	return o.Normalization
}

func (o *OperatorConfiguration) GetOperatorType() OperatorType {
	if o == nil {
		return OperatorType("")
	}
	return o.OperatorType
}

func (o *OperatorConfiguration) GetWebhook() *OperatorWebhook {
	if o == nil {
		return nil
	}
	return o.Webhook
}