/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceActivateType
type AdGetV2DataAudienceActivateType string

// List of ad_get_v2_data_audience_activate_type
const (
	WITH_IN_A_MONTH_AdGetV2DataAudienceActivateType         AdGetV2DataAudienceActivateType = "WITH_IN_A_MONTH"
	ONE_MONTH_2_THREE_MONTH_AdGetV2DataAudienceActivateType AdGetV2DataAudienceActivateType = "ONE_MONTH_2_THREE_MONTH"
	UNLIMITED_AdGetV2DataAudienceActivateType               AdGetV2DataAudienceActivateType = "UNLIMITED"
	THREE_MONTH_EAILIER_AdGetV2DataAudienceActivateType     AdGetV2DataAudienceActivateType = "THREE_MONTH_EAILIER"
)

// Ptr returns reference to ad_get_v2_data_audience_activate_type value
func (v AdGetV2DataAudienceActivateType) Ptr() *AdGetV2DataAudienceActivateType {
	return &v
}
