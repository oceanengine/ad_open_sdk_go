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

// QianchuanAdUpdateV10AudienceAge
type QianchuanAdUpdateV10AudienceAge string

// List of qianchuan_ad_update_v1.0_audience_age
const (
	AGE_ABOVE_50_QianchuanAdUpdateV10AudienceAge      QianchuanAdUpdateV10AudienceAge = "AGE_ABOVE_50"
	AGE_BETWEEN_18_23_QianchuanAdUpdateV10AudienceAge QianchuanAdUpdateV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAdUpdateV10AudienceAge QianchuanAdUpdateV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAdUpdateV10AudienceAge QianchuanAdUpdateV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_QianchuanAdUpdateV10AudienceAge QianchuanAdUpdateV10AudienceAge = "AGE_BETWEEN_41_49"
)

// All allowed values of QianchuanAdUpdateV10AudienceAge enum
var AllowedQianchuanAdUpdateV10AudienceAgeEnumValues = []QianchuanAdUpdateV10AudienceAge{
	"AGE_ABOVE_50",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewQianchuanAdUpdateV10AudienceAgeFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAgeFromValue(v string) (*QianchuanAdUpdateV10AudienceAge, error) {
	ev := QianchuanAdUpdateV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAge: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_age value
func (v QianchuanAdUpdateV10AudienceAge) Ptr() *QianchuanAdUpdateV10AudienceAge {
	return &v
}
