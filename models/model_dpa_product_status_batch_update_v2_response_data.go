/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductStatusBatchUpdateV2ResponseData
type DpaProductStatusBatchUpdateV2ResponseData struct {
	// 修改状态失败的商品列表
	ErrorList []*DpaProductStatusBatchUpdateV2ResponseDataErrorListInner `json:"error_list,omitempty"`
	// 修改状态成功的商品列表
	SuccessList []int64 `json:"success_list,omitempty"`
}
