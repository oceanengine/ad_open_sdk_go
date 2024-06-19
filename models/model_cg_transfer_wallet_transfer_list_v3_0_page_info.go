/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferListV30PageInfo
type CgTransferWalletTransferListV30PageInfo struct {
	// 页码，从1开始，仅支持page_size * page_num < 1w的深度分页
	PageNum int64 `json:"page_num"`
	// 分页大小，范围：1-100
	PageSize int64 `json:"page_size"`
}
