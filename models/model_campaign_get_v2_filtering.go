/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2Filtering
type CampaignGetV2Filtering struct {
	CampaignBudgetOptimization *CampaignGetV2FilteringCampaignBudgetOptimization `json:"campaign_budget_optimization,omitempty"`
	//
	CampaignCreateTime **string `json:"campaign_create_time,omitempty"`
	//
	CampaignName *string                             `json:"campaign_name,omitempty"`
	DedicateType *CampaignGetV2FilteringDedicateType `json:"dedicate_type,omitempty"`
	//
	Ids          []int64                             `json:"ids,omitempty"`
	LandingType  *CampaignGetV2FilteringLandingType  `json:"landing_type,omitempty"`
	SmartBidType *CampaignGetV2FilteringSmartBidType `json:"smart_bid_type,omitempty"`
	Status       *CampaignGetV2FilteringStatus       `json:"status,omitempty"`
}
