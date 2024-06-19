/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferListV30ResponseDataRecordListInner struct for CgTransferWalletTransferListV30ResponseDataRecordListInner
type CgTransferWalletTransferListV30ResponseDataRecordListInner struct {
	// 转账申请幂等id
	BizRequestNo *string `json:"biz_request_no,omitempty"`
	// 加款方id
	PayeeId *int64 `json:"payee_id,omitempty"`
	// 转账备注
	Remark *string `json:"remark,omitempty"`
	// 减款方id
	RemitterId *int64 `json:"remitter_id,omitempty"`
	// 转账资金列表
	TransferCapitalRecordList []*CgTransferWalletTransferListV30ResponseDataRecordListInnerTransferCapitalRecordListInner `json:"transfer_capital_record_list,omitempty"`
	// 转账单编号
	TransferOrderSerial *string `json:"transfer_order_serial,omitempty"`
	// 加款方-减款方间转账金额(单位：分)
	TransferTargetAmount *int64 `json:"transfer_target_amount,omitempty"`
	// 完成时间yyyy-MM-dd HH:mm:ss
	TransferTargetCreateTime *string `json:"transfer_target_create_time,omitempty"`
	// 创建时间yyyy-MM-dd HH:mm:ss
	TransferTargetFinishTime *string                                                            `json:"transfer_target_finish_time,omitempty"`
	TransferTargetStatus     *CgTransferWalletTransferListV30DataRecordListTransferTargetStatus `json:"transfer_target_status,omitempty"`
}
