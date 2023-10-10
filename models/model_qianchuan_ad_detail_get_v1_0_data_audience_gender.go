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

// QianchuanAdDetailGetV10DataAudienceGender
type QianchuanAdDetailGetV10DataAudienceGender string

// List of qianchuan_ad_detail_get_v1.0_data_audience_gender
const (
	GENDER_FEMALE_QianchuanAdDetailGetV10DataAudienceGender QianchuanAdDetailGetV10DataAudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanAdDetailGetV10DataAudienceGender   QianchuanAdDetailGetV10DataAudienceGender = "GENDER_MALE"
	NONE_QianchuanAdDetailGetV10DataAudienceGender          QianchuanAdDetailGetV10DataAudienceGender = "NONE"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceGender enum
var AllowedQianchuanAdDetailGetV10DataAudienceGenderEnumValues = []QianchuanAdDetailGetV10DataAudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewQianchuanAdDetailGetV10DataAudienceGenderFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceGenderFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceGender, error) {
	ev := QianchuanAdDetailGetV10DataAudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceGender: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_gender value
func (v QianchuanAdDetailGetV10DataAudienceGender) Ptr() *QianchuanAdDetailGetV10DataAudienceGender {
	return &v
}
