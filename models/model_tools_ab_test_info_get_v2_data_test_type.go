/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestInfoGetV2DataTestType
type ToolsAbTestInfoGetV2DataTestType string

// List of tools_ab_test_info_get_v2_data_test_type
const (
	AD_ToolsAbTestInfoGetV2DataTestType       ToolsAbTestInfoGetV2DataTestType = "AD"
	CAMPAIGN_ToolsAbTestInfoGetV2DataTestType ToolsAbTestInfoGetV2DataTestType = "CAMPAIGN"
)

// Ptr returns reference to tools_ab_test_info_get_v2_data_test_type value
func (v ToolsAbTestInfoGetV2DataTestType) Ptr() *ToolsAbTestInfoGetV2DataTestType {
	return &v
}
