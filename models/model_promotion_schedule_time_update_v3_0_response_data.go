/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionScheduleTimeUpdateV30ResponseData
type PromotionScheduleTimeUpdateV30ResponseData struct {
	// 更新失败的广告ID列表
	Errors []*PromotionScheduleTimeUpdateV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 更新成功的广告ID集合
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
