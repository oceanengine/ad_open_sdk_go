/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordBatchGetV30Request struct for ToolsPrivativeWordBatchGetV30Request
type ToolsPrivativeWordBatchGetV30Request struct {
	// 广告主id
	AdvertiserId int64 `json:"advertiser_id"`
	// 项目id列表
	ProjectIds []int64 `json:"project_ids,omitempty"`
	//
	PromotionIds []int64 `json:"promotion_ids,omitempty"`
}
