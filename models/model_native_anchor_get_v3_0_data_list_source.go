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

// NativeAnchorGetV30DataListSource
type NativeAnchorGetV30DataListSource string

// List of native_anchor_get_v3.0_data_list_source
const (
	AUTO_NativeAnchorGetV30DataListSource             NativeAnchorGetV30DataListSource = "AUTO"
	CUSTOM_NativeAnchorGetV30DataListSource           NativeAnchorGetV30DataListSource = "CUSTOM"
	REPLACE_DOWNLOAD_NativeAnchorGetV30DataListSource NativeAnchorGetV30DataListSource = "REPLACE_DOWNLOAD"
)

// All allowed values of NativeAnchorGetV30DataListSource enum
var AllowedNativeAnchorGetV30DataListSourceEnumValues = []NativeAnchorGetV30DataListSource{
	"AUTO",
	"CUSTOM",
	"REPLACE_DOWNLOAD",
}

// NewNativeAnchorGetV30DataListSourceFromValue returns a pointer to a valid NativeAnchorGetV30DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorGetV30DataListSourceFromValue(v string) (*NativeAnchorGetV30DataListSource, error) {
	ev := NativeAnchorGetV30DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorGetV30DataListSource: valid values are %v", v, AllowedNativeAnchorGetV30DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorGetV30DataListSource) IsValid() bool {
	for _, existing := range AllowedNativeAnchorGetV30DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_get_v3.0_data_list_source value
func (v NativeAnchorGetV30DataListSource) Ptr() *NativeAnchorGetV30DataListSource {
	return &v
}
