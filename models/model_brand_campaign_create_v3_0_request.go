/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignCreateV30Request struct for BrandCampaignCreateV30Request
type BrandCampaignCreateV30Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 广告组名称
	CampaignName string `json:"campaign_name"`
	// 合同ID
	MainContractId string `json:"main_contract_id"`
	// 代理商员工id
	StaffId *int64 `json:"staff_id,omitempty"`
}
