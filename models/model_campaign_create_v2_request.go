/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignCreateV2Request struct for CampaignCreateV2Request
type CampaignCreateV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Budget                     *float64                                    `json:"budget,omitempty"`
	BudgetMode                 CampaignCreateV2BudgetMode                  `json:"budget_mode"`
	CampaignBudgetOptimization *CampaignCreateV2CampaignBudgetOptimization `json:"campaign_budget_optimization,omitempty"`
	//
	CampaignName       string                              `json:"campaign_name"`
	CampaignType       *CampaignCreateV2CampaignType       `json:"campaign_type,omitempty"`
	DedicateType       *CampaignCreateV2DedicateType       `json:"dedicate_type,omitempty"`
	DeliveryMode       *CampaignCreateV2DeliveryMode       `json:"delivery_mode,omitempty"`
	DeliveryRelatedNum *CampaignCreateV2DeliveryRelatedNum `json:"delivery_related_num,omitempty"`
	LandingType        CampaignCreateV2LandingType         `json:"landing_type"`
	MarketingPurpose   *CampaignCreateV2MarketingPurpose   `json:"marketing_purpose,omitempty"`
	MarketingScene     *CampaignCreateV2MarketingScene     `json:"marketing_scene,omitempty"`
	Operation          *CampaignCreateV2Operation          `json:"operation,omitempty"`
	//
	OptPermission *int64                        `json:"opt_permission,omitempty"`
	SmartBidType  *CampaignCreateV2SmartBidType `json:"smart_bid_type,omitempty"`
}
