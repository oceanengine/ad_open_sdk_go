/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceLocalOnlyUnallocatedBalance 巨量本地推业务线待分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceLocalOnlyUnallocatedBalance struct {
	AvailableBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceLocalOnlyUnallocatedBalanceAvailableBalance   `json:"available_balance,omitempty"`
	UnavailableBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoUnallocatedBalanceLocalOnlyUnallocatedBalanceUnavailableBalance `json:"unavailable_balance,omitempty"`
}
