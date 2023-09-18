/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanFinanceDetailGetV10Type
type QianchuanFinanceDetailGetV10Type string

// List of qianchuan_finance_detail_get_v1.0_type
const (
	ALL_QianchuanFinanceDetailGetV10Type     QianchuanFinanceDetailGetV10Type = "ALL"
	DEFAULT_QianchuanFinanceDetailGetV10Type QianchuanFinanceDetailGetV10Type = "DEFAULT"
	BRAND_QianchuanFinanceDetailGetV10Type   QianchuanFinanceDetailGetV10Type = "BRAND"
)

// All allowed values of QianchuanFinanceDetailGetV10Type enum
var AllowedQianchuanFinanceDetailGetV10TypeEnumValues = []QianchuanFinanceDetailGetV10Type{
	"ALL",
	"DEFAULT",
	"BRAND",
}

// NewQianchuanFinanceDetailGetV10TypeFromValue returns a pointer to a valid QianchuanFinanceDetailGetV10Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanFinanceDetailGetV10TypeFromValue(v string) (*QianchuanFinanceDetailGetV10Type, error) {
	ev := QianchuanFinanceDetailGetV10Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanFinanceDetailGetV10Type: valid values are %v", v, AllowedQianchuanFinanceDetailGetV10TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanFinanceDetailGetV10Type) IsValid() bool {
	for _, existing := range AllowedQianchuanFinanceDetailGetV10TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_finance_detail_get_v1.0_type value
func (v QianchuanFinanceDetailGetV10Type) Ptr() *QianchuanFinanceDetailGetV10Type {
	return &v
}