/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignUpdateV10BudgetMode
type QianchuanCampaignUpdateV10BudgetMode string

// List of qianchuan_campaign_update_v1.0_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanCampaignUpdateV10BudgetMode      QianchuanCampaignUpdateV10BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanCampaignUpdateV10BudgetMode QianchuanCampaignUpdateV10BudgetMode = "BUDGET_MODE_INFINITE"
)

// Ptr returns reference to qianchuan_campaign_update_v1.0_budget_mode value
func (v QianchuanCampaignUpdateV10BudgetMode) Ptr() *QianchuanCampaignUpdateV10BudgetMode {
	return &v
}
