/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner struct for CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner
type CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInnerCapitalDetailListInner struct {
	CapitalType CgTransferCanTransferBalanceGetV30DataCanTransferDetailListCapitalDetailListCapitalType `json:"capital_type"`
	// 转出方可转资金余额(单位：分)
	TransferBalance int64 `json:"transfer_balance"`
}
