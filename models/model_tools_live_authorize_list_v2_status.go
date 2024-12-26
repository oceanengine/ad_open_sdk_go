/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsLiveAuthorizeListV2Status
type ToolsLiveAuthorizeListV2Status string

// List of tools_live_authorize_list_v2_status
const (
	AUTHORIZING_ToolsLiveAuthorizeListV2Status       ToolsLiveAuthorizeListV2Status = "AUTHORIZING"
	APPLY_OVERDUE_ToolsLiveAuthorizeListV2Status     ToolsLiveAuthorizeListV2Status = "APPLY_OVERDUE"
	AUTHORIZE_OVERDUE_ToolsLiveAuthorizeListV2Status ToolsLiveAuthorizeListV2Status = "AUTHORIZE_OVERDUE"
	APPLYING_ToolsLiveAuthorizeListV2Status          ToolsLiveAuthorizeListV2Status = "APPLYING"
)

// Ptr returns reference to tools_live_authorize_list_v2_status value
func (v ToolsLiveAuthorizeListV2Status) Ptr() *ToolsLiveAuthorizeListV2Status {
	return &v
}
