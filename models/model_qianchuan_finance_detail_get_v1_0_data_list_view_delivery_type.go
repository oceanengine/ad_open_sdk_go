/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanFinanceDetailGetV10DataListViewDeliveryType
type QianchuanFinanceDetailGetV10DataListViewDeliveryType string

// List of qianchuan_finance_detail_get_v1.0_data_list_view_delivery_type
const (
	ALL_QianchuanFinanceDetailGetV10DataListViewDeliveryType     QianchuanFinanceDetailGetV10DataListViewDeliveryType = "ALL"
	DEFAULT_QianchuanFinanceDetailGetV10DataListViewDeliveryType QianchuanFinanceDetailGetV10DataListViewDeliveryType = "DEFAULT"
	BRAND_QianchuanFinanceDetailGetV10DataListViewDeliveryType   QianchuanFinanceDetailGetV10DataListViewDeliveryType = "BRAND"
)

// All allowed values of QianchuanFinanceDetailGetV10DataListViewDeliveryType enum
var AllowedQianchuanFinanceDetailGetV10DataListViewDeliveryTypeEnumValues = []QianchuanFinanceDetailGetV10DataListViewDeliveryType{
	"ALL",
	"DEFAULT",
	"BRAND",
}

// NewQianchuanFinanceDetailGetV10DataListViewDeliveryTypeFromValue returns a pointer to a valid QianchuanFinanceDetailGetV10DataListViewDeliveryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanFinanceDetailGetV10DataListViewDeliveryTypeFromValue(v string) (*QianchuanFinanceDetailGetV10DataListViewDeliveryType, error) {
	ev := QianchuanFinanceDetailGetV10DataListViewDeliveryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanFinanceDetailGetV10DataListViewDeliveryType: valid values are %v", v, AllowedQianchuanFinanceDetailGetV10DataListViewDeliveryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanFinanceDetailGetV10DataListViewDeliveryType) IsValid() bool {
	for _, existing := range AllowedQianchuanFinanceDetailGetV10DataListViewDeliveryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_finance_detail_get_v1.0_data_list_view_delivery_type value
func (v QianchuanFinanceDetailGetV10DataListViewDeliveryType) Ptr() *QianchuanFinanceDetailGetV10DataListViewDeliveryType {
	return &v
}
