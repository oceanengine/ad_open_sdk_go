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

// QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays
type QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays string

// List of qianchuan_ad_create_v1.0_audience_aweme_fan_behaviors_days
const (
	DAYS_15_QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays = "DAYS_15"
	DAYS_30_QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays = "DAYS_30"
	DAYS_60_QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays = "DAYS_60"
)

// All allowed values of QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays enum
var AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsDaysEnumValues = []QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays{
	"DAYS_15",
	"DAYS_30",
	"DAYS_60",
}

// NewQianchuanAdCreateV10AudienceAwemeFanBehaviorsDaysFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAwemeFanBehaviorsDaysFromValue(v string) (*QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays, error) {
	ev := QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_aweme_fan_behaviors_days value
func (v QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays) Ptr() *QianchuanAdCreateV10AudienceAwemeFanBehaviorsDays {
	return &v
}
