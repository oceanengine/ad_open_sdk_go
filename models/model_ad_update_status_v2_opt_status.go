/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUpdateStatusV2OptStatus
type AdUpdateStatusV2OptStatus string

// List of ad_update_status_v2_opt_status
const (
	DELETE_AdUpdateStatusV2OptStatus  AdUpdateStatusV2OptStatus = "delete"
	ENABLE_AdUpdateStatusV2OptStatus  AdUpdateStatusV2OptStatus = "enable"
	DISABLE_AdUpdateStatusV2OptStatus AdUpdateStatusV2OptStatus = "disable"
)

// Ptr returns reference to ad_update_status_v2_opt_status value
func (v AdUpdateStatusV2OptStatus) Ptr() *AdUpdateStatusV2OptStatus {
	return &v
}
