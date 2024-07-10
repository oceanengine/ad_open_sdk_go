/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AccountFundGetV30AccountType
type AccountFundGetV30AccountType string

// List of account_fund_get_v3.0_account_type
const (
	AD_AccountFundGetV30AccountType    AccountFundGetV30AccountType = "AD"
	STAR_AccountFundGetV30AccountType  AccountFundGetV30AccountType = "STAR"
	LOCAL_AccountFundGetV30AccountType AccountFundGetV30AccountType = "LOCAL"
)

// All allowed values of AccountFundGetV30AccountType enum
var AllowedAccountFundGetV30AccountTypeEnumValues = []AccountFundGetV30AccountType{
	"AD",
	"STAR",
	"LOCAL",
}

// NewAccountFundGetV30AccountTypeFromValue returns a pointer to a valid AccountFundGetV30AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountFundGetV30AccountTypeFromValue(v string) (*AccountFundGetV30AccountType, error) {
	ev := AccountFundGetV30AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountFundGetV30AccountType: valid values are %v", v, AllowedAccountFundGetV30AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountFundGetV30AccountType) IsValid() bool {
	for _, existing := range AllowedAccountFundGetV30AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to account_fund_get_v3.0_account_type value
func (v AccountFundGetV30AccountType) Ptr() *AccountFundGetV30AccountType {
	return &v
}
