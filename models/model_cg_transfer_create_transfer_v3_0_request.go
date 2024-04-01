/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCreateTransferV30Request struct for CgTransferCreateTransferV30Request
type CgTransferCreateTransferV30Request struct {
	// 锚定账户id，1:N的1
	AccountId int64 `json:"account_id"`
	// 代理商账户id，用于鉴权
	AgentId int64 `json:"agent_id"`
	// 请求幂等id，同一biz_request_no发起转账代表同一转账申请，返回的转账单号相同，推荐uuid
	BizRequestNo string `json:"biz_request_no"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 目标账户列表，1:N的N，需要列表内账户类型相同，最多支持100个
	TargetAccountDetailList []*CgTransferCreateTransferV30RequestTargetAccountDetailListInner `json:"target_account_detail_list"`
	TransferDirection       CgTransferCreateTransferV30TransferDirection                      `json:"transfer_direction"`
}
