/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10DeliverySettingBudgetMode
type QianchuanAdCreateV10DeliverySettingBudgetMode string

// List of qianchuan_ad_create_v1.0_delivery_setting_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanAdCreateV10DeliverySettingBudgetMode             QianchuanAdCreateV10DeliverySettingBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_SEVEN_DAY_CYCLE_QianchuanAdCreateV10DeliverySettingBudgetMode QianchuanAdCreateV10DeliverySettingBudgetMode = "BUDGET_MODE_SEVEN_DAY_CYCLE"
	BUDGET_MODE_TOTAL_QianchuanAdCreateV10DeliverySettingBudgetMode           QianchuanAdCreateV10DeliverySettingBudgetMode = "BUDGET_MODE_TOTAL"
)

// Ptr returns reference to qianchuan_ad_create_v1.0_delivery_setting_budget_mode value
func (v QianchuanAdCreateV10DeliverySettingBudgetMode) Ptr() *QianchuanAdCreateV10DeliverySettingBudgetMode {
	return &v
}
