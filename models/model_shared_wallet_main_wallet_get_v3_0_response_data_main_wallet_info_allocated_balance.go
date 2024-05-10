/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalance 钱包已分配余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalance struct {
	AdOnlyAllocatedBalance    *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalance    `json:"ad_only_allocated_balance,omitempty"`
	AdSharedAllocatedBalance  *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdSharedAllocatedBalance  `json:"ad_shared_allocated_balance,omitempty"`
	EcpOnlyAllocatedBalance   *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceEcpOnlyAllocatedBalance   `json:"ecp_only_allocated_balance,omitempty"`
	LocalOnlyAllocatedBalance *SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceLocalOnlyAllocatedBalance `json:"local_only_allocated_balance,omitempty"`
}
