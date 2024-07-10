/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryRebateBalanceV2FilteringStandard
type QueryRebateBalanceV2FilteringStandard string

// List of query_rebate_balance_v2_filtering_standard
const (
	YES_QueryRebateBalanceV2FilteringStandard QueryRebateBalanceV2FilteringStandard = "YES"
	NO_QueryRebateBalanceV2FilteringStandard  QueryRebateBalanceV2FilteringStandard = "NO"
)

// All allowed values of QueryRebateBalanceV2FilteringStandard enum
var AllowedQueryRebateBalanceV2FilteringStandardEnumValues = []QueryRebateBalanceV2FilteringStandard{
	"YES",
	"NO",
}

// NewQueryRebateBalanceV2FilteringStandardFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringStandard
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringStandardFromValue(v string) (*QueryRebateBalanceV2FilteringStandard, error) {
	ev := QueryRebateBalanceV2FilteringStandard(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringStandard: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringStandardEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringStandard) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringStandardEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_standard value
func (v QueryRebateBalanceV2FilteringStandard) Ptr() *QueryRebateBalanceV2FilteringStandard {
	return &v
}
