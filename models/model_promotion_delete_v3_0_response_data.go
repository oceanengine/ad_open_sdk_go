/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionDeleteV30ResponseData
type PromotionDeleteV30ResponseData struct {
	// 删除失败的广告ID列表
	Errors []*PromotionBidUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 删除成功的广告ID列表
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
