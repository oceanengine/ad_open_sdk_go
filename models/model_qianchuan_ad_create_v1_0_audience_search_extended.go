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

// QianchuanAdCreateV10AudienceSearchExtended
type QianchuanAdCreateV10AudienceSearchExtended int64

// List of qianchuan_ad_create_v1.0_audience_search_extended
const (
	Enum_1_QianchuanAdCreateV10AudienceSearchExtended QianchuanAdCreateV10AudienceSearchExtended = 1
	Enum_2_QianchuanAdCreateV10AudienceSearchExtended QianchuanAdCreateV10AudienceSearchExtended = 2
)

// All allowed values of QianchuanAdCreateV10AudienceSearchExtended enum
var AllowedQianchuanAdCreateV10AudienceSearchExtendedEnumValues = []QianchuanAdCreateV10AudienceSearchExtended{
	1,
	2,
}

// NewQianchuanAdCreateV10AudienceSearchExtendedFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceSearchExtended
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceSearchExtendedFromValue(v int64) (*QianchuanAdCreateV10AudienceSearchExtended, error) {
	ev := QianchuanAdCreateV10AudienceSearchExtended(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceSearchExtended: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceSearchExtendedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceSearchExtended) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceSearchExtendedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_search_extended value
func (v QianchuanAdCreateV10AudienceSearchExtended) Ptr() *QianchuanAdCreateV10AudienceSearchExtended {
	return &v
}
