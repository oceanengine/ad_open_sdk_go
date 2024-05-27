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

// QianchuanAdDetailGetV10DataAudienceSearchExtended
type QianchuanAdDetailGetV10DataAudienceSearchExtended int64

// List of qianchuan_ad_detail_get_v1.0_data_audience_search_extended
const (
	Enum_1_QianchuanAdDetailGetV10DataAudienceSearchExtended QianchuanAdDetailGetV10DataAudienceSearchExtended = 1
	Enum_2_QianchuanAdDetailGetV10DataAudienceSearchExtended QianchuanAdDetailGetV10DataAudienceSearchExtended = 2
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceSearchExtended enum
var AllowedQianchuanAdDetailGetV10DataAudienceSearchExtendedEnumValues = []QianchuanAdDetailGetV10DataAudienceSearchExtended{
	1,
	2,
}

// NewQianchuanAdDetailGetV10DataAudienceSearchExtendedFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceSearchExtended
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceSearchExtendedFromValue(v int64) (*QianchuanAdDetailGetV10DataAudienceSearchExtended, error) {
	ev := QianchuanAdDetailGetV10DataAudienceSearchExtended(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceSearchExtended: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceSearchExtendedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceSearchExtended) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceSearchExtendedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_search_extended value
func (v QianchuanAdDetailGetV10DataAudienceSearchExtended) Ptr() *QianchuanAdDetailGetV10DataAudienceSearchExtended {
	return &v
}
