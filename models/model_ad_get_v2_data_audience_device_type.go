/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceDeviceType
type AdGetV2DataAudienceDeviceType string

// List of ad_get_v2_data_audience_device_type
const (
	MOBILE_AdGetV2DataAudienceDeviceType AdGetV2DataAudienceDeviceType = "MOBILE"
	PAD_AdGetV2DataAudienceDeviceType    AdGetV2DataAudienceDeviceType = "PAD"
)

// Ptr returns reference to ad_get_v2_data_audience_device_type value
func (v AdGetV2DataAudienceDeviceType) Ptr() *AdGetV2DataAudienceDeviceType {
	return &v
}
