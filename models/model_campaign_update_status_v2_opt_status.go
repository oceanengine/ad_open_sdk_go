/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignUpdateStatusV2OptStatus
type CampaignUpdateStatusV2OptStatus string

// List of campaign_update_status_v2_opt_status
const (
	DELETE_CampaignUpdateStatusV2OptStatus  CampaignUpdateStatusV2OptStatus = "delete"
	ENABLE_CampaignUpdateStatusV2OptStatus  CampaignUpdateStatusV2OptStatus = "enable"
	DISABLE_CampaignUpdateStatusV2OptStatus CampaignUpdateStatusV2OptStatus = "disable"
)

// Ptr returns reference to campaign_update_status_v2_opt_status value
func (v CampaignUpdateStatusV2OptStatus) Ptr() *CampaignUpdateStatusV2OptStatus {
	return &v
}
