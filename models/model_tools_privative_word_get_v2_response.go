/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPrivativeWordGetV2Response struct for ToolsPrivativeWordGetV2Response
type ToolsPrivativeWordGetV2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *ToolsPrivativeWordGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
