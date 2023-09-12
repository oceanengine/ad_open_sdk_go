/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueWechatPoolListV2DataItemsWechatType
type ClueWechatPoolListV2DataItemsWechatType string

// List of clue_wechat_pool_list_v2_data_items_wechat_type
const (
	PERSONAL_ClueWechatPoolListV2DataItemsWechatType   ClueWechatPoolListV2DataItemsWechatType = "PERSONAL"
	PUBLIC_ClueWechatPoolListV2DataItemsWechatType     ClueWechatPoolListV2DataItemsWechatType = "PUBLIC"
	ENTERPRISE_ClueWechatPoolListV2DataItemsWechatType ClueWechatPoolListV2DataItemsWechatType = "ENTERPRISE"
)

// All allowed values of ClueWechatPoolListV2DataItemsWechatType enum
var AllowedClueWechatPoolListV2DataItemsWechatTypeEnumValues = []ClueWechatPoolListV2DataItemsWechatType{
	"PERSONAL",
	"PUBLIC",
	"ENTERPRISE",
}

// NewClueWechatPoolListV2DataItemsWechatTypeFromValue returns a pointer to a valid ClueWechatPoolListV2DataItemsWechatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatPoolListV2DataItemsWechatTypeFromValue(v string) (*ClueWechatPoolListV2DataItemsWechatType, error) {
	ev := ClueWechatPoolListV2DataItemsWechatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatPoolListV2DataItemsWechatType: valid values are %v", v, AllowedClueWechatPoolListV2DataItemsWechatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatPoolListV2DataItemsWechatType) IsValid() bool {
	for _, existing := range AllowedClueWechatPoolListV2DataItemsWechatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_pool_list_v2_data_items_wechat_type value
func (v ClueWechatPoolListV2DataItemsWechatType) Ptr() *ClueWechatPoolListV2DataItemsWechatType {
	return &v
}
