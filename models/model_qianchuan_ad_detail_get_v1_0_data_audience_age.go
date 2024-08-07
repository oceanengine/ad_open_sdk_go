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

// QianchuanAdDetailGetV10DataAudienceAge
type QianchuanAdDetailGetV10DataAudienceAge string

// List of qianchuan_ad_detail_get_v1.0_data_audience_age
const (
	AGE_ABOVE_50_QianchuanAdDetailGetV10DataAudienceAge      QianchuanAdDetailGetV10DataAudienceAge = "AGE_ABOVE_50"
	AGE_BETWEEN_18_23_QianchuanAdDetailGetV10DataAudienceAge QianchuanAdDetailGetV10DataAudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAdDetailGetV10DataAudienceAge QianchuanAdDetailGetV10DataAudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAdDetailGetV10DataAudienceAge QianchuanAdDetailGetV10DataAudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_49_QianchuanAdDetailGetV10DataAudienceAge QianchuanAdDetailGetV10DataAudienceAge = "AGE_BETWEEN_41_49"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAge enum
var AllowedQianchuanAdDetailGetV10DataAudienceAgeEnumValues = []QianchuanAdDetailGetV10DataAudienceAge{
	"AGE_ABOVE_50",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_49",
}

// NewQianchuanAdDetailGetV10DataAudienceAgeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAgeFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAge, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAge: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_age value
func (v QianchuanAdDetailGetV10DataAudienceAge) Ptr() *QianchuanAdDetailGetV10DataAudienceAge {
	return &v
}
