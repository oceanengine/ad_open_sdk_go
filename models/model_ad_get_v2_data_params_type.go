/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataParamsType
type AdGetV2DataParamsType string

// List of ad_get_v2_data_params_type
const (
	CUSTOM_AdGetV2DataParamsType AdGetV2DataParamsType = "CUSTOM"
	DPA_AdGetV2DataParamsType    AdGetV2DataParamsType = "DPA"
)

// Ptr returns reference to ad_get_v2_data_params_type value
func (v AdGetV2DataParamsType) Ptr() *AdGetV2DataParamsType {
	return &v
}
