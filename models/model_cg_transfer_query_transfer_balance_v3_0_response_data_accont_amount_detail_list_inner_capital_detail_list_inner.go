/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferQueryTransferBalanceV30ResponseDataAccontAmountDetailListInnerCapitalDetailListInner struct for CgTransferQueryTransferBalanceV30ResponseDataAccontAmountDetailListInnerCapitalDetailListInner
type CgTransferQueryTransferBalanceV30ResponseDataAccontAmountDetailListInnerCapitalDetailListInner struct {
	CapitalType CgTransferQueryTransferBalanceV30DataAccontAmountDetailListCapitalDetailListCapitalType `json:"capital_type"`
	// 可转资金金额(单位：分)
	TransferBalance int64 `json:"transfer_balance"`
}
