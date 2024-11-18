/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2Gender
type ToolsEstimateAudienceV2Gender string

// List of tools_estimate_audience_v2_gender
const (
	NONE_ToolsEstimateAudienceV2Gender             ToolsEstimateAudienceV2Gender = "NONE"
	GENDER_MALE_ToolsEstimateAudienceV2Gender      ToolsEstimateAudienceV2Gender = "GENDER_MALE"
	GENDER_UNLIMITED_ToolsEstimateAudienceV2Gender ToolsEstimateAudienceV2Gender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ToolsEstimateAudienceV2Gender    ToolsEstimateAudienceV2Gender = "GENDER_FEMALE"
)

// Ptr returns reference to tools_estimate_audience_v2_gender value
func (v ToolsEstimateAudienceV2Gender) Ptr() *ToolsEstimateAudienceV2Gender {
	return &v
}
