/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2ResponseData
type StarDemandCreateAssignV2ResponseData struct {
	BillInfo *StarDemandCreateAssignV2ResponseDataBillInfo `json:"bill_info,omitempty"`
	// 需求ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 任务ID列表
	OrderIds []int64 `json:"order_ids,omitempty"`
}
