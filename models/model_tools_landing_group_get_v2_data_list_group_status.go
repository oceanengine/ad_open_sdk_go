/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLandingGroupGetV2DataListGroupStatus
type ToolsLandingGroupGetV2DataListGroupStatus string

// List of tools_landing_group_get_v2_data_list_group_status
const (
	LANDING_GROUP_STATUS_AVAILABLE_ToolsLandingGroupGetV2DataListGroupStatus   ToolsLandingGroupGetV2DataListGroupStatus = "LANDING_GROUP_STATUS_AVAILABLE"
	LANDING_GROUP_STATUS_UNAVAILABLE_ToolsLandingGroupGetV2DataListGroupStatus ToolsLandingGroupGetV2DataListGroupStatus = "LANDING_GROUP_STATUS_UNAVAILABLE"
)

// Ptr returns reference to tools_landing_group_get_v2_data_list_group_status value
func (v ToolsLandingGroupGetV2DataListGroupStatus) Ptr() *ToolsLandingGroupGetV2DataListGroupStatus {
	return &v
}
