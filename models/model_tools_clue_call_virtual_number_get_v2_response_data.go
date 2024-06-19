/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallVirtualNumberGetV2ResponseData
type ToolsClueCallVirtualNumberGetV2ResponseData struct {
	// 呼叫结果返回码
	CallResultCode *int64 `json:"call_result_code,omitempty"`
	// 呼叫结果返回信息
	CallResultMessage *string `json:"call_result_message,omitempty"`
	// 呼叫记录id, 用来标识一次外呼行为
	ContactId *string `json:"contact_id,omitempty"`
	// 中间号，使用该接口外呼线索号码时，显示的虚拟号（运营商提供）；<br>三天内，用户回拨中间号可以拨打给广告主，并记录对应的话单
	VirtualNumber *string `json:"virtual_number,omitempty"`
}
