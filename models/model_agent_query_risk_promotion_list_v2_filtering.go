/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AgentQueryRiskPromotionListV2Filtering
type AgentQueryRiskPromotionListV2Filtering struct {
	// 广告主账户ID，最多支持100个
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	// 广告主账户名称，模糊搜索，长度不能超过30
	AdvertiserName   *string                                                 `json:"advertiser_name,omitempty"`
	FinalOperatorTag *AgentQueryRiskPromotionListV2FilteringFinalOperatorTag `json:"final_operator_tag,omitempty"`
	// 违规素材ids，最多支持100个
	IllegalMaterialIds []int64 `json:"illegal_material_ids,omitempty"`
	// 广告ID，最多支持100个
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
	// 广告名称，模糊搜索，长度不能超过30
	PromotionName   *string                                                `json:"promotion_name,omitempty"`
	PromotionStatus *AgentQueryRiskPromotionListV2FilteringPromotionStatus `json:"promotion_status,omitempty"`
}
