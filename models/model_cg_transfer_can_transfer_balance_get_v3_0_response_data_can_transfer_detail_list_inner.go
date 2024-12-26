/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInner struct for CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInner
type CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInner struct {
	// 转出方可转资金列表
	CapitalDetailList []*CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner `json:"capital_detail_list,omitempty"`
	// 转入方可转余额信息列表
	PayeeTransferAmountDetailList []*CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInnerPayeeTransferAmountDetailListInner `json:"payee_transfer_amount_detail_list,omitempty"`
	// 转出方账户id
	RemitterTargetId *int64 `json:"remitter_target_id,omitempty"`
}