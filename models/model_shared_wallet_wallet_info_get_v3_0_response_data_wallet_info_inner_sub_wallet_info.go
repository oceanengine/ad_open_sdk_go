/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerSubWalletInfo 投放钱包(小钱包)信息
type SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerSubWalletInfo struct {
	// 钱包下的 adv 数量
	AdvCnt *int64 `json:"adv_cnt,omitempty"`
	// 所属大钱包ID
	MainWalletId   *int64                                                                              `json:"main_wallet_id,omitempty"`
	SubSharedRange *SharedWalletWalletInfoGetV30ResponseDataWalletInfoInnerSubWalletInfoSubSharedRange `json:"sub_shared_range,omitempty"`
}
