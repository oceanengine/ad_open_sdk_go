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

// QianchuanAdDetailGetV10DataAudienceAutoExtendTargets
type QianchuanAdDetailGetV10DataAudienceAutoExtendTargets string

// List of qianchuan_ad_detail_get_v1.0_data_audience_auto_extend_targets
const (
	AD_TAG_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets          QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "AD_TAG"
	AGE_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets             QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets          QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "INTEREST_ACTION"
	INTEREST_TAG_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets    QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_QianchuanAdDetailGetV10DataAudienceAutoExtendTargets          QianchuanAdDetailGetV10DataAudienceAutoExtendTargets = "REGION"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAutoExtendTargets enum
var AllowedQianchuanAdDetailGetV10DataAudienceAutoExtendTargetsEnumValues = []QianchuanAdDetailGetV10DataAudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"INTEREST_TAG",
	"REGION",
}

// NewQianchuanAdDetailGetV10DataAudienceAutoExtendTargetsFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAutoExtendTargetsFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAutoExtendTargets, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAutoExtendTargets: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_auto_extend_targets value
func (v QianchuanAdDetailGetV10DataAudienceAutoExtendTargets) Ptr() *QianchuanAdDetailGetV10DataAudienceAutoExtendTargets {
	return &v
}
