/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBudgetUpdateV30DataBudgetMode
type PromotionBudgetUpdateV30DataBudgetMode string

// List of promotion_budget_update_v3.0_data_budget_mode
const (
	BUDGET_MODE_DAY_PromotionBudgetUpdateV30DataBudgetMode   PromotionBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_PromotionBudgetUpdateV30DataBudgetMode PromotionBudgetUpdateV30DataBudgetMode = "BUDGET_MODE_TOTAL"
)

// Ptr returns reference to promotion_budget_update_v3.0_data_budget_mode value
func (v PromotionBudgetUpdateV30DataBudgetMode) Ptr() *PromotionBudgetUpdateV30DataBudgetMode {
	return &v
}
