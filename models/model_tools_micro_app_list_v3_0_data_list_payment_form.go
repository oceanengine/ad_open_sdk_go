/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroAppListV30DataListPaymentForm
type ToolsMicroAppListV30DataListPaymentForm string

// List of tools_micro_app_list_v3.0_data_list_payment_form
const (
	BOTH_OPTIONS_AVAILABLE_ToolsMicroAppListV30DataListPaymentForm ToolsMicroAppListV30DataListPaymentForm = "BOTH_OPTIONS_AVAILABLE"
	CONTENT_OR_SERVICES_ToolsMicroAppListV30DataListPaymentForm    ToolsMicroAppListV30DataListPaymentForm = "CONTENT_OR_SERVICES"
	UNLOCK_FULL_FEATURES_ToolsMicroAppListV30DataListPaymentForm   ToolsMicroAppListV30DataListPaymentForm = "UNLOCK_FULL_FEATURES"
)

// Ptr returns reference to tools_micro_app_list_v3.0_data_list_payment_form value
func (v ToolsMicroAppListV30DataListPaymentForm) Ptr() *ToolsMicroAppListV30DataListPaymentForm {
	return &v
}
