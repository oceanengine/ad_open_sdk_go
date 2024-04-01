/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderCreateV30TargetingAudienceMode
type DouplusOrderCreateV30TargetingAudienceMode string

// List of douplus_order_create_v3.0_targeting_audience_mode
const (
	AUTO_DouplusOrderCreateV30TargetingAudienceMode         DouplusOrderCreateV30TargetingAudienceMode = "AUTO"
	CUSTOM_DouplusOrderCreateV30TargetingAudienceMode       DouplusOrderCreateV30TargetingAudienceMode = "CUSTOM"
	AUDIENCE_PKG_DouplusOrderCreateV30TargetingAudienceMode DouplusOrderCreateV30TargetingAudienceMode = "AUDIENCE_PKG"
	DMP_PKG_DouplusOrderCreateV30TargetingAudienceMode      DouplusOrderCreateV30TargetingAudienceMode = "DMP_PKG"
)

// All allowed values of DouplusOrderCreateV30TargetingAudienceMode enum
var AllowedDouplusOrderCreateV30TargetingAudienceModeEnumValues = []DouplusOrderCreateV30TargetingAudienceMode{
	"AUTO",
	"CUSTOM",
	"AUDIENCE_PKG",
	"DMP_PKG",
}

// NewDouplusOrderCreateV30TargetingAudienceModeFromValue returns a pointer to a valid DouplusOrderCreateV30TargetingAudienceMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30TargetingAudienceModeFromValue(v string) (*DouplusOrderCreateV30TargetingAudienceMode, error) {
	ev := DouplusOrderCreateV30TargetingAudienceMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30TargetingAudienceMode: valid values are %v", v, AllowedDouplusOrderCreateV30TargetingAudienceModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30TargetingAudienceMode) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30TargetingAudienceModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_targeting_audience_mode value
func (v DouplusOrderCreateV30TargetingAudienceMode) Ptr() *DouplusOrderCreateV30TargetingAudienceMode {
	return &v
}
