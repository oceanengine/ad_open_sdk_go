/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsIsSupportUniversalGetV2DeepBidType
type ToolsIsSupportUniversalGetV2DeepBidType string

// List of tools_is_support_universal_get_v2_deep_bid_type
const (
	AUTO_MIN_SECOND_STAGE_ToolsIsSupportUniversalGetV2DeepBidType ToolsIsSupportUniversalGetV2DeepBidType = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_ToolsIsSupportUniversalGetV2DeepBidType        ToolsIsSupportUniversalGetV2DeepBidType = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_ToolsIsSupportUniversalGetV2DeepBidType      ToolsIsSupportUniversalGetV2DeepBidType = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_ToolsIsSupportUniversalGetV2DeepBidType          ToolsIsSupportUniversalGetV2DeepBidType = "DEEP_BID_MIN"
	DEEP_BID_PACING_ToolsIsSupportUniversalGetV2DeepBidType       ToolsIsSupportUniversalGetV2DeepBidType = "DEEP_BID_PACING"
	MIN_SECOND_STAGE_ToolsIsSupportUniversalGetV2DeepBidType      ToolsIsSupportUniversalGetV2DeepBidType = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_ToolsIsSupportUniversalGetV2DeepBidType   ToolsIsSupportUniversalGetV2DeepBidType = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_ToolsIsSupportUniversalGetV2DeepBidType       ToolsIsSupportUniversalGetV2DeepBidType = "ROI_COEFFICIENT"
	ROI_PACING_ToolsIsSupportUniversalGetV2DeepBidType            ToolsIsSupportUniversalGetV2DeepBidType = "ROI_PACING"
	SMARTBID_ToolsIsSupportUniversalGetV2DeepBidType              ToolsIsSupportUniversalGetV2DeepBidType = "SMARTBID"
)

// All allowed values of ToolsIsSupportUniversalGetV2DeepBidType enum
var AllowedToolsIsSupportUniversalGetV2DeepBidTypeEnumValues = []ToolsIsSupportUniversalGetV2DeepBidType{
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"ROI_COEFFICIENT",
	"ROI_PACING",
	"SMARTBID",
}

// NewToolsIsSupportUniversalGetV2DeepBidTypeFromValue returns a pointer to a valid ToolsIsSupportUniversalGetV2DeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsIsSupportUniversalGetV2DeepBidTypeFromValue(v string) (*ToolsIsSupportUniversalGetV2DeepBidType, error) {
	ev := ToolsIsSupportUniversalGetV2DeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsIsSupportUniversalGetV2DeepBidType: valid values are %v", v, AllowedToolsIsSupportUniversalGetV2DeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsIsSupportUniversalGetV2DeepBidType) IsValid() bool {
	for _, existing := range AllowedToolsIsSupportUniversalGetV2DeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_is_support_universal_get_v2_deep_bid_type value
func (v ToolsIsSupportUniversalGetV2DeepBidType) Ptr() *ToolsIsSupportUniversalGetV2DeepBidType {
	return &v
}
