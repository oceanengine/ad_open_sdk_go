/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferCreateV30RequestTargetWalletDetailListInnerTransferCapitalDetailListInner struct for CgTransferWalletTransferCreateV30RequestTargetWalletDetailListInnerTransferCapitalDetailListInner
type CgTransferWalletTransferCreateV30RequestTargetWalletDetailListInnerTransferCapitalDetailListInner struct {
	CapitalType CgTransferWalletTransferCreateV30TargetWalletDetailListTransferCapitalDetailListCapitalType `json:"capital_type"`
	Platform    CgTransferWalletTransferCreateV30TargetWalletDetailListTransferCapitalDetailListPlatform    `json:"platform"`
	// 申请转账金额（单位：分）
	TransferAmount int64 `json:"transfer_amount"`
}
