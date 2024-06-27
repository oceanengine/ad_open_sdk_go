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

// ToolsQuotaGetV2DeliveryRange
type ToolsQuotaGetV2DeliveryRange string

// List of tools_quota_get_v2_delivery_range
const (
	DEFAULT_ToolsQuotaGetV2DeliveryRange ToolsQuotaGetV2DeliveryRange = "DEFAULT"
	UNION_ToolsQuotaGetV2DeliveryRange   ToolsQuotaGetV2DeliveryRange = "UNION"
)

// All allowed values of ToolsQuotaGetV2DeliveryRange enum
var AllowedToolsQuotaGetV2DeliveryRangeEnumValues = []ToolsQuotaGetV2DeliveryRange{
	"DEFAULT",
	"UNION",
}

// NewToolsQuotaGetV2DeliveryRangeFromValue returns a pointer to a valid ToolsQuotaGetV2DeliveryRange
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsQuotaGetV2DeliveryRangeFromValue(v string) (*ToolsQuotaGetV2DeliveryRange, error) {
	ev := ToolsQuotaGetV2DeliveryRange(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsQuotaGetV2DeliveryRange: valid values are %v", v, AllowedToolsQuotaGetV2DeliveryRangeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsQuotaGetV2DeliveryRange) IsValid() bool {
	for _, existing := range AllowedToolsQuotaGetV2DeliveryRangeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_quota_get_v2_delivery_range value
func (v ToolsQuotaGetV2DeliveryRange) Ptr() *ToolsQuotaGetV2DeliveryRange {
	return &v
}
