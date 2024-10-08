/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageGetV2Filtering
type FileImageGetV2Filtering struct {
	// 根据图片上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
	EndTime *string `json:"end_time,omitempty"`
	// 图片高度
	Height *float64 `json:"height,omitempty"`
	// 图片ids，可以根据图片ids（创意中使用的图片key，存在一张图片对应多个image_ids的情况）进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	ImageIds []string `json:"image_ids,omitempty"`
	// 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 图片宽高比，eg: [1.7, 2.5]，输入1.7则搜索满足宽高比介于1.65-1.75之间的图片，即精度上下浮动0.05
	Ratio []float64 `json:"ratio,omitempty"`
	// md5值列表，可以根据素材的md5进行过滤 数量限制：<=100 注意：image_ids、material_ids、signatures只能选择一个进行过滤
	Signatures []string `json:"signatures,omitempty"`
	// 根据图片上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	StartTime *string `json:"start_time,omitempty"`
	// 图片宽度
	Width *float64 `json:"width,omitempty"`
}
