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

// CreativeProceduralCreativeCreateV2AdDataEnableSmartSource
type CreativeProceduralCreativeCreateV2AdDataEnableSmartSource string

// List of creative_procedural_creative_create_v2_ad_data_enable_smart_source
const (
	OFF_CreativeProceduralCreativeCreateV2AdDataEnableSmartSource CreativeProceduralCreativeCreateV2AdDataEnableSmartSource = "OFF"
	ON_CreativeProceduralCreativeCreateV2AdDataEnableSmartSource  CreativeProceduralCreativeCreateV2AdDataEnableSmartSource = "ON"
)

// All allowed values of CreativeProceduralCreativeCreateV2AdDataEnableSmartSource enum
var AllowedCreativeProceduralCreativeCreateV2AdDataEnableSmartSourceEnumValues = []CreativeProceduralCreativeCreateV2AdDataEnableSmartSource{
	"OFF",
	"ON",
}

// NewCreativeProceduralCreativeCreateV2AdDataEnableSmartSourceFromValue returns a pointer to a valid CreativeProceduralCreativeCreateV2AdDataEnableSmartSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeCreateV2AdDataEnableSmartSourceFromValue(v string) (*CreativeProceduralCreativeCreateV2AdDataEnableSmartSource, error) {
	ev := CreativeProceduralCreativeCreateV2AdDataEnableSmartSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeCreateV2AdDataEnableSmartSource: valid values are %v", v, AllowedCreativeProceduralCreativeCreateV2AdDataEnableSmartSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeCreateV2AdDataEnableSmartSource) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeCreateV2AdDataEnableSmartSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_create_v2_ad_data_enable_smart_source value
func (v CreativeProceduralCreativeCreateV2AdDataEnableSmartSource) Ptr() *CreativeProceduralCreativeCreateV2AdDataEnableSmartSource {
	return &v
}
