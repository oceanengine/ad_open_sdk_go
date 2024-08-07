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

// SharedWalletTransactionDetailGetV30AccountType
type SharedWalletTransactionDetailGetV30AccountType string

// List of shared_wallet_transaction_detail_get_v3.0_account_type
const (
	AD_SharedWalletTransactionDetailGetV30AccountType        SharedWalletTransactionDetailGetV30AccountType = "AD"
	AGENT_SharedWalletTransactionDetailGetV30AccountType     SharedWalletTransactionDetailGetV30AccountType = "AGENT"
	LOCAL_SharedWalletTransactionDetailGetV30AccountType     SharedWalletTransactionDetailGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletTransactionDetailGetV30AccountType SharedWalletTransactionDetailGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletTransactionDetailGetV30AccountType enum
var AllowedSharedWalletTransactionDetailGetV30AccountTypeEnumValues = []SharedWalletTransactionDetailGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletTransactionDetailGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletTransactionDetailGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletTransactionDetailGetV30AccountTypeFromValue(v string) (*SharedWalletTransactionDetailGetV30AccountType, error) {
	ev := SharedWalletTransactionDetailGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletTransactionDetailGetV30AccountType: valid values are %v", v, AllowedSharedWalletTransactionDetailGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletTransactionDetailGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletTransactionDetailGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_transaction_detail_get_v3.0_account_type value
func (v SharedWalletTransactionDetailGetV30AccountType) Ptr() *SharedWalletTransactionDetailGetV30AccountType {
	return &v
}
