/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsEstimateAudienceV2Carrier
type ToolsEstimateAudienceV2Carrier string

// List of tools_estimate_audience_v2_carrier
const (
	TELCOM_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "TELCOM"
	MOBILE_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "MOBILE"
	UNICOM_ToolsEstimateAudienceV2Carrier ToolsEstimateAudienceV2Carrier = "UNICOM"
)

// Ptr returns reference to tools_estimate_audience_v2_carrier value
func (v ToolsEstimateAudienceV2Carrier) Ptr() *ToolsEstimateAudienceV2Carrier {
	return &v
}
