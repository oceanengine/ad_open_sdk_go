/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletRelationGetV30ResponseDataResultsInner struct for SharedWalletWalletRelationGetV30ResponseDataResultsInner
type SharedWalletWalletRelationGetV30ResponseDataResultsInner struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 共享钱包ID
	SharedWalletId *int64 `json:"shared_wallet_id,omitempty"`
}