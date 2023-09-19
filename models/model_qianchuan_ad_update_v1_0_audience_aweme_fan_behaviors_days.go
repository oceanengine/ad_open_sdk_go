/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays
type QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays string

// List of qianchuan_ad_update_v1.0_audience_aweme_fan_behaviors_days
const (
	DAYS_15_QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays = "DAYS_15"
	DAYS_30_QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays = "DAYS_30"
	DAYS_60_QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays = "DAYS_60"
)

// All allowed values of QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays enum
var AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsDaysEnumValues = []QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays{
	"DAYS_15",
	"DAYS_30",
	"DAYS_60",
}

// NewQianchuanAdUpdateV10AudienceAwemeFanBehaviorsDaysFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAwemeFanBehaviorsDaysFromValue(v string) (*QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays, error) {
	ev := QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_aweme_fan_behaviors_days value
func (v QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays) Ptr() *QianchuanAdUpdateV10AudienceAwemeFanBehaviorsDays {
	return &v
}
