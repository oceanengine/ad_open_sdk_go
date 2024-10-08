/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceUnavailableBalance 不可用余额(单位元)
type SharedWalletMainWalletGetV30ResponseDataMainWalletInfoAllocatedBalanceAdOnlyAllocatedBalanceUnavailableBalance struct {
	// 授信竞价不可用余额(单位元)
	CreditBiddingBalance *float64 `json:"credit_bidding_balance,omitempty"`
	// 授信品牌不可用余额(单位元)
	CreditBrandBalance *float64 `json:"credit_brand_balance,omitempty"`
	// 授信通用不可用余额(单位元)
	CreditGeneralBalance *float64 `json:"credit_general_balance,omitempty"`
	// 预付竞价不可用余额(单位元)
	PrepayBiddingBalance *float64 `json:"prepay_bidding_balance,omitempty"`
	// 预付品牌不可用余额(单位元)
	PrepayBrandBalance *float64 `json:"prepay_brand_balance,omitempty"`
	// 预付通用不可用余额(单位元)
	PrepayGeneralBalance *float64 `json:"prepay_general_balance,omitempty"`
}
