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

// EnterpriseOverviewDataGetV10OrderType
type EnterpriseOverviewDataGetV10OrderType string

// List of enterprise_overview_data_get_v1.0_order_type
const (
	ASC_EnterpriseOverviewDataGetV10OrderType  EnterpriseOverviewDataGetV10OrderType = "ASC"
	DESC_EnterpriseOverviewDataGetV10OrderType EnterpriseOverviewDataGetV10OrderType = "DESC"
)

// All allowed values of EnterpriseOverviewDataGetV10OrderType enum
var AllowedEnterpriseOverviewDataGetV10OrderTypeEnumValues = []EnterpriseOverviewDataGetV10OrderType{
	"ASC",
	"DESC",
}

// NewEnterpriseOverviewDataGetV10OrderTypeFromValue returns a pointer to a valid EnterpriseOverviewDataGetV10OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseOverviewDataGetV10OrderTypeFromValue(v string) (*EnterpriseOverviewDataGetV10OrderType, error) {
	ev := EnterpriseOverviewDataGetV10OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseOverviewDataGetV10OrderType: valid values are %v", v, AllowedEnterpriseOverviewDataGetV10OrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseOverviewDataGetV10OrderType) IsValid() bool {
	for _, existing := range AllowedEnterpriseOverviewDataGetV10OrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_overview_data_get_v1.0_order_type value
func (v EnterpriseOverviewDataGetV10OrderType) Ptr() *EnterpriseOverviewDataGetV10OrderType {
	return &v
}
