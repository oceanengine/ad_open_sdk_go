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

// CreativeDetailGetV30DataAdDataEnableSmartSource
type CreativeDetailGetV30DataAdDataEnableSmartSource string

// List of creative_detail_get_v3.0_data_ad_data_enable_smart_source
const (
	OFF_CreativeDetailGetV30DataAdDataEnableSmartSource CreativeDetailGetV30DataAdDataEnableSmartSource = "OFF"
	ON_CreativeDetailGetV30DataAdDataEnableSmartSource  CreativeDetailGetV30DataAdDataEnableSmartSource = "ON"
)

// All allowed values of CreativeDetailGetV30DataAdDataEnableSmartSource enum
var AllowedCreativeDetailGetV30DataAdDataEnableSmartSourceEnumValues = []CreativeDetailGetV30DataAdDataEnableSmartSource{
	"OFF",
	"ON",
}

// NewCreativeDetailGetV30DataAdDataEnableSmartSourceFromValue returns a pointer to a valid CreativeDetailGetV30DataAdDataEnableSmartSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeDetailGetV30DataAdDataEnableSmartSourceFromValue(v string) (*CreativeDetailGetV30DataAdDataEnableSmartSource, error) {
	ev := CreativeDetailGetV30DataAdDataEnableSmartSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeDetailGetV30DataAdDataEnableSmartSource: valid values are %v", v, AllowedCreativeDetailGetV30DataAdDataEnableSmartSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeDetailGetV30DataAdDataEnableSmartSource) IsValid() bool {
	for _, existing := range AllowedCreativeDetailGetV30DataAdDataEnableSmartSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_detail_get_v3.0_data_ad_data_enable_smart_source value
func (v CreativeDetailGetV30DataAdDataEnableSmartSource) Ptr() *CreativeDetailGetV30DataAdDataEnableSmartSource {
	return &v
}
