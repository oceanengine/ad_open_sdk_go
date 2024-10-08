/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueInfoUpdateV2Request struct for ToolsClueInfoUpdateV2Request
type ToolsClueInfoUpdateV2Request struct {
	// 修改线索所属的广告主id
	AdvertiserIds []int64 `json:"advertiser_ids"`
	// 操作人，不能为空
	OpName *string `json:"op_name,omitempty"`
	// 更新的线索信息
	UpdateInfo []*ToolsClueInfoUpdateV2RequestUpdateInfoInner `json:"update_info"`
}
