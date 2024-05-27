/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCreateTransferV30RequestTargetAccountDetailListInner struct for CgTransferCreateTransferV30RequestTargetAccountDetailListInner
type CgTransferCreateTransferV30RequestTargetAccountDetailListInner struct {
	// 目标账户id
	AccountId int64 `json:"account_id"`
	// 锚定账户与目标账户转账资金列表
	TransferCapitalDetailList []*CgTransferCreateTransferV30RequestTargetAccountDetailListInnerTransferCapitalDetailListInner `json:"transfer_capital_detail_list"`
}
