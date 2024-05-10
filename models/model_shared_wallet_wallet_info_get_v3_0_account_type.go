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

// SharedWalletWalletInfoGetV30AccountType
type SharedWalletWalletInfoGetV30AccountType string

// List of shared_wallet_wallet_info_get_v3.0_account_type
const (
	AD_SharedWalletWalletInfoGetV30AccountType        SharedWalletWalletInfoGetV30AccountType = "AD"
	AGENT_SharedWalletWalletInfoGetV30AccountType     SharedWalletWalletInfoGetV30AccountType = "AGENT"
	LOCAL_SharedWalletWalletInfoGetV30AccountType     SharedWalletWalletInfoGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletWalletInfoGetV30AccountType SharedWalletWalletInfoGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletWalletInfoGetV30AccountType enum
var AllowedSharedWalletWalletInfoGetV30AccountTypeEnumValues = []SharedWalletWalletInfoGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletWalletInfoGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletWalletInfoGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletWalletInfoGetV30AccountTypeFromValue(v string) (*SharedWalletWalletInfoGetV30AccountType, error) {
	ev := SharedWalletWalletInfoGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletWalletInfoGetV30AccountType: valid values are %v", v, AllowedSharedWalletWalletInfoGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletWalletInfoGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletWalletInfoGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_wallet_info_get_v3.0_account_type value
func (v SharedWalletWalletInfoGetV30AccountType) Ptr() *SharedWalletWalletInfoGetV30AccountType {
	return &v
}
