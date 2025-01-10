/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignAddV30Request struct for BrandCampaignAddV30Request
type BrandCampaignAddV30Request struct {
	// 广告主
	AdvertiserId int64 `json:"advertiser_id"`
	// 组名称
	CampaignName *string `json:"campaign_name,omitempty"`
	// 合同ID
	MainContractId *string `json:"main_contract_id,omitempty"`
	// 代理商员工ID
	StaffId *int64 `json:"staff_id,omitempty"`
}
