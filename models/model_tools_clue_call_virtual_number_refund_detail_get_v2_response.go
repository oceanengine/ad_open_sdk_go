/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallVirtualNumberRefundDetailGetV2Response struct for ToolsClueCallVirtualNumberRefundDetailGetV2Response
type ToolsClueCallVirtualNumberRefundDetailGetV2Response struct {
	//
	Code *int64                                                   `json:"code,omitempty"`
	Data *ToolsClueCallVirtualNumberRefundDetailGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
