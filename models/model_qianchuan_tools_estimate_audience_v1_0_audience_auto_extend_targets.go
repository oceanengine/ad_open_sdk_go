/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets
type QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets string

// List of qianchuan_tools_estimate_audience_v1.0_audience_auto_extend_targets
const (
	AGE_QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets             QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets          QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets = "GENDER"
	INTEREST_ACTION_QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets = "INTEREST_ACTION"
	REGION_QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets          QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets = "REGION"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendTargetsEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets{
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_ACTION",
	"REGION",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAutoExtendTargetsFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAutoExtendTargetsFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_auto_extend_targets value
func (v QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAutoExtendTargets {
	return &v
}
