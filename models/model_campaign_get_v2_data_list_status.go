/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2DataListStatus
type CampaignGetV2DataListStatus string

// List of campaign_get_v2_data_list_status
const (
	CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED_CampaignGetV2DataListStatus CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_ADVERTISER_BUDGET_EXCEED"
	CAMPAIGN_STATUS_DELETE_CampaignGetV2DataListStatus                   CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_DELETE"
	CAMPAIGN_STATUS_DISABLE_CampaignGetV2DataListStatus                  CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_DISABLE"
	CAMPAIGN_STATUS_ENABLE_CampaignGetV2DataListStatus                   CampaignGetV2DataListStatus = "CAMPAIGN_STATUS_ENABLE"
)

// Ptr returns reference to campaign_get_v2_data_list_status value
func (v CampaignGetV2DataListStatus) Ptr() *CampaignGetV2DataListStatus {
	return &v
}
