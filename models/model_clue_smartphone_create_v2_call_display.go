/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueSmartphoneCreateV2CallDisplay
type ClueSmartphoneCreateV2CallDisplay int64

// List of clue_smartphone_create_v2_call_display
const (
	Enum_0_ClueSmartphoneCreateV2CallDisplay ClueSmartphoneCreateV2CallDisplay = 0
	Enum_1_ClueSmartphoneCreateV2CallDisplay ClueSmartphoneCreateV2CallDisplay = 1
)

// All allowed values of ClueSmartphoneCreateV2CallDisplay enum
var AllowedClueSmartphoneCreateV2CallDisplayEnumValues = []ClueSmartphoneCreateV2CallDisplay{
	0,
	1,
}

// NewClueSmartphoneCreateV2CallDisplayFromValue returns a pointer to a valid ClueSmartphoneCreateV2CallDisplay
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueSmartphoneCreateV2CallDisplayFromValue(v int64) (*ClueSmartphoneCreateV2CallDisplay, error) {
	ev := ClueSmartphoneCreateV2CallDisplay(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueSmartphoneCreateV2CallDisplay: valid values are %v", v, AllowedClueSmartphoneCreateV2CallDisplayEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueSmartphoneCreateV2CallDisplay) IsValid() bool {
	for _, existing := range AllowedClueSmartphoneCreateV2CallDisplayEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_smartphone_create_v2_call_display value
func (v ClueSmartphoneCreateV2CallDisplay) Ptr() *ClueSmartphoneCreateV2CallDisplay {
	return &v
}
