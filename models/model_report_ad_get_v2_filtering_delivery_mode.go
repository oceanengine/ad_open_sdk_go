/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2FilteringDeliveryMode
type ReportAdGetV2FilteringDeliveryMode string

// List of report_ad_get_v2_filtering_delivery_mode
const (
	STANDARD_ReportAdGetV2FilteringDeliveryMode   ReportAdGetV2FilteringDeliveryMode = "STANDARD"
	ADLAB_FREE_ReportAdGetV2FilteringDeliveryMode ReportAdGetV2FilteringDeliveryMode = "ADLAB_FREE"
)

// Ptr returns reference to report_ad_get_v2_filtering_delivery_mode value
func (v ReportAdGetV2FilteringDeliveryMode) Ptr() *ReportAdGetV2FilteringDeliveryMode {
	return &v
}
