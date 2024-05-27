/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInner struct for CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInner
type CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInner struct {
	// 减款方可转资金列表
	CapitalDetailList []*CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner `json:"capital_detail_list,omitempty"`
	// 加款方可转余额信息列表
	PayeeTransferAmountDetailList []*CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner `json:"payee_transfer_amount_detail_list,omitempty"`
	// 减款方账户id
	RemitterAccountId *int64 `json:"remitter_account_id,omitempty"`
}
