// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InvalidInputExceptionInfo struct {
	ExceptionClassName *string                `json:"exceptionClassName,omitempty"`
	ExceptionStack     []string               `json:"exceptionStack,omitempty"`
	Message            string                 `json:"message"`
	ValidationErrors   []InvalidInputProperty `json:"validationErrors"`
}

func (o *InvalidInputExceptionInfo) GetExceptionClassName() *string {
	if o == nil {
		return nil
	}
	return o.ExceptionClassName
}

func (o *InvalidInputExceptionInfo) GetExceptionStack() []string {
	if o == nil {
		return nil
	}
	return o.ExceptionStack
}

func (o *InvalidInputExceptionInfo) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *InvalidInputExceptionInfo) GetValidationErrors() []InvalidInputProperty {
	if o == nil {
		return []InvalidInputProperty{}
	}
	return o.ValidationErrors
}
