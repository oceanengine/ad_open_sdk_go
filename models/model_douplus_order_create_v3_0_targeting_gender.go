/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderCreateV30TargetingGender
type DouplusOrderCreateV30TargetingGender string

// List of douplus_order_create_v3.0_targeting_gender
const (
	GENDER_FEMALE_DouplusOrderCreateV30TargetingGender DouplusOrderCreateV30TargetingGender = "GENDER_FEMALE"
	GENDER_MALE_DouplusOrderCreateV30TargetingGender   DouplusOrderCreateV30TargetingGender = "GENDER_MALE"
)

// All allowed values of DouplusOrderCreateV30TargetingGender enum
var AllowedDouplusOrderCreateV30TargetingGenderEnumValues = []DouplusOrderCreateV30TargetingGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
}

// NewDouplusOrderCreateV30TargetingGenderFromValue returns a pointer to a valid DouplusOrderCreateV30TargetingGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30TargetingGenderFromValue(v string) (*DouplusOrderCreateV30TargetingGender, error) {
	ev := DouplusOrderCreateV30TargetingGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30TargetingGender: valid values are %v", v, AllowedDouplusOrderCreateV30TargetingGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30TargetingGender) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30TargetingGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_targeting_gender value
func (v DouplusOrderCreateV30TargetingGender) Ptr() *DouplusOrderCreateV30TargetingGender {
	return &v
}
