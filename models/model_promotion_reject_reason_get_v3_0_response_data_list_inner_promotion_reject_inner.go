/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionRejectReasonGetV30ResponseDataListInnerPromotionRejectInner struct for PromotionRejectReasonGetV30ResponseDataListInnerPromotionRejectInner
type PromotionRejectReasonGetV30ResponseDataListInnerPromotionRejectInner struct {
	Content *PromotionRejectReasonGetV30DataListPromotionRejectContent `json:"content,omitempty"`
	// 拒绝建议
	RejectReason []string `json:"reject_reason,omitempty"`
	// 审核建议
	Suggestion []string `json:"suggestion,omitempty"`
}
