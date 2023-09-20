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

// DmpCustomAudienceSelectV2SelectType
type DmpCustomAudienceSelectV2SelectType int64

// List of dmp_custom_audience_select_v2_select_type
const (
	Enum_0_DmpCustomAudienceSelectV2SelectType DmpCustomAudienceSelectV2SelectType = 0
	Enum_1_DmpCustomAudienceSelectV2SelectType DmpCustomAudienceSelectV2SelectType = 1
)

// All allowed values of DmpCustomAudienceSelectV2SelectType enum
var AllowedDmpCustomAudienceSelectV2SelectTypeEnumValues = []DmpCustomAudienceSelectV2SelectType{
	0,
	1,
}

// NewDmpCustomAudienceSelectV2SelectTypeFromValue returns a pointer to a valid DmpCustomAudienceSelectV2SelectType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDmpCustomAudienceSelectV2SelectTypeFromValue(v int64) (*DmpCustomAudienceSelectV2SelectType, error) {
	ev := DmpCustomAudienceSelectV2SelectType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DmpCustomAudienceSelectV2SelectType: valid values are %v", v, AllowedDmpCustomAudienceSelectV2SelectTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DmpCustomAudienceSelectV2SelectType) IsValid() bool {
	for _, existing := range AllowedDmpCustomAudienceSelectV2SelectTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dmp_custom_audience_select_v2_select_type value
func (v DmpCustomAudienceSelectV2SelectType) Ptr() *DmpCustomAudienceSelectV2SelectType {
	return &v
}
