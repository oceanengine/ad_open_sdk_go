/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30DataListAnchorShareType
type NativeAnchorGetV30DataListAnchorShareType string

// List of native_anchor_get_v3.0_data_list_anchor_share_type
const (
	OWN_ANCHOR_NativeAnchorGetV30DataListAnchorShareType   NativeAnchorGetV30DataListAnchorShareType = "OWN_ANCHOR"
	SHARE_ANCHOR_NativeAnchorGetV30DataListAnchorShareType NativeAnchorGetV30DataListAnchorShareType = "SHARE_ANCHOR"
)

// All allowed values of NativeAnchorGetV30DataListAnchorShareType enum
var AllowedNativeAnchorGetV30DataListAnchorShareTypeEnumValues = []NativeAnchorGetV30DataListAnchorShareType{
	"OWN_ANCHOR",
	"SHARE_ANCHOR",
}

// NewNativeAnchorGetV30DataListAnchorShareTypeFromValue returns a pointer to a valid NativeAnchorGetV30DataListAnchorShareType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30DataListAnchorShareTypeFromValue(v string) (*NativeAnchorGetV30DataListAnchorShareType, error) {
	ev := NativeAnchorGetV30DataListAnchorShareType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30DataListAnchorShareType: valid values are %v", v, AllowedNativeAnchorGetV30DataListAnchorShareTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30DataListAnchorShareType) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30DataListAnchorShareTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_data_list_anchor_share_type value
func (v NativeAnchorGetV30DataListAnchorShareType) Ptr() *NativeAnchorGetV30DataListAnchorShareType {
	return &v
}
