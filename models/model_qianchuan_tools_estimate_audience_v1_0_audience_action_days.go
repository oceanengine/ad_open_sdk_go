/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceActionDays
type QianchuanToolsEstimateAudienceV10AudienceActionDays int64

// List of qianchuan_tools_estimate_audience_v1.0_audience_action_days
const (
	Enum_7_QianchuanToolsEstimateAudienceV10AudienceActionDays   QianchuanToolsEstimateAudienceV10AudienceActionDays = 7
	Enum_15_QianchuanToolsEstimateAudienceV10AudienceActionDays  QianchuanToolsEstimateAudienceV10AudienceActionDays = 15
	Enum_30_QianchuanToolsEstimateAudienceV10AudienceActionDays  QianchuanToolsEstimateAudienceV10AudienceActionDays = 30
	Enum_60_QianchuanToolsEstimateAudienceV10AudienceActionDays  QianchuanToolsEstimateAudienceV10AudienceActionDays = 60
	Enum_90_QianchuanToolsEstimateAudienceV10AudienceActionDays  QianchuanToolsEstimateAudienceV10AudienceActionDays = 90
	Enum_180_QianchuanToolsEstimateAudienceV10AudienceActionDays QianchuanToolsEstimateAudienceV10AudienceActionDays = 180
	Enum_365_QianchuanToolsEstimateAudienceV10AudienceActionDays QianchuanToolsEstimateAudienceV10AudienceActionDays = 365
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceActionDays enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceActionDaysEnumValues = []QianchuanToolsEstimateAudienceV10AudienceActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewQianchuanToolsEstimateAudienceV10AudienceActionDaysFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceActionDaysFromValue(v int64) (*QianchuanToolsEstimateAudienceV10AudienceActionDays, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceActionDays: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceActionDays) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_action_days value
func (v QianchuanToolsEstimateAudienceV10AudienceActionDays) Ptr() *QianchuanToolsEstimateAudienceV10AudienceActionDays {
	return &v
}
