/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValue struct for SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValue
type SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValue struct {
	BasicBalanceInfo   *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueBasicBalanceInfo   `json:"basic_balance_info,omitempty"`
	GeneralBalanceInfo *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfo `json:"general_balance_info,omitempty"`
	// 共享钱包ID
	WalletId *int64 `json:"wallet_id,omitempty"`
}
