/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoAgentGetV2Filtering
type FileVideoAgentGetV2Filtering struct {
	// 根据视频上传时间进行过滤的截止时间，与start_time搭配使用，格式：yyyy-mm-dd
	EndTime *string `json:"end_time,omitempty"`
	// 视频高度
	Height *int64 `json:"height,omitempty"`
	//
	Labels []string `json:"labels,omitempty"`
	// 素材id列表，可以根据material_ids（素材报表使用的id，一个素材唯一对应一个素材id）进行过滤 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	MaterialIds []int64 `json:"material_ids,omitempty"`
	// 视频宽高比，示例: [1.7, 2.5] 输入1.7则搜索满足宽高比介于1.65-1.75之间的视频，即精度上下浮动0.05
	Ratio []float64 `json:"ratio,omitempty"`
	// md5值列表，可以根据素材的md5进行过滤 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	Signatures []string `json:"signatures,omitempty"`
	// 素材来源枚举 https://open.oceanengine.com/labels/7/docs/1696710760171535
	Source []string `json:"source,omitempty"`
	// 根据视频上传时间进行过滤的起始时间，与end_time搭配使用，格式：yyyy-mm-dd
	StartTime *string `json:"start_time,omitempty"`
	// 视频ids，示例: [\"86adb23eaa21229fc04ef932b5089bb8\"] 数量限制：<=100 注意：video_ids、material_ids、signatures只能选择一个进行过滤
	VideoIds []string `json:"video_ids,omitempty"`
	// 视频宽度
	Width *int64 `json:"width,omitempty"`
}
