/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataAudienceAppBehaviorTarget
type AdGetV2DataAudienceAppBehaviorTarget string

// List of ad_get_v2_data_audience_app_behavior_target
const (
	CATEGORY_AdGetV2DataAudienceAppBehaviorTarget AdGetV2DataAudienceAppBehaviorTarget = "CATEGORY"
	NONE_AdGetV2DataAudienceAppBehaviorTarget     AdGetV2DataAudienceAppBehaviorTarget = "NONE"
	APP_AdGetV2DataAudienceAppBehaviorTarget      AdGetV2DataAudienceAppBehaviorTarget = "APP"
)

// Ptr returns reference to ad_get_v2_data_audience_app_behavior_target value
func (v AdGetV2DataAudienceAppBehaviorTarget) Ptr() *AdGetV2DataAudienceAppBehaviorTarget {
	return &v
}
