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

// QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays
type QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays string

// List of qianchuan_tools_estimate_audience_v1.0_audience_aweme_fan_behaviors_days
const (
	DAYS_15_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays = "DAYS_15"
	DAYS_30_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays = "DAYS_30"
	DAYS_60_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays = "DAYS_60"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDaysEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays{
	"DAYS_15",
	"DAYS_30",
	"DAYS_60",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDaysFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDaysFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_aweme_fan_behaviors_days value
func (v QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsDays {
	return &v
}
