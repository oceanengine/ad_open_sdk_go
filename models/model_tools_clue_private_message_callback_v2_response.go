/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCluePrivateMessageCallbackV2Response struct for ToolsCluePrivateMessageCallbackV2Response
type ToolsCluePrivateMessageCallbackV2Response struct {
	//
	Code *int64                                         `json:"code,omitempty"`
	Data *ToolsCluePrivateMessageCallbackV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
