/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAudiencePushV10AccountType
type QianchuanAudiencePushV10AccountType string

// List of qianchuan_audience_push_v1.0_account_type
const (
	LOCAL_QianchuanAudiencePushV10AccountType    QianchuanAudiencePushV10AccountType = "LOCAL"
	NO_LOCAL_QianchuanAudiencePushV10AccountType QianchuanAudiencePushV10AccountType = "NO_LOCAL"
)

// All allowed values of QianchuanAudiencePushV10AccountType enum
var AllowedQianchuanAudiencePushV10AccountTypeEnumValues = []QianchuanAudiencePushV10AccountType{
	"LOCAL",
	"NO_LOCAL",
}

// NewQianchuanAudiencePushV10AccountTypeFromValue returns a pointer to a valid QianchuanAudiencePushV10AccountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAudiencePushV10AccountTypeFromValue(v string) (*QianchuanAudiencePushV10AccountType, error) {
	ev := QianchuanAudiencePushV10AccountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAudiencePushV10AccountType: valid values are %v", v, AllowedQianchuanAudiencePushV10AccountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAudiencePushV10AccountType) IsValid() bool {
	for _, existing := range AllowedQianchuanAudiencePushV10AccountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_audience_push_v1.0_account_type value
func (v QianchuanAudiencePushV10AccountType) Ptr() *QianchuanAudiencePushV10AccountType {
	return &v
}
