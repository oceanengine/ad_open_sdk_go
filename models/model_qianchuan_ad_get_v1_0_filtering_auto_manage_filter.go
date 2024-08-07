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

// QianchuanAdGetV10FilteringAutoManageFilter
type QianchuanAdGetV10FilteringAutoManageFilter string

// List of qianchuan_ad_get_v1.0_filtering_auto_manage_filter
const (
	ALL_QianchuanAdGetV10FilteringAutoManageFilter         QianchuanAdGetV10FilteringAutoManageFilter = "ALL"
	AUTO_MANAGE_QianchuanAdGetV10FilteringAutoManageFilter QianchuanAdGetV10FilteringAutoManageFilter = "AUTO_MANAGE"
	NORMAL_QianchuanAdGetV10FilteringAutoManageFilter      QianchuanAdGetV10FilteringAutoManageFilter = "NORMAL"
)

// All allowed values of QianchuanAdGetV10FilteringAutoManageFilter enum
var AllowedQianchuanAdGetV10FilteringAutoManageFilterEnumValues = []QianchuanAdGetV10FilteringAutoManageFilter{
	"ALL",
	"AUTO_MANAGE",
	"NORMAL",
}

// NewQianchuanAdGetV10FilteringAutoManageFilterFromValue returns a pointer to a valid QianchuanAdGetV10FilteringAutoManageFilter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10FilteringAutoManageFilterFromValue(v string) (*QianchuanAdGetV10FilteringAutoManageFilter, error) {
	ev := QianchuanAdGetV10FilteringAutoManageFilter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10FilteringAutoManageFilter: valid values are %v", v, AllowedQianchuanAdGetV10FilteringAutoManageFilterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10FilteringAutoManageFilter) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10FilteringAutoManageFilterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_filtering_auto_manage_filter value
func (v QianchuanAdGetV10FilteringAutoManageFilter) Ptr() *QianchuanAdGetV10FilteringAutoManageFilter {
	return &v
}
