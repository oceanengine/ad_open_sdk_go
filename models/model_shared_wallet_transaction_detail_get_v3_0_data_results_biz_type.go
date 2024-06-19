/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// SharedWalletTransactionDetailGetV30DataResultsBizType
type SharedWalletTransactionDetailGetV30DataResultsBizType string

// List of shared_wallet_transaction_detail_get_v3.0_data_results_biz_type
const (
	ADJUST_DECREASE_SharedWalletTransactionDetailGetV30DataResultsBizType        SharedWalletTransactionDetailGetV30DataResultsBizType = "ADJUST_DECREASE"
	ADJUST_FREEZE_SharedWalletTransactionDetailGetV30DataResultsBizType          SharedWalletTransactionDetailGetV30DataResultsBizType = "ADJUST_FREEZE"
	ADJUST_INCREASE_SharedWalletTransactionDetailGetV30DataResultsBizType        SharedWalletTransactionDetailGetV30DataResultsBizType = "ADJUST_INCREASE"
	CREDIT_RECHARGE_SharedWalletTransactionDetailGetV30DataResultsBizType        SharedWalletTransactionDetailGetV30DataResultsBizType = "CREDIT_RECHARGE"
	CREDIT_REFUND_SharedWalletTransactionDetailGetV30DataResultsBizType          SharedWalletTransactionDetailGetV30DataResultsBizType = "CREDIT_REFUND"
	INIT_SharedWalletTransactionDetailGetV30DataResultsBizType                   SharedWalletTransactionDetailGetV30DataResultsBizType = "INIT"
	ORDER_PAY_SharedWalletTransactionDetailGetV30DataResultsBizType              SharedWalletTransactionDetailGetV30DataResultsBizType = "ORDER_PAY"
	ORDER_WITHDRAW_SharedWalletTransactionDetailGetV30DataResultsBizType         SharedWalletTransactionDetailGetV30DataResultsBizType = "ORDER_WITHDRAW"
	REFUND_FREEZE_SharedWalletTransactionDetailGetV30DataResultsBizType          SharedWalletTransactionDetailGetV30DataResultsBizType = "REFUND_FREEZE"
	SHARED_CANCEL_RECHARGE_SharedWalletTransactionDetailGetV30DataResultsBizType SharedWalletTransactionDetailGetV30DataResultsBizType = "SHARED_CANCEL_RECHARGE"
	SHARED_FROZEN_REFUND_SharedWalletTransactionDetailGetV30DataResultsBizType   SharedWalletTransactionDetailGetV30DataResultsBizType = "SHARED_FROZEN_REFUND"
	SHARED_RECHARGE_SharedWalletTransactionDetailGetV30DataResultsBizType        SharedWalletTransactionDetailGetV30DataResultsBizType = "SHARED_RECHARGE"
	TRANSFER_SharedWalletTransactionDetailGetV30DataResultsBizType               SharedWalletTransactionDetailGetV30DataResultsBizType = "TRANSFER"
	TRANSFER_IN_SharedWalletTransactionDetailGetV30DataResultsBizType            SharedWalletTransactionDetailGetV30DataResultsBizType = "TRANSFER_IN"
	TRANSFER_OUT_SharedWalletTransactionDetailGetV30DataResultsBizType           SharedWalletTransactionDetailGetV30DataResultsBizType = "TRANSFER_OUT"
)

// All allowed values of SharedWalletTransactionDetailGetV30DataResultsBizType enum
var AllowedSharedWalletTransactionDetailGetV30DataResultsBizTypeEnumValues = []SharedWalletTransactionDetailGetV30DataResultsBizType{
	"ADJUST_DECREASE",
	"ADJUST_FREEZE",
	"ADJUST_INCREASE",
	"CREDIT_RECHARGE",
	"CREDIT_REFUND",
	"INIT",
	"ORDER_PAY",
	"ORDER_WITHDRAW",
	"REFUND_FREEZE",
	"SHARED_CANCEL_RECHARGE",
	"SHARED_FROZEN_REFUND",
	"SHARED_RECHARGE",
	"TRANSFER",
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewSharedWalletTransactionDetailGetV30DataResultsBizTypeFromValue returns a pointer to a valid SharedWalletTransactionDetailGetV30DataResultsBizType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletTransactionDetailGetV30DataResultsBizTypeFromValue(v string) (*SharedWalletTransactionDetailGetV30DataResultsBizType, error) {
	ev := SharedWalletTransactionDetailGetV30DataResultsBizType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletTransactionDetailGetV30DataResultsBizType: valid values are %v", v, AllowedSharedWalletTransactionDetailGetV30DataResultsBizTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletTransactionDetailGetV30DataResultsBizType) IsValid() bool {
	for _, existing := range AllowedSharedWalletTransactionDetailGetV30DataResultsBizTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_transaction_detail_get_v3.0_data_results_biz_type value
func (v SharedWalletTransactionDetailGetV30DataResultsBizType) Ptr() *SharedWalletTransactionDetailGetV30DataResultsBizType {
	return &v
}
