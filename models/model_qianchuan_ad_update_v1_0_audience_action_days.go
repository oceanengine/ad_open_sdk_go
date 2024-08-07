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

// QianchuanAdUpdateV10AudienceActionDays
type QianchuanAdUpdateV10AudienceActionDays int64

// List of qianchuan_ad_update_v1.0_audience_action_days
const (
	Enum_7_QianchuanAdUpdateV10AudienceActionDays   QianchuanAdUpdateV10AudienceActionDays = 7
	Enum_15_QianchuanAdUpdateV10AudienceActionDays  QianchuanAdUpdateV10AudienceActionDays = 15
	Enum_30_QianchuanAdUpdateV10AudienceActionDays  QianchuanAdUpdateV10AudienceActionDays = 30
	Enum_60_QianchuanAdUpdateV10AudienceActionDays  QianchuanAdUpdateV10AudienceActionDays = 60
	Enum_90_QianchuanAdUpdateV10AudienceActionDays  QianchuanAdUpdateV10AudienceActionDays = 90
	Enum_180_QianchuanAdUpdateV10AudienceActionDays QianchuanAdUpdateV10AudienceActionDays = 180
	Enum_365_QianchuanAdUpdateV10AudienceActionDays QianchuanAdUpdateV10AudienceActionDays = 365
)

// All allowed values of QianchuanAdUpdateV10AudienceActionDays enum
var AllowedQianchuanAdUpdateV10AudienceActionDaysEnumValues = []QianchuanAdUpdateV10AudienceActionDays{
	7,
	15,
	30,
	60,
	90,
	180,
	365,
}

// NewQianchuanAdUpdateV10AudienceActionDaysFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceActionDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceActionDaysFromValue(v int64) (*QianchuanAdUpdateV10AudienceActionDays, error) {
	ev := QianchuanAdUpdateV10AudienceActionDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceActionDays: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceActionDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceActionDays) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceActionDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_action_days value
func (v QianchuanAdUpdateV10AudienceActionDays) Ptr() *QianchuanAdUpdateV10AudienceActionDays {
	return &v
}
