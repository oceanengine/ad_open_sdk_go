/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertDeepbidReadV2SmartBidType
type ToolsAdConvertDeepbidReadV2SmartBidType string

// List of tools_ad_convert_deepbid_read_v2_smart_bid_type
const (
	SMART_BID_NO_BID_ToolsAdConvertDeepbidReadV2SmartBidType       ToolsAdConvertDeepbidReadV2SmartBidType = "SMART_BID_NO_BID"
	SMART_BID_CUSTOM_ToolsAdConvertDeepbidReadV2SmartBidType       ToolsAdConvertDeepbidReadV2SmartBidType = "SMART_BID_CUSTOM"
	SMART_BID_CONSERVATIVE_ToolsAdConvertDeepbidReadV2SmartBidType ToolsAdConvertDeepbidReadV2SmartBidType = "SMART_BID_CONSERVATIVE"
	SMART_BID_RADICAL_ToolsAdConvertDeepbidReadV2SmartBidType      ToolsAdConvertDeepbidReadV2SmartBidType = "SMART_BID_RADICAL"
)

// All allowed values of ToolsAdConvertDeepbidReadV2SmartBidType enum
var AllowedToolsAdConvertDeepbidReadV2SmartBidTypeEnumValues = []ToolsAdConvertDeepbidReadV2SmartBidType{
	"SMART_BID_NO_BID",
	"SMART_BID_CUSTOM",
	"SMART_BID_CONSERVATIVE",
	"SMART_BID_RADICAL",
}

// NewToolsAdConvertDeepbidReadV2SmartBidTypeFromValue returns a pointer to a valid ToolsAdConvertDeepbidReadV2SmartBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertDeepbidReadV2SmartBidTypeFromValue(v string) (*ToolsAdConvertDeepbidReadV2SmartBidType, error) {
	ev := ToolsAdConvertDeepbidReadV2SmartBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertDeepbidReadV2SmartBidType: valid values are %v", v, AllowedToolsAdConvertDeepbidReadV2SmartBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertDeepbidReadV2SmartBidType) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertDeepbidReadV2SmartBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_deepbid_read_v2_smart_bid_type value
func (v ToolsAdConvertDeepbidReadV2SmartBidType) Ptr() *ToolsAdConvertDeepbidReadV2SmartBidType {
	return &v
}
