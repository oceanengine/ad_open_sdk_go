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

// QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays
type QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays string

// List of qianchuan_ad_detail_get_v1.0_data_audience_aweme_fan_behaviors_days
const (
	DAYS_15_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays = "DAYS_15"
	DAYS_30_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays = "DAYS_30"
	DAYS_60_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays = "DAYS_60"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays enum
var AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDaysEnumValues = []QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays{
	"DAYS_15",
	"DAYS_30",
	"DAYS_60",
}

// NewQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDaysFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDaysFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDaysEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDaysEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_aweme_fan_behaviors_days value
func (v QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays) Ptr() *QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsDays {
	return &v
}
