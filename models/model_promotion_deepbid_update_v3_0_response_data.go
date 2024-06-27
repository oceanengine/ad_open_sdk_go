/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionDeepbidUpdateV30ResponseData
type PromotionDeepbidUpdateV30ResponseData struct {
	// 更新失败的广告列表
	Errors []*PromotionDeepbidUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的广告ID集合
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
