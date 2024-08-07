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

// SharedWalletDailyStatGetV30AccountType
type SharedWalletDailyStatGetV30AccountType string

// List of shared_wallet_daily_stat_get_v3.0_account_type
const (
	AD_SharedWalletDailyStatGetV30AccountType        SharedWalletDailyStatGetV30AccountType = "AD"
	AGENT_SharedWalletDailyStatGetV30AccountType     SharedWalletDailyStatGetV30AccountType = "AGENT"
	LOCAL_SharedWalletDailyStatGetV30AccountType     SharedWalletDailyStatGetV30AccountType = "LOCAL"
	QIANCHUAN_SharedWalletDailyStatGetV30AccountType SharedWalletDailyStatGetV30AccountType = "QIANCHUAN"
)

// All allowed values of SharedWalletDailyStatGetV30AccountType enum
var AllowedSharedWalletDailyStatGetV30AccountTypeEnumValues = []SharedWalletDailyStatGetV30AccountType{
	"AD",
	"AGENT",
	"LOCAL",
	"QIANCHUAN",
}

// NewSharedWalletDailyStatGetV30AccountTypeFromValue returns a pointer to a valid SharedWalletDailyStatGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSharedWalletDailyStatGetV30AccountTypeFromValue(v string) (*SharedWalletDailyStatGetV30AccountType, error) {
	ev := SharedWalletDailyStatGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SharedWalletDailyStatGetV30AccountType: valid values are %v", v, AllowedSharedWalletDailyStatGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SharedWalletDailyStatGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedSharedWalletDailyStatGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to shared_wallet_daily_stat_get_v3.0_account_type value
func (v SharedWalletDailyStatGetV30AccountType) Ptr() *SharedWalletDailyStatGetV30AccountType {
	return &v
}
