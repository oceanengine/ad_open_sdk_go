/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2AppBehaviorTarget
type ToolsEstimateAudienceV2AppBehaviorTarget string

// List of tools_estimate_audience_v2_app_behavior_target
const (
	NONE_ToolsEstimateAudienceV2AppBehaviorTarget     ToolsEstimateAudienceV2AppBehaviorTarget = "NONE"
	APP_ToolsEstimateAudienceV2AppBehaviorTarget      ToolsEstimateAudienceV2AppBehaviorTarget = "APP"
	CATEGORY_ToolsEstimateAudienceV2AppBehaviorTarget ToolsEstimateAudienceV2AppBehaviorTarget = "CATEGORY"
)

// Ptr returns reference to tools_estimate_audience_v2_app_behavior_target value
func (v ToolsEstimateAudienceV2AppBehaviorTarget) Ptr() *ToolsEstimateAudienceV2AppBehaviorTarget {
	return &v
}
