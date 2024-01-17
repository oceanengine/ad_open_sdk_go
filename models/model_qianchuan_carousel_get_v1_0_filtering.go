/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselGetV10Filtering
type QianchuanCarouselGetV10Filtering struct {
	// 根据图片上传时间进行过滤的截止时间，与start_time搭配使用，格式：\"yyyy-mm-dd\"
	EndTime *string `json:"end_time,omitempty"`
	// 素材类型 CAROUSEL 图文
	ImageMode []*QianchuanCarouselGetV10FilteringImageMode `json:"image_mode,omitempty"`
	// 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤 数量限制：<=100
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 支持根据图文名称/ID搜索
	QueryString *string `json:"query_string,omitempty"`
	// 图文素材来源 JICHAUNG 即创
	Sources []*QianchuanCarouselGetV10FilteringSources `json:"sources,omitempty"`
	// 根据图片上传时间进行过滤的起始时间，与end_time搭配使用，格式：\"yyyy-mm-dd\"
	StartTime *string `json:"start_time,omitempty"`
}
