/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingBudgetMode
type QianchuanAdGetV10DataListDeliverySettingBudgetMode string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanAdGetV10DataListDeliverySettingBudgetMode             QianchuanAdGetV10DataListDeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanAdGetV10DataListDeliverySettingBudgetMode        QianchuanAdGetV10DataListDeliverySettingBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_SEVEN_DAY_CYCLE_QianchuanAdGetV10DataListDeliverySettingBudgetMode QianchuanAdGetV10DataListDeliverySettingBudgetMode = "BUDGET_MODE_SEVEN_DAY_CYCLE"
	BUDGET_MODE_TOTAL_QianchuanAdGetV10DataListDeliverySettingBudgetMode           QianchuanAdGetV10DataListDeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingBudgetMode enum
var AllowedQianchuanAdGetV10DataListDeliverySettingBudgetModeEnumValues = []QianchuanAdGetV10DataListDeliverySettingBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_SEVEN_DAY_CYCLE",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanAdGetV10DataListDeliverySettingBudgetModeFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingBudgetModeFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingBudgetMode, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingBudgetMode: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_budget_mode value
func (v QianchuanAdGetV10DataListDeliverySettingBudgetMode) Ptr() *QianchuanAdGetV10DataListDeliverySettingBudgetMode {
	return &v
}
