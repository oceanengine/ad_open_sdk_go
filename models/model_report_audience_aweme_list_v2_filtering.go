/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceAwemeListV2Filtering
type ReportAudienceAwemeListV2Filtering struct {
	//
	AdIds         []int64                                          `json:"ad_ids,omitempty"`
	AudienceLevel *ReportAudienceAwemeListV2FilteringAudienceLevel `json:"audience_level,omitempty"`
	//
	Behaviors []*ReportAudienceAwemeListV2FilteringBehaviors `json:"behaviors,omitempty"`
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
}
