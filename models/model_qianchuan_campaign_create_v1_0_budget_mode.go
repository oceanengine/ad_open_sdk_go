/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignCreateV10BudgetMode
type QianchuanCampaignCreateV10BudgetMode string

// List of qianchuan_campaign_create_v1.0_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanCampaignCreateV10BudgetMode      QianchuanCampaignCreateV10BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanCampaignCreateV10BudgetMode QianchuanCampaignCreateV10BudgetMode = "BUDGET_MODE_INFINITE"
)

// Ptr returns reference to qianchuan_campaign_create_v1.0_budget_mode value
func (v QianchuanCampaignCreateV10BudgetMode) Ptr() *QianchuanCampaignCreateV10BudgetMode {
	return &v
}
