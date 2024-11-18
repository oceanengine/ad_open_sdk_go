/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2FilteringStatus
type CampaignGetV2FilteringStatus string

// List of campaign_get_v2_filtering_status
const (
	CAMPAIGN_STATUS_NOT_DELETE_CampaignGetV2FilteringStatus CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_NOT_DELETE"
	CAMPAIGN_STATUS_DELETE_CampaignGetV2FilteringStatus     CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DELETE"
	CAMPAIGN_STATUS_DISABLE_CampaignGetV2FilteringStatus    CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_DISABLE"
	CAMPAIGN_STATUS_ALL_CampaignGetV2FilteringStatus        CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ALL"
	CAMPAIGN_STATUS_ENABLE_CampaignGetV2FilteringStatus     CampaignGetV2FilteringStatus = "CAMPAIGN_STATUS_ENABLE"
)

// Ptr returns reference to campaign_get_v2_filtering_status value
func (v CampaignGetV2FilteringStatus) Ptr() *CampaignGetV2FilteringStatus {
	return &v
}
