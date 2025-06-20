/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignUpdateV2BudgetMode
type CampaignUpdateV2BudgetMode string

// List of campaign_update_v2_budget_mode
const (
	BUDGET_MODE_DAY_CampaignUpdateV2BudgetMode      CampaignUpdateV2BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_TOTAL_CampaignUpdateV2BudgetMode    CampaignUpdateV2BudgetMode = "BUDGET_MODE_TOTAL"
	BUDGET_MODE_INFINITE_CampaignUpdateV2BudgetMode CampaignUpdateV2BudgetMode = "BUDGET_MODE_INFINITE"
)

// Ptr returns reference to campaign_update_v2_budget_mode value
func (v CampaignUpdateV2BudgetMode) Ptr() *CampaignUpdateV2BudgetMode {
	return &v
}
