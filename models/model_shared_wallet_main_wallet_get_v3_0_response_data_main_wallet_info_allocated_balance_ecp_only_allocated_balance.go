/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceEcpOnlyAllocatedBalance 巨量千川业务线已分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceEcpOnlyAllocatedBalance struct {
	AvailableBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceEcpOnlyAllocatedBalanceAvailableBalance   `json:"available_balance,omitempty"`
	UnavailableBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceEcpOnlyAllocatedBalanceUnavailableBalance `json:"unavailable_balance,omitempty"`
}
