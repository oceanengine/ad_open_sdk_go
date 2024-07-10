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

// QianchuanAdCreateV10AudienceAge
type QianchuanAdCreateV10AudienceAge string

// List of qianchuan_ad_create_v1.0_audience_age
const (
	AGE_ABOVE_50_QianchuanAdCreateV10AudienceAge      QianchuanAdCreateV10AudienceAge = "AGE_ABOVE_50"
	AGE_BETWEEN_18_23_QianchuanAdCreateV10AudienceAge QianchuanAdCreateV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAdCreateV10AudienceAge QianchuanAdCreateV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAdCreateV10AudienceAge QianchuanAdCreateV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_QianchuanAdCreateV10AudienceAge QianchuanAdCreateV10AudienceAge = "AGE_BETWEEN_41_49"
)

// All allowed values of QianchuanAdCreateV10AudienceAge enum
var AllowedQianchuanAdCreateV10AudienceAgeEnumValues = []QianchuanAdCreateV10AudienceAge{
	"AGE_ABOVE_50",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewQianchuanAdCreateV10AudienceAgeFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAgeFromValue(v string) (*QianchuanAdCreateV10AudienceAge, error) {
	ev := QianchuanAdCreateV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAge: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_age value
func (v QianchuanAdCreateV10AudienceAge) Ptr() *QianchuanAdCreateV10AudienceAge {
	return &v
}
