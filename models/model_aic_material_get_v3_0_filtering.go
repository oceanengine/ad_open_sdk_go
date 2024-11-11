/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AicMaterialGetV30Filtering
type AicMaterialGetV30Filtering struct {
	// 根据素材创建时间进行过滤结束始时间，可与create_start_time搭配使用 格式为yyyy-mm-dd HH:MM:SS
	CreateEndTime *string `json:"create_end_time,omitempty"`
	// 根据素材创建时间进行过滤的起始时间，可与create_end_time搭配使用 格式为yyyy-mm-dd HH:MM:SS
	CreateStartTime *string `json:"create_start_time,omitempty"`
	// 素材名称，模糊匹配
	MaterialName *string                          `json:"material_name,omitempty"`
	Radio        *AicMaterialGetV30FilteringRadio `json:"radio,omitempty"`
	// 视频ID数组，数量限制：50
	VideoIds []string `json:"video_ids,omitempty"`
}
