/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsClueCallVirtualNumberGetV2Response struct for ToolsClueCallVirtualNumberGetV2Response
type ToolsClueCallVirtualNumberGetV2Response struct {
	//
	Code *int64                                       `json:"code,omitempty"`
	Data *ToolsClueCallVirtualNumberGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
