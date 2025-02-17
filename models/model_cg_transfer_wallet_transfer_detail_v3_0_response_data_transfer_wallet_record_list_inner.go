/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInner struct for CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInner
type CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInner struct {
	// 大钱包id
	MainWalletId *int64 `json:"main_wallet_id,omitempty"`
	// 小钱包id
	SubWalletId *int64 `json:"sub_wallet_id,omitempty"`
	// 转账金额(单位：分)
	TransferAmount *int64 `json:"transfer_amount,omitempty"`
	// 转账资金类型列表
	TransferCapitalRecordList []*CgTransferWalletTransferDetailV30ResponseDataTransferWalletRecordListInnerTransferCapitalRecordListInner `json:"transfer_capital_record_list,omitempty"`
	TransferStatus            *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferStatus                                `json:"transfer_status,omitempty"`
}
