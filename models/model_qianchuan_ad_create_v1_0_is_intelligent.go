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

// QianchuanAdCreateV10IsIntelligent
type QianchuanAdCreateV10IsIntelligent int64

// List of qianchuan_ad_create_v1.0_is_intelligent
const (
	Enum_0_QianchuanAdCreateV10IsIntelligent QianchuanAdCreateV10IsIntelligent = 0
	Enum_1_QianchuanAdCreateV10IsIntelligent QianchuanAdCreateV10IsIntelligent = 1
)

// All allowed values of QianchuanAdCreateV10IsIntelligent enum
var AllowedQianchuanAdCreateV10IsIntelligentEnumValues = []QianchuanAdCreateV10IsIntelligent{
	0,
	1,
}

// NewQianchuanAdCreateV10IsIntelligentFromValue returns a pointer to a valid QianchuanAdCreateV10IsIntelligent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10IsIntelligentFromValue(v int64) (*QianchuanAdCreateV10IsIntelligent, error) {
	ev := QianchuanAdCreateV10IsIntelligent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10IsIntelligent: valid values are %v", v, AllowedQianchuanAdCreateV10IsIntelligentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10IsIntelligent) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10IsIntelligentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_is_intelligent value
func (v QianchuanAdCreateV10IsIntelligent) Ptr() *QianchuanAdCreateV10IsIntelligent {
	return &v
}
