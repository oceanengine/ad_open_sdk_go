/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserFundGrantTransactionGetV2DataListTransactionType
type AdvertiserFundGrantTransactionGetV2DataListTransactionType string

// List of advertiser_fund_grant_transaction_get_v2_data_list_transaction_type
const (
	RECHARGE_AdvertiserFundGrantTransactionGetV2DataListTransactionType AdvertiserFundGrantTransactionGetV2DataListTransactionType = "RECHARGE"
	TRANSFER_AdvertiserFundGrantTransactionGetV2DataListTransactionType AdvertiserFundGrantTransactionGetV2DataListTransactionType = "TRANSFER"
)

// All allowed values of AdvertiserFundGrantTransactionGetV2DataListTransactionType enum
var AllowedAdvertiserFundGrantTransactionGetV2DataListTransactionTypeEnumValues = []AdvertiserFundGrantTransactionGetV2DataListTransactionType{
	"RECHARGE",
	"TRANSFER",
}

// NewAdvertiserFundGrantTransactionGetV2DataListTransactionTypeFromValue returns a pointer to a valid AdvertiserFundGrantTransactionGetV2DataListTransactionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundGrantTransactionGetV2DataListTransactionTypeFromValue(v string) (*AdvertiserFundGrantTransactionGetV2DataListTransactionType, error) {
	ev := AdvertiserFundGrantTransactionGetV2DataListTransactionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundGrantTransactionGetV2DataListTransactionType: valid values are %v", v, AllowedAdvertiserFundGrantTransactionGetV2DataListTransactionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundGrantTransactionGetV2DataListTransactionType) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundGrantTransactionGetV2DataListTransactionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_grant_transaction_get_v2_data_list_transaction_type value
func (v AdvertiserFundGrantTransactionGetV2DataListTransactionType) Ptr() *AdvertiserFundGrantTransactionGetV2DataListTransactionType {
	return &v
}
