/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalReportAccountGetV30FilteringDeliveryMode
type LocalReportAccountGetV30FilteringDeliveryMode string

// List of local_report_account_get_v3.0_filtering_delivery_mode
const (
	CDP_AUTO_MODE_LocalReportAccountGetV30FilteringDeliveryMode LocalReportAccountGetV30FilteringDeliveryMode = "CDP_AUTO_MODE"
	MANUAL_MODE_LocalReportAccountGetV30FilteringDeliveryMode   LocalReportAccountGetV30FilteringDeliveryMode = "MANUAL_MODE"
)

// Ptr returns reference to local_report_account_get_v3.0_filtering_delivery_mode value
func (v LocalReportAccountGetV30FilteringDeliveryMode) Ptr() *LocalReportAccountGetV30FilteringDeliveryMode {
	return &v
}
