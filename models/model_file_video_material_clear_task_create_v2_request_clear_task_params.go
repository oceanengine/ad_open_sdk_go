/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileVideoMaterialClearTaskCreateV2RequestClearTaskParams 任务参数
type FileVideoMaterialClearTaskCreateV2RequestClearTaskParams struct {
	// 清理素材类型， 允许值： INEFFICIENT_MATERIAL低效素材； SIMILAR_MATERIAL 同质化挤压严重素材； SIMILAR_QUEUE_MATERIAL 同质化排队素材;
	ClearMaterialTypes []*FileVideoMaterialClearTaskCreateV2ClearTaskParamsClearMaterialTypes `json:"clear_material_types,omitempty"`
	// 累积转化数@PM补充描述是UPPER值
	Convert *int64 `json:"convert,omitempty"`
	// 累积消耗@PM补充描述是UPPER值
	Cost *float64 `json:"cost,omitempty"`
	// 清理创建时间\"yyyy-mm-dd\"之前的素材，包括这一天
	CreateTimeUpper *string `json:"create_time_upper,omitempty"`
	// 数据统计结束时间\"yyyy-mm-dd\"
	EndTime *string `json:"end_time,omitempty"`
	// 待清理素材列表，单次最多支持100个
	MaterialIds    []int64                                                         `json:"material_ids,omitempty"`
	MaterialSource FileVideoMaterialClearTaskCreateV2ClearTaskParamsMaterialSource `json:"material_source"`
	// 数据统计开始时间\"yyyy-mm-dd\"
	StartTime *string `json:"start_time,omitempty"`
}
