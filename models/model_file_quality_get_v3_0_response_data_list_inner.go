/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileQualityGetV30ResponseDataListInner struct for FileQualityGetV30ResponseDataListInner
type FileQualityGetV30ResponseDataListInner struct {
	//
	Errmsg string `json:"errmsg"`
	// 是否是首发素材
	IsFirstPublishMaterial *bool `json:"is_first_publish_material,omitempty"`
	// 是否同质化素材风险-未投放预计排队素材
	IsSimilarExpectedQueueMaterial *bool `json:"is_similar_expected_queue_material,omitempty"`
	// 是否同质化素材风险-排队投放素材
	IsSimilarQueueMaterial *bool `json:"is_similar_queue_material,omitempty"`
	//
	MaterialId int64                           `json:"material_id"`
	Status     FileQualityGetV30DataListStatus `json:"status"`
}
