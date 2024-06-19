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

// SharedWalletMainWalletGetV30AccountType
type SharedWalletMainWalletGetV30AccountType string

// List of shared_wallet_main_wallet_get_v3.0_account_type
const (
	AD_SharedWalletMainWalletGetV30AccountType        SharedWalletMainWalletGetV30AccountType = "AD"
	AGENT_SharedWalletMainWalletGetV30AccountType     SharedWalletMainWalletGetV30AccountType = "AGENT"
	LOCAL_SharedWalletMainWalletGetV30AccountType     SharedWalletMainWalletGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletMainWalletGetV30AccountType SharedWalletMainWalletGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletMainWalletGetV30AccountType enum
var AllowedSharedWalletMainWalletGetV30AccountTypeEnumValues = []SharedWalletMainWalletGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletMainWalletGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletMainWalletGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletMainWalletGetV30AccountTypeFromValue(v string) (*SharedWalletMainWalletGetV30AccountType, error) {
	ev := SharedWalletMainWalletGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletMainWalletGetV30AccountType: valid values are %v", v, AllowedSharedWalletMainWalletGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletMainWalletGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletMainWalletGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_main_wallet_get_v3.0_account_type value
func (v SharedWalletMainWalletGetV30AccountType) Ptr() *SharedWalletMainWalletGetV30AccountType {
	return &v
}
