/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalFileVideoGetV30Filtering
type LocalFileVideoGetV30Filtering struct {
	// 评估类型
	AnalysisType []*LocalFileVideoGetV30FilteringAnalysisType `json:"analysis_type,omitempty"`
	// 指标统计结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 素材类型
	ImageMode []*LocalFileVideoGetV30FilteringImageMode `json:"image_mode,omitempty"`
	// 是否过滤低质素材，默认true
	IsFilterUnqualified *bool `json:"is_filter_unqualified,omitempty"`
	// 素材来源
	MaterialSource []*LocalFileVideoGetV30FilteringMaterialSource `json:"material_source,omitempty"`
	// 搜索词
	SearchKeyWord *string `json:"search_key_word,omitempty"`
	// 指标统计开始时间
	StartTime *string `json:"start_time,omitempty"`
}
