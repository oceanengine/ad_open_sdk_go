/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2RetargetingType
type ToolsEstimateAudienceV2RetargetingType string

// List of tools_estimate_audience_v2_retargeting_type
const (
	RETARGETING_NONE_ToolsEstimateAudienceV2RetargetingType    ToolsEstimateAudienceV2RetargetingType = "RETARGETING_NONE"
	RETARGETING_EXCLUDE_ToolsEstimateAudienceV2RetargetingType ToolsEstimateAudienceV2RetargetingType = "RETARGETING_EXCLUDE"
	NONE_ToolsEstimateAudienceV2RetargetingType                ToolsEstimateAudienceV2RetargetingType = "NONE"
	RETARGETING_INCLUDE_ToolsEstimateAudienceV2RetargetingType ToolsEstimateAudienceV2RetargetingType = "RETARGETING_INCLUDE"
)

// Ptr returns reference to tools_estimate_audience_v2_retargeting_type value
func (v ToolsEstimateAudienceV2RetargetingType) Ptr() *ToolsEstimateAudienceV2RetargetingType {
	return &v
}
