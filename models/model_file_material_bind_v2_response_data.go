/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileMaterialBindV2ResponseData
type FileMaterialBindV2ResponseData struct {
	// 推送失败列表
	FailList []*FileMaterialBindV2ResponseDataFailListInner `json:"fail_list,omitempty"`
}
