/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceInterestActionListV2Filtering
type ReportAudienceInterestActionListV2Filtering struct {
	ActionDays *ReportAudienceInterestActionListV2FilteringActionDays `json:"action_days,omitempty"`
	//
	ActionScene []*ReportAudienceInterestActionListV2FilteringActionScene `json:"action_scene,omitempty"`
	//
	AdIds         []int64                                                   `json:"ad_ids,omitempty"`
	AudienceLevel *ReportAudienceInterestActionListV2FilteringAudienceLevel `json:"audience_level,omitempty"`
	//
	CampaignIds        []int64                                                        `json:"campaign_ids,omitempty"`
	InterestActionType *ReportAudienceInterestActionListV2FilteringInterestActionType `json:"interest_action_type,omitempty"`
}
