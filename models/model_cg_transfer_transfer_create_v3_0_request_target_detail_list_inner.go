/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferTransferCreateV30RequestTargetDetailListInner struct for CgTransferTransferCreateV30RequestTargetDetailListInner
type CgTransferTransferCreateV30RequestTargetDetailListInner struct {
	// 目标账户id
	TargetId int64 `json:"target_id"`
	// 锚定账户与目标账户转账资金列表
	TransferCapitalDetailList []*CgTransferTransferCreateV30RequestTargetDetailListInnerTransferCapitalDetailListInner `json:"transfer_capital_detail_list"`
}
