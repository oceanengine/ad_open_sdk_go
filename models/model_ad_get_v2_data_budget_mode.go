/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataBudgetMode
type AdGetV2DataBudgetMode string

// List of ad_get_v2_data_budget_mode
const (
	BUDGET_MODE_TOTAL_AdGetV2DataBudgetMode    AdGetV2DataBudgetMode = "BUDGET_MODE_TOTAL"
	BUDGET_MODE_DAY_AdGetV2DataBudgetMode      AdGetV2DataBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_AdGetV2DataBudgetMode AdGetV2DataBudgetMode = "BUDGET_MODE_INFINITE"
)

// Ptr returns reference to ad_get_v2_data_budget_mode value
func (v AdGetV2DataBudgetMode) Ptr() *AdGetV2DataBudgetMode {
	return &v
}
