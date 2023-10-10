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

// ClueWechatInstanceDetailV2DataSubType
type ClueWechatInstanceDetailV2DataSubType string

// List of clue_wechat_instance_detail_v2_data_sub_type
const (
	NORMAL_ClueWechatInstanceDetailV2DataSubType             ClueWechatInstanceDetailV2DataSubType = "NORMAL"
	SINGLE_ClueWechatInstanceDetailV2DataSubType             ClueWechatInstanceDetailV2DataSubType = "SINGLE"
	ENTERPRISE_DEFAULT_ClueWechatInstanceDetailV2DataSubType ClueWechatInstanceDetailV2DataSubType = "ENTERPRISE_DEFAULT"
	ENTERPRISE_CODE_ClueWechatInstanceDetailV2DataSubType    ClueWechatInstanceDetailV2DataSubType = "ENTERPRISE_CODE"
)

// All allowed values of ClueWechatInstanceDetailV2DataSubType enum
var AllowedClueWechatInstanceDetailV2DataSubTypeEnumValues = []ClueWechatInstanceDetailV2DataSubType{
	"NORMAL",
	"SINGLE",
	"ENTERPRISE_DEFAULT",
	"ENTERPRISE_CODE",
}

// NewClueWechatInstanceDetailV2DataSubTypeFromValue returns a pointer to a valid ClueWechatInstanceDetailV2DataSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceDetailV2DataSubTypeFromValue(v string) (*ClueWechatInstanceDetailV2DataSubType, error) {
	ev := ClueWechatInstanceDetailV2DataSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceDetailV2DataSubType: valid values are %v", v, AllowedClueWechatInstanceDetailV2DataSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceDetailV2DataSubType) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceDetailV2DataSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_detail_v2_data_sub_type value
func (v ClueWechatInstanceDetailV2DataSubType) Ptr() *ClueWechatInstanceDetailV2DataSubType {
	return &v
}
