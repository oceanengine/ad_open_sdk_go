/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2Platform
type ToolsEstimateAudienceV2Platform string

// List of tools_estimate_audience_v2_platform
const (
	NONE_ToolsEstimateAudienceV2Platform    ToolsEstimateAudienceV2Platform = "NONE"
	ANDROID_ToolsEstimateAudienceV2Platform ToolsEstimateAudienceV2Platform = "ANDROID"
	PC_ToolsEstimateAudienceV2Platform      ToolsEstimateAudienceV2Platform = "PC"
	IPAD_ToolsEstimateAudienceV2Platform    ToolsEstimateAudienceV2Platform = "IPAD"
	WAP_ToolsEstimateAudienceV2Platform     ToolsEstimateAudienceV2Platform = "WAP"
	IOS_ToolsEstimateAudienceV2Platform     ToolsEstimateAudienceV2Platform = "IOS"
)

// Ptr returns reference to tools_estimate_audience_v2_platform value
func (v ToolsEstimateAudienceV2Platform) Ptr() *ToolsEstimateAudienceV2Platform {
	return &v
}
