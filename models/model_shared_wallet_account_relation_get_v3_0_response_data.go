/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletAccountRelationGetV30ResponseData
type SharedWalletAccountRelationGetV30ResponseData struct {
	// 客户公司ID
	CompanyId *int64 `json:"company_id,omitempty"`
	// 共享钱包ID
	MainWalletId *int64 `json:"main_wallet_id,omitempty"`
	// 子钱包ID列表
	SubWalletIds []int64 `json:"sub_wallet_ids,omitempty"`
}
