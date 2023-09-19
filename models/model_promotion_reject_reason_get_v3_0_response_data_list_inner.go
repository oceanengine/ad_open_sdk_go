/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionRejectReasonGetV30ResponseDataListInner struct for PromotionRejectReasonGetV30ResponseDataListInner
type PromotionRejectReasonGetV30ResponseDataListInner struct {
	// 素材维度审核建议列表
	MaterialReject []*PromotionRejectReasonGetV30ResponseDataListInnerMaterialRejectInner `json:"material_reject,omitempty"`
	// 广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
	// 广告维度审核建议列表
	PromotionReject []*PromotionRejectReasonGetV30ResponseDataListInnerPromotionRejectInner `json:"promotion_reject,omitempty"`
}
