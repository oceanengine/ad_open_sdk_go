/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCanTransferBalanceGetV30ResponseData
type CgTransferCanTransferBalanceGetV30ResponseData struct {
	// 可转余额信息列表
	CanTransferDetailList []*CgTransferCanTransferBalanceGetV30ResponseDataCanTransferDetailListInner `json:"can_transfer_detail_list,omitempty"`
}
