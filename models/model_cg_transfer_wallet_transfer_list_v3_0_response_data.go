/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferListV30ResponseData
type CgTransferWalletTransferListV30ResponseData struct {
	// 转账记录列表
	RecordList []*CgTransferWalletTransferListV30ResponseDataRecordListInner `json:"record_list,omitempty"`
}
