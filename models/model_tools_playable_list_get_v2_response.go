/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsPlayableListGetV2Response struct for ToolsPlayableListGetV2Response
type ToolsPlayableListGetV2Response struct {
	//
	Code *int64                              `json:"code,omitempty"`
	Data *ToolsPlayableListGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
