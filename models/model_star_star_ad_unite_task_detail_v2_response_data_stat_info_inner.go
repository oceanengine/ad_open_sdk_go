/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarStarAdUniteTaskDetailV2ResponseDataStatInfoInner struct for StarStarAdUniteTaskDetailV2ResponseDataStatInfoInner
type StarStarAdUniteTaskDetailV2ResponseDataStatInfoInner struct {
	// 安卓转化数
	AndroidConvert int64 `json:"android_convert"`
	// 安卓最新出价，单位：元*1000，建议获取后除以1000展示为元单位
	AndroidConvertUnitAmount int64 `json:"android_convert_unit_amount"`
	// 安卓消耗，单位：元*100000，建议获取后除以100000展示为元单位
	AndroidCost int64 `json:"android_cost"`
	// 安卓深度转化数
	AndroidDeepConvert int64 `json:"android_deep_convert"`
	// iOS转化数
	IosConvert int64 `json:"ios_convert"`
	// iOS最新出价，单位：元*1000，建议获取后除以1000展示为元单位
	IosConvertUnitAmount int64 `json:"ios_convert_unit_amount"`
	// iOS消耗，单位：元*100000，建议获取后除以100000展示为元单位
	IosCost int64 `json:"ios_cost"`
	// iOS深度转化数
	IosDeepConvert int64 `json:"ios_deep_convert"`
	// 数据日期，只和安卓/iOS消耗、转化数、深度转化数相关
	StatDate string `json:"stat_date"`
}
