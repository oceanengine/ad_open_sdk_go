/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BudgetGroupCreateV30Request struct for BudgetGroupCreateV30Request
type BudgetGroupCreateV30Request struct {
	// 广告主 ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 日预算，单位：元，范围：[300，9,999,999]，预算组中的每个项目在实际消耗时，以【项目、预算组】中较低预算为准
	BudgetGroupBudget float64 `json:"budget_group_budget"`
	// 预算组名称，50个字以内，预算组之间名称不可重复，不可为空
	BudgetGroupName string                           `json:"budget_group_name"`
	CampaignType    BudgetGroupCreateV30CampaignType `json:"campaign_type"`
	DeliveryMode    BudgetGroupCreateV30DeliveryMode `json:"delivery_mode"`
}
