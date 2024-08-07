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

// QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory
type QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory string

// List of qianchuan_finance_wallet_get_v1.0_data_share_expiring_detail_list_category
const (
	COMMON_QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory  QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory = "COMMON"
	DEFAULT_QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory = "DEFAULT"
	SEARCH_QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory  QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory = "SEARCH"
	UNION_QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory   QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory = "UNION"
)

// All allowed values of QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory enum
var AllowedQianchuanFinanceWalletGetV10DataShareExpiringDetailListCategoryEnumValues = []QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory{
	"COMMON",
	"DEFAULT",
	"SEARCH",
	"UNION",
}

// NewQianchuanFinanceWalletGetV10DataShareExpiringDetailListCategoryFromValue returns a pointer to a valid QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanFinanceWalletGetV10DataShareExpiringDetailListCategoryFromValue(v string) (*QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory, error) {
	ev := QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory: valid values are %v", v, AllowedQianchuanFinanceWalletGetV10DataShareExpiringDetailListCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory) IsValid() bool {
	for _, existing := range AllowedQianchuanFinanceWalletGetV10DataShareExpiringDetailListCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_finance_wallet_get_v1.0_data_share_expiring_detail_list_category value
func (v QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory) Ptr() *QianchuanFinanceWalletGetV10DataShareExpiringDetailListCategory {
	return &v
}
