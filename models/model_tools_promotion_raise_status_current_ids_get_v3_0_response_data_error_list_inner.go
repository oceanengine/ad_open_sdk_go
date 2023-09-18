/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataErrorListInner struct for ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataErrorListInner
type ToolsPromotionRaiseStatusCurrentIdsGetV30ResponseDataErrorListInner struct {
	// 获取一键起量状态失败的错误原因
	ErrorReason *string `json:"error_reason,omitempty"`
	// 获取一键起量状态失败的广告ID
	PromotionId *int64 `json:"promotion_id,omitempty"`
}