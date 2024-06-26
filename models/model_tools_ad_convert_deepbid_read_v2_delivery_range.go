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

// ToolsAdConvertDeepbidReadV2DeliveryRange
type ToolsAdConvertDeepbidReadV2DeliveryRange string

// List of tools_ad_convert_deepbid_read_v2_delivery_range
const (
	UNIVERSAL_ToolsAdConvertDeepbidReadV2DeliveryRange ToolsAdConvertDeepbidReadV2DeliveryRange = "UNIVERSAL"
	DEFAULT_ToolsAdConvertDeepbidReadV2DeliveryRange   ToolsAdConvertDeepbidReadV2DeliveryRange = "DEFAULT"
	UNION_ToolsAdConvertDeepbidReadV2DeliveryRange     ToolsAdConvertDeepbidReadV2DeliveryRange = "UNION"
)

// All allowed values of ToolsAdConvertDeepbidReadV2DeliveryRange enum
var AllowedToolsAdConvertDeepbidReadV2DeliveryRangeEnumValues = []ToolsAdConvertDeepbidReadV2DeliveryRange{
	"UNIVERSAL",
	"DEFAULT",
	"UNION",
}

// NewToolsAdConvertDeepbidReadV2DeliveryRangeFromValue returns a pointer to a valid ToolsAdConvertDeepbidReadV2DeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertDeepbidReadV2DeliveryRangeFromValue(v string) (*ToolsAdConvertDeepbidReadV2DeliveryRange, error) {
	ev := ToolsAdConvertDeepbidReadV2DeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertDeepbidReadV2DeliveryRange: valid values are %v", v, AllowedToolsAdConvertDeepbidReadV2DeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertDeepbidReadV2DeliveryRange) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertDeepbidReadV2DeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_deepbid_read_v2_delivery_range value
func (v ToolsAdConvertDeepbidReadV2DeliveryRange) Ptr() *ToolsAdConvertDeepbidReadV2DeliveryRange {
	return &v
}
