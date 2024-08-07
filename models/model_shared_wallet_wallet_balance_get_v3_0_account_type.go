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

// SharedWalletWalletBalanceGetV30AccountType
type SharedWalletWalletBalanceGetV30AccountType string

// List of shared_wallet_wallet_balance_get_v3.0_account_type
const (
	AD_SharedWalletWalletBalanceGetV30AccountType        SharedWalletWalletBalanceGetV30AccountType = "AD"
	AGENT_SharedWalletWalletBalanceGetV30AccountType     SharedWalletWalletBalanceGetV30AccountType = "AGENT"
	LOCAL_SharedWalletWalletBalanceGetV30AccountType     SharedWalletWalletBalanceGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletWalletBalanceGetV30AccountType SharedWalletWalletBalanceGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletWalletBalanceGetV30AccountType enum
var AllowedSharedWalletWalletBalanceGetV30AccountTypeEnumValues = []SharedWalletWalletBalanceGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletWalletBalanceGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletWalletBalanceGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletWalletBalanceGetV30AccountTypeFromValue(v string) (*SharedWalletWalletBalanceGetV30AccountType, error) {
	ev := SharedWalletWalletBalanceGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletWalletBalanceGetV30AccountType: valid values are %v", v, AllowedSharedWalletWalletBalanceGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletWalletBalanceGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletWalletBalanceGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_wallet_balance_get_v3.0_account_type value
func (v SharedWalletWalletBalanceGetV30AccountType) Ptr() *SharedWalletWalletBalanceGetV30AccountType {
	return &v
}
