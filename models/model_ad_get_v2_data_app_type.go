/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAppType
type AdGetV2DataAppType string

// List of ad_get_v2_data_app_type
const (
	APP_ANDROID_AdGetV2DataAppType AdGetV2DataAppType = "APP_ANDROID"
	APP_IOS_AdGetV2DataAppType     AdGetV2DataAppType = "APP_IOS"
)

// Ptr returns reference to ad_get_v2_data_app_type value
func (v AdGetV2DataAppType) Ptr() *AdGetV2DataAppType {
	return &v
}
