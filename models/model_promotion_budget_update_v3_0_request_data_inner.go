/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionBudgetUpdateV30RequestDataInner struct for PromotionBudgetUpdateV30RequestDataInner
type PromotionBudgetUpdateV30RequestDataInner struct {
	// 预算，单位元
	Budget     float64                                 `json:"budget"`
	BudgetMode *PromotionBudgetUpdateV30DataBudgetMode `json:"budget_mode,omitempty"`
	// 广告ID
	PromotionId int64 `json:"promotion_id"`
}
