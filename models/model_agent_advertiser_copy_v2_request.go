/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentAdvertiserCopyV2Request struct for AgentAdvertiserCopyV2Request
type AgentAdvertiserCopyV2Request struct {
	// 被复制广告主账户ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 代理商账户ID
	AgentId int64 `json:"agent_id"`
	// 是否复制投放资质，默认不复制。原广告主账户的投放资质数量>0且<=100时才会复制。
	CopyDelivery *bool `json:"copy_delivery,omitempty"`
	// 是否复制原广告主账户的账户标签
	CopyTag *bool `json:"copy_tag,omitempty"`
	// 复制账户信息
	Item       []*AgentAdvertiserCopyV2RequestItemInner `json:"item"`
	ReportType *AgentAdvertiserCopyV2ReportType         `json:"report_type,omitempty"`
}
