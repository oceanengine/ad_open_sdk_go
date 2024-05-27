/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner struct for CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner
type CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner struct {
	// 加款方可转资金列表
	CapitalDetailList []*CgTransferQueryCanTransferBalanceV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInnerCapitalDetailListInner `json:"capital_detail_list"`
	// 加款方账户id
	PayeeAccountId int64 `json:"payee_account_id"`
}
