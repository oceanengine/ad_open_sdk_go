/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCreativeWordSelectV2Response struct for ToolsCreativeWordSelectV2Response
type ToolsCreativeWordSelectV2Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *ToolsCreativeWordSelectV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
