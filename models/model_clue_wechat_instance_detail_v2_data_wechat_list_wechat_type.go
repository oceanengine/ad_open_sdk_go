/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueWechatInstanceDetailV2DataWechatListWechatType
type ClueWechatInstanceDetailV2DataWechatListWechatType string

// List of clue_wechat_instance_detail_v2_data_wechat_list_wechat_type
const (
	PERSONAL_ClueWechatInstanceDetailV2DataWechatListWechatType   ClueWechatInstanceDetailV2DataWechatListWechatType = "PERSONAL"
	PUBLIC_ClueWechatInstanceDetailV2DataWechatListWechatType     ClueWechatInstanceDetailV2DataWechatListWechatType = "PUBLIC"
	ENTERPRISE_ClueWechatInstanceDetailV2DataWechatListWechatType ClueWechatInstanceDetailV2DataWechatListWechatType = "ENTERPRISE"
)

// All allowed values of ClueWechatInstanceDetailV2DataWechatListWechatType enum
var AllowedClueWechatInstanceDetailV2DataWechatListWechatTypeEnumValues = []ClueWechatInstanceDetailV2DataWechatListWechatType{
	"PERSONAL",
	"PUBLIC",
	"ENTERPRISE",
}

// NewClueWechatInstanceDetailV2DataWechatListWechatTypeFromValue returns a pointer to a valid ClueWechatInstanceDetailV2DataWechatListWechatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceDetailV2DataWechatListWechatTypeFromValue(v string) (*ClueWechatInstanceDetailV2DataWechatListWechatType, error) {
	ev := ClueWechatInstanceDetailV2DataWechatListWechatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceDetailV2DataWechatListWechatType: valid values are %v", v, AllowedClueWechatInstanceDetailV2DataWechatListWechatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceDetailV2DataWechatListWechatType) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceDetailV2DataWechatListWechatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_detail_v2_data_wechat_list_wechat_type value
func (v ClueWechatInstanceDetailV2DataWechatListWechatType) Ptr() *ClueWechatInstanceDetailV2DataWechatListWechatType {
	return &v
}
