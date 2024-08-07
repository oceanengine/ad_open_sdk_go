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

// EnterpriseOverviewDataGetV10TimeDimension
type EnterpriseOverviewDataGetV10TimeDimension string

// List of enterprise_overview_data_get_v1.0_time_dimension
const (
	SUM_EnterpriseOverviewDataGetV10TimeDimension   EnterpriseOverviewDataGetV10TimeDimension = "SUM"
	DAILY_EnterpriseOverviewDataGetV10TimeDimension EnterpriseOverviewDataGetV10TimeDimension = "DAILY"
)

// All allowed values of EnterpriseOverviewDataGetV10TimeDimension enum
var AllowedEnterpriseOverviewDataGetV10TimeDimensionEnumValues = []EnterpriseOverviewDataGetV10TimeDimension{
	"SUM",
	"DAILY",
}

// NewEnterpriseOverviewDataGetV10TimeDimensionFromValue returns a pointer to a valid EnterpriseOverviewDataGetV10TimeDimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseOverviewDataGetV10TimeDimensionFromValue(v string) (*EnterpriseOverviewDataGetV10TimeDimension, error) {
	ev := EnterpriseOverviewDataGetV10TimeDimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseOverviewDataGetV10TimeDimension: valid values are %v", v, AllowedEnterpriseOverviewDataGetV10TimeDimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseOverviewDataGetV10TimeDimension) IsValid() bool {
	for _, existing := range AllowedEnterpriseOverviewDataGetV10TimeDimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to enterprise_overview_data_get_v1.0_time_dimension value
func (v EnterpriseOverviewDataGetV10TimeDimension) Ptr() *EnterpriseOverviewDataGetV10TimeDimension {
	return &v
}
