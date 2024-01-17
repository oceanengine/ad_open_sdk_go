/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceGender
type QianchuanToolsEstimateAudienceV10AudienceGender string

// List of qianchuan_tools_estimate_audience_v1.0_audience_gender
const (
	GENDER_FEMALE_QianchuanToolsEstimateAudienceV10AudienceGender QianchuanToolsEstimateAudienceV10AudienceGender = "GENDER_FEMALE"
	GENDER_MALE_QianchuanToolsEstimateAudienceV10AudienceGender   QianchuanToolsEstimateAudienceV10AudienceGender = "GENDER_MALE"
	NONE_QianchuanToolsEstimateAudienceV10AudienceGender          QianchuanToolsEstimateAudienceV10AudienceGender = "NONE"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceGender enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceGenderEnumValues = []QianchuanToolsEstimateAudienceV10AudienceGender{
	"GENDER_FEMALE",
	"GENDER_MALE",
	"NONE",
}

// NewQianchuanToolsEstimateAudienceV10AudienceGenderFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceGenderFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceGender, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceGender: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceGender) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_gender value
func (v QianchuanToolsEstimateAudienceV10AudienceGender) Ptr() *QianchuanToolsEstimateAudienceV10AudienceGender {
	return &v
}
