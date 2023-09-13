/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataIsHomepageHide
type QianchuanAdDetailGetV10DataIsHomepageHide int64

// List of qianchuan_ad_detail_get_v1.0_data_is_homepage_hide
const (
	Enum_0_QianchuanAdDetailGetV10DataIsHomepageHide QianchuanAdDetailGetV10DataIsHomepageHide = 0
	Enum_1_QianchuanAdDetailGetV10DataIsHomepageHide QianchuanAdDetailGetV10DataIsHomepageHide = 1
)

// All allowed values of QianchuanAdDetailGetV10DataIsHomepageHide enum
var AllowedQianchuanAdDetailGetV10DataIsHomepageHideEnumValues = []QianchuanAdDetailGetV10DataIsHomepageHide{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataIsHomepageHideFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataIsHomepageHide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataIsHomepageHideFromValue(v int64) (*QianchuanAdDetailGetV10DataIsHomepageHide, error) {
	ev := QianchuanAdDetailGetV10DataIsHomepageHide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataIsHomepageHide: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataIsHomepageHideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataIsHomepageHide) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataIsHomepageHideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_is_homepage_hide value
func (v QianchuanAdDetailGetV10DataIsHomepageHide) Ptr() *QianchuanAdDetailGetV10DataIsHomepageHide {
	return &v
}
