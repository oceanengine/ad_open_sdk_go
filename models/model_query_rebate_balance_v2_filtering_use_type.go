/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QueryRebateBalanceV2FilteringUseType
type QueryRebateBalanceV2FilteringUseType string

// List of query_rebate_balance_v2_filtering_use_type
const (
	CASH_QueryRebateBalanceV2FilteringUseType       QueryRebateBalanceV2FilteringUseType = "CASH"
	CHARGE_QueryRebateBalanceV2FilteringUseType     QueryRebateBalanceV2FilteringUseType = "CHARGE"
	HEDGING_QueryRebateBalanceV2FilteringUseType    QueryRebateBalanceV2FilteringUseType = "HEDGING"
	NONPAYMENT_QueryRebateBalanceV2FilteringUseType QueryRebateBalanceV2FilteringUseType = "NONPAYMENT"
	REVERT_QueryRebateBalanceV2FilteringUseType     QueryRebateBalanceV2FilteringUseType = "REVERT"
)

// All allowed values of QueryRebateBalanceV2FilteringUseType enum
var AllowedQueryRebateBalanceV2FilteringUseTypeEnumValues = []QueryRebateBalanceV2FilteringUseType{
	"CASH",
	"CHARGE",
	"HEDGING",
	"NONPAYMENT",
	"REVERT",
}

// NewQueryRebateBalanceV2FilteringUseTypeFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringUseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringUseTypeFromValue(v string) (*QueryRebateBalanceV2FilteringUseType, error) {
	ev := QueryRebateBalanceV2FilteringUseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringUseType: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringUseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringUseType) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringUseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_use_type value
func (v QueryRebateBalanceV2FilteringUseType) Ptr() *QueryRebateBalanceV2FilteringUseType {
	return &v
}