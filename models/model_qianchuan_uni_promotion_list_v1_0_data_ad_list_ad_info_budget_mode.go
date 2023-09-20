/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode
type QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode string

// List of qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode      QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode    QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode enum
var AllowedQianchuanUniPromotionListV10DataAdListAdInfoBudgetModeEnumValues = []QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanUniPromotionListV10DataAdListAdInfoBudgetModeFromValue returns a pointer to a valid QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanUniPromotionListV10DataAdListAdInfoBudgetModeFromValue(v string) (*QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode, error) {
	ev := QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode: valid values are %v", v, AllowedQianchuanUniPromotionListV10DataAdListAdInfoBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanUniPromotionListV10DataAdListAdInfoBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_uni_promotion_list_v1.0_data_ad_list_ad_info_budget_mode value
func (v QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode) Ptr() *QianchuanUniPromotionListV10DataAdListAdInfoBudgetMode {
	return &v
}
