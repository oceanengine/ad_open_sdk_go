/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsTaskRaiseDataGetV2Response struct for ToolsTaskRaiseDataGetV2Response
type ToolsTaskRaiseDataGetV2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *ToolsTaskRaiseDataGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
