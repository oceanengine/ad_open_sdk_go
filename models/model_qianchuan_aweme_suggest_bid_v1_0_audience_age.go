/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeSuggestBidV10AudienceAge
type QianchuanAwemeSuggestBidV10AudienceAge string

// List of qianchuan_aweme_suggest_bid_v1.0_audience_age
const (
	AGE_BETWEEN_18_23_QianchuanAwemeSuggestBidV10AudienceAge QianchuanAwemeSuggestBidV10AudienceAge = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_24_30_QianchuanAwemeSuggestBidV10AudienceAge QianchuanAwemeSuggestBidV10AudienceAge = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_31_40_QianchuanAwemeSuggestBidV10AudienceAge QianchuanAwemeSuggestBidV10AudienceAge = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_41_50_QianchuanAwemeSuggestBidV10AudienceAge QianchuanAwemeSuggestBidV10AudienceAge = "AGE_BETWEEN_41_50"
)

// All allowed values of QianchuanAwemeSuggestBidV10AudienceAge enum
var AllowedQianchuanAwemeSuggestBidV10AudienceAgeEnumValues = []QianchuanAwemeSuggestBidV10AudienceAge{
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_41_50",
}

// NewQianchuanAwemeSuggestBidV10AudienceAgeFromValue returns a pointer to a valid QianchuanAwemeSuggestBidV10AudienceAge
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeSuggestBidV10AudienceAgeFromValue(v string) (*QianchuanAwemeSuggestBidV10AudienceAge, error) {
	ev := QianchuanAwemeSuggestBidV10AudienceAge(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeSuggestBidV10AudienceAge: valid values are %v", v, AllowedQianchuanAwemeSuggestBidV10AudienceAgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeSuggestBidV10AudienceAge) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeSuggestBidV10AudienceAgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_suggest_bid_v1.0_audience_age value
func (v QianchuanAwemeSuggestBidV10AudienceAge) Ptr() *QianchuanAwemeSuggestBidV10AudienceAge {
	return &v
}
