/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30DataListAuthorizationStatus
type ToolsWechatGameListV30DataListAuthorizationStatus string

// List of tools_wechat_game_list_v3.0_data_list_authorization_status
const (
	AUTHORIZED_ToolsWechatGameListV30DataListAuthorizationStatus           ToolsWechatGameListV30DataListAuthorizationStatus = "AUTHORIZED"
	AUTHORIZATION_FAILED_ToolsWechatGameListV30DataListAuthorizationStatus ToolsWechatGameListV30DataListAuthorizationStatus = "AUTHORIZATION_FAILED"
	UNAUTHORIZED_ToolsWechatGameListV30DataListAuthorizationStatus         ToolsWechatGameListV30DataListAuthorizationStatus = "UNAUTHORIZED"
)

// All allowed values of ToolsWechatGameListV30DataListAuthorizationStatus enum
var AllowedToolsWechatGameListV30DataListAuthorizationStatusEnumValues = []ToolsWechatGameListV30DataListAuthorizationStatus{
	"AUTHORIZED",
	"AUTHORIZATION_FAILED",
	"UNAUTHORIZED",
}

// NewToolsWechatGameListV30DataListAuthorizationStatusFromValue returns a pointer to a valid ToolsWechatGameListV30DataListAuthorizationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30DataListAuthorizationStatusFromValue(v string) (*ToolsWechatGameListV30DataListAuthorizationStatus, error) {
	ev := ToolsWechatGameListV30DataListAuthorizationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30DataListAuthorizationStatus: valid values are %v", v, AllowedToolsWechatGameListV30DataListAuthorizationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30DataListAuthorizationStatus) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30DataListAuthorizationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_data_list_authorization_status value
func (v ToolsWechatGameListV30DataListAuthorizationStatus) Ptr() *ToolsWechatGameListV30DataListAuthorizationStatus {
	return &v
}
