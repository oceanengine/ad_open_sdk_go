/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCanTransferTargetListV30ResponseData
type CgTransferCanTransferTargetListV30ResponseData struct {
	// 可转账户列表
	CanTransferTargetList []*CgTransferCanTransferTargetListV30ResponseDataCanTransferTargetListInner `json:"can_transfer_target_list,omitempty"`
	PageInfo              *CgTransferCanTransferTargetListV30ResponseDataPageInfo                     `json:"page_info,omitempty"`
}