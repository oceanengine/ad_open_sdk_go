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

// QueryRebateBalanceV2FilteringInvoiceStatusList
type QueryRebateBalanceV2FilteringInvoiceStatusList string

// List of query_rebate_balance_v2_filtering_invoice_status_list
const (
	NO_QueryRebateBalanceV2FilteringInvoiceStatusList       QueryRebateBalanceV2FilteringInvoiceStatusList = "NO"
	PART_QueryRebateBalanceV2FilteringInvoiceStatusList     QueryRebateBalanceV2FilteringInvoiceStatusList = "PART"
	COMPLETE_QueryRebateBalanceV2FilteringInvoiceStatusList QueryRebateBalanceV2FilteringInvoiceStatusList = "COMPLETE"
)

// All allowed values of QueryRebateBalanceV2FilteringInvoiceStatusList enum
var AllowedQueryRebateBalanceV2FilteringInvoiceStatusListEnumValues = []QueryRebateBalanceV2FilteringInvoiceStatusList{
	"NO",
	"PART",
	"COMPLETE",
}

// NewQueryRebateBalanceV2FilteringInvoiceStatusListFromValue returns a pointer to a valid QueryRebateBalanceV2FilteringInvoiceStatusList
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryRebateBalanceV2FilteringInvoiceStatusListFromValue(v string) (*QueryRebateBalanceV2FilteringInvoiceStatusList, error) {
	ev := QueryRebateBalanceV2FilteringInvoiceStatusList(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryRebateBalanceV2FilteringInvoiceStatusList: valid values are %v", v, AllowedQueryRebateBalanceV2FilteringInvoiceStatusListEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryRebateBalanceV2FilteringInvoiceStatusList) IsValid() bool {
	for _, existing := range AllowedQueryRebateBalanceV2FilteringInvoiceStatusListEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to query_rebate_balance_v2_filtering_invoice_status_list value
func (v QueryRebateBalanceV2FilteringInvoiceStatusList) Ptr() *QueryRebateBalanceV2FilteringInvoiceStatusList {
	return &v
}
