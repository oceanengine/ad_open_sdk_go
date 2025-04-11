/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLiveAuthorizeListV2DataListStatus
type ToolsLiveAuthorizeListV2DataListStatus string

// List of tools_live_authorize_list_v2_data_list_status
const (
	APPLY_OVERDUE_ToolsLiveAuthorizeListV2DataListStatus     ToolsLiveAuthorizeListV2DataListStatus = "APPLY_OVERDUE"
	AUTHORIZE_OVERDUE_ToolsLiveAuthorizeListV2DataListStatus ToolsLiveAuthorizeListV2DataListStatus = "AUTHORIZE_OVERDUE"
	APPLYING_ToolsLiveAuthorizeListV2DataListStatus          ToolsLiveAuthorizeListV2DataListStatus = "APPLYING"
	AUTHORIZING_ToolsLiveAuthorizeListV2DataListStatus       ToolsLiveAuthorizeListV2DataListStatus = "AUTHORIZING"
)

// Ptr returns reference to tools_live_authorize_list_v2_data_list_status value
func (v ToolsLiveAuthorizeListV2DataListStatus) Ptr() *ToolsLiveAuthorizeListV2DataListStatus {
	return &v
}
