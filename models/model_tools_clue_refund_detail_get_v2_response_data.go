/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueRefundDetailGetV2ResponseData
type ToolsClueRefundDetailGetV2ResponseData struct {
	PageInfo *ToolsClueRefundDetailGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	RefundDetailList []*ToolsClueRefundDetailGetV2ResponseDataRefundDetailListInner `json:"refund_detail_list,omitempty"`
}
