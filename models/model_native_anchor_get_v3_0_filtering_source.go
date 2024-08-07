/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// NativeAnchorGetV30FilteringSource
type NativeAnchorGetV30FilteringSource string

// List of native_anchor_get_v3.0_filtering_source
const (
	AUTO_NativeAnchorGetV30FilteringSource   NativeAnchorGetV30FilteringSource = "AUTO"
	CUSTOM_NativeAnchorGetV30FilteringSource NativeAnchorGetV30FilteringSource = "CUSTOM"
	MANUAL_NativeAnchorGetV30FilteringSource NativeAnchorGetV30FilteringSource = "MANUAL"
)

// All allowed values of NativeAnchorGetV30FilteringSource enum
var AllowedNativeAnchorGetV30FilteringSourceEnumValues = []NativeAnchorGetV30FilteringSource{
	"AUTO",
	"CUSTOM",
	"MANUAL",
}

// NewNativeAnchorGetV30FilteringSourceFromValue returns a pointer to a valid NativeAnchorGetV30FilteringSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30FilteringSourceFromValue(v string) (*NativeAnchorGetV30FilteringSource, error) {
	ev := NativeAnchorGetV30FilteringSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30FilteringSource: valid values are %v", v, AllowedNativeAnchorGetV30FilteringSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30FilteringSource) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30FilteringSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_filtering_source value
func (v NativeAnchorGetV30FilteringSource) Ptr() *NativeAnchorGetV30FilteringSource {
	return &v
}
