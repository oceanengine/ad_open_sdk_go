/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30DataListAccountType
type ToolsWechatGameListV30DataListAccountType string

// List of tools_wechat_game_list_v3.0_data_list_account_type
const (
	AD_ToolsWechatGameListV30DataListAccountType ToolsWechatGameListV30DataListAccountType = "AD"
	BP_ToolsWechatGameListV30DataListAccountType ToolsWechatGameListV30DataListAccountType = "BP"
)

// All allowed values of ToolsWechatGameListV30DataListAccountType enum
var AllowedToolsWechatGameListV30DataListAccountTypeEnumValues = []ToolsWechatGameListV30DataListAccountType{
	"AD",
	"BP",
}

// NewToolsWechatGameListV30DataListAccountTypeFromValue returns a pointer to a valid ToolsWechatGameListV30DataListAccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30DataListAccountTypeFromValue(v string) (*ToolsWechatGameListV30DataListAccountType, error) {
	ev := ToolsWechatGameListV30DataListAccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30DataListAccountType: valid values are %v", v, AllowedToolsWechatGameListV30DataListAccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30DataListAccountType) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30DataListAccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_data_list_account_type value
func (v ToolsWechatGameListV30DataListAccountType) Ptr() *ToolsWechatGameListV30DataListAccountType {
	return &v
}
