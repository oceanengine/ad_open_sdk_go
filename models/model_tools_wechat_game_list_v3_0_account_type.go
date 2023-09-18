/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsWechatGameListV30AccountType
type ToolsWechatGameListV30AccountType string

// List of tools_wechat_game_list_v3.0_account_type
const (
	BP_ToolsWechatGameListV30AccountType ToolsWechatGameListV30AccountType = "BP"
	AD_ToolsWechatGameListV30AccountType ToolsWechatGameListV30AccountType = "AD"
)

// All allowed values of ToolsWechatGameListV30AccountType enum
var AllowedToolsWechatGameListV30AccountTypeEnumValues = []ToolsWechatGameListV30AccountType{
	"BP",
	"AD",
}

// NewToolsWechatGameListV30AccountTypeFromValue returns a pointer to a valid ToolsWechatGameListV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsWechatGameListV30AccountTypeFromValue(v string) (*ToolsWechatGameListV30AccountType, error) {
	ev := ToolsWechatGameListV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsWechatGameListV30AccountType: valid values are %v", v, AllowedToolsWechatGameListV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsWechatGameListV30AccountType) IsValid() bool {
	for _, existing := range AllowedToolsWechatGameListV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_wechat_game_list_v3.0_account_type value
func (v ToolsWechatGameListV30AccountType) Ptr() *ToolsWechatGameListV30AccountType {
	return &v
}