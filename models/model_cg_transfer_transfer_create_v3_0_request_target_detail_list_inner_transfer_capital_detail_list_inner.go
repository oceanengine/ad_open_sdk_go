/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferTransferCreateV30RequestTargetDetailListInnerTransferCapitalDetailListInner struct for CgTransferTransferCreateV30RequestTargetDetailListInnerTransferCapitalDetailListInner
type CgTransferTransferCreateV30RequestTargetDetailListInnerTransferCapitalDetailListInner struct {
	CapitalType CgTransferTransferCreateV30TargetDetailListTransferCapitalDetailListCapitalType `json:"capital_type"`
	// 转账资金金额(单位：分)
	TransferAmount int64 `json:"transfer_amount"`
}
