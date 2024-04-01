/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdvertiserFundGrantTransactionGetV2DataListGrantType
type AdvertiserFundGrantTransactionGetV2DataListGrantType string

// List of advertiser_fund_grant_transaction_get_v2_data_list_grant_type
const (
	UNKNOWN_AdvertiserFundGrantTransactionGetV2DataListGrantType      AdvertiserFundGrantTransactionGetV2DataListGrantType = "unknown"
	ACTIVITY_AdvertiserFundGrantTransactionGetV2DataListGrantType     AdvertiserFundGrantTransactionGetV2DataListGrantType = "activity"
	REFUND_AdvertiserFundGrantTransactionGetV2DataListGrantType       AdvertiserFundGrantTransactionGetV2DataListGrantType = "refund"
	INTERNAL_AD_AdvertiserFundGrantTransactionGetV2DataListGrantType  AdvertiserFundGrantTransactionGetV2DataListGrantType = "internal_ad"
	RETURN_GOODS_AdvertiserFundGrantTransactionGetV2DataListGrantType AdvertiserFundGrantTransactionGetV2DataListGrantType = "return_goods"
	EXCHANGE_AdvertiserFundGrantTransactionGetV2DataListGrantType     AdvertiserFundGrantTransactionGetV2DataListGrantType = "exchange"
	SUPPLEMENT_AdvertiserFundGrantTransactionGetV2DataListGrantType   AdvertiserFundGrantTransactionGetV2DataListGrantType = "supplement"
	SUBSIDIARY_AdvertiserFundGrantTransactionGetV2DataListGrantType   AdvertiserFundGrantTransactionGetV2DataListGrantType = "subsidiary"
)

// All allowed values of AdvertiserFundGrantTransactionGetV2DataListGrantType enum
var AllowedAdvertiserFundGrantTransactionGetV2DataListGrantTypeEnumValues = []AdvertiserFundGrantTransactionGetV2DataListGrantType{
	"unknown",
	"activity",
	"refund",
	"internal_ad",
	"return_goods",
	"exchange",
	"supplement",
	"subsidiary",
}

// NewAdvertiserFundGrantTransactionGetV2DataListGrantTypeFromValue returns a pointer to a valid AdvertiserFundGrantTransactionGetV2DataListGrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserFundGrantTransactionGetV2DataListGrantTypeFromValue(v string) (*AdvertiserFundGrantTransactionGetV2DataListGrantType, error) {
	ev := AdvertiserFundGrantTransactionGetV2DataListGrantType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserFundGrantTransactionGetV2DataListGrantType: valid values are %v", v, AllowedAdvertiserFundGrantTransactionGetV2DataListGrantTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserFundGrantTransactionGetV2DataListGrantType) IsValid() bool {
	for _, existing := range AllowedAdvertiserFundGrantTransactionGetV2DataListGrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_fund_grant_transaction_get_v2_data_list_grant_type value
func (v AdvertiserFundGrantTransactionGetV2DataListGrantType) Ptr() *AdvertiserFundGrantTransactionGetV2DataListGrantType {
	return &v
}
