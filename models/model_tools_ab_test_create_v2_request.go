/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestCreateV2Request struct for ToolsAbTestCreateV2Request
type ToolsAbTestCreateV2Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 实验对象列表
	ObjectList []*ToolsAbTestCreateV2RequestObjectListInner `json:"object_list"`
	// 实验结束时间，格式为：\"2020-12-25 07:12:08\"，
	TestEndTime string `json:"test_end_time"`
	// 实验名称，最多100个字符，中文算两个字符，不支持emoji，不支持与现有实验重名。
	TestName string `json:"test_name"`
	// 实验开始时间，格式为：\"2020-12-25 07:12:08\"，开始时间不能早于当前时间。
	TestStartTime string                      `json:"test_start_time"`
	TestType      ToolsAbTestCreateV2TestType `json:"test_type"`
}
