/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfo 通用余额信息
type SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfo struct {
	AdOnlyBalanceInfo    *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoAdOnlyBalanceInfo    `json:"ad_only_balance_info,omitempty"`
	AdSharedBalanceInfo  *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoAdSharedBalanceInfo  `json:"ad_shared_balance_info,omitempty"`
	EcpOnlyBalanceInfo   *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoEcpOnlyBalanceInfo   `json:"ecp_only_balance_info,omitempty"`
	LocalOnlyBalanceInfo *SharedWalletWalletBalanceGetV30ResponseDataSharedWalletBalanceInfoValueGeneralBalanceInfoLocalOnlyBalanceInfo `json:"local_only_balance_info,omitempty"`
}
