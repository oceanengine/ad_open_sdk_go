/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletRelationGetV30ResponseData
type SharedWalletWalletRelationGetV30ResponseData struct {
	PageInfo *SharedWalletWalletRelationGetV30ResponseDataPageInfo `json:"page_info,omitempty"`
	// 共享钱包关系
	Results []*SharedWalletWalletRelationGetV30ResponseDataResultsInner `json:"results,omitempty"`
	// 共享钱包ID
	SharedWalletId *int64 `json:"shared_wallet_id,omitempty"`
}
