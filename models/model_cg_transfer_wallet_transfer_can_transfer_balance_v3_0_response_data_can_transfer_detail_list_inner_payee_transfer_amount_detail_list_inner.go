/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner struct for CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner
type CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner struct {
	// 加款钱包可转资金列表
	CapitalDetailList []*CgTransferWalletTransferCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInnerCapitalDetailListInner `json:"capital_detail_list,omitempty"`
	// 加款钱包非品牌资金最小转入金额（单位：分）
	NonBrandMinTransferBalance *int64 `json:"non_brand_min_transfer_balance,omitempty"`
	// 加款钱包id
	PayeeWalletId *int64 `json:"payee_wallet_id,omitempty"`
}
