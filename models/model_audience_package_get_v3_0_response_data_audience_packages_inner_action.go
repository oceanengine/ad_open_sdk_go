/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV30ResponseDataAudiencePackagesInnerAction
type AudiencePackageGetV30ResponseDataAudiencePackagesInnerAction struct {
	// 行为分类
	ActionCategories []int64 `json:"action_categories,omitempty"`
	// 行为天数
	ActionDays *int64 `json:"action_days,omitempty"`
	// 行为关键词
	ActionWords []int64 `json:"action_words,omitempty"`
}
