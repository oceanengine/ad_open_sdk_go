/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallVirtualNumberRefundDetailGetV2ResponseData
type ToolsClueCallVirtualNumberRefundDetailGetV2ResponseData struct {
	PageInfo *ToolsClueCallVirtualNumberRefundDetailGetV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	RefundDetailList []*ToolsClueCallVirtualNumberRefundDetailGetV2ResponseDataRefundDetailListInner `json:"refund_detail_list,omitempty"`
}
