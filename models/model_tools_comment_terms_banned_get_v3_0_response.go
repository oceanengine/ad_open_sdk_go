/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCommentTermsBannedGetV30Response struct for ToolsCommentTermsBannedGetV30Response
type ToolsCommentTermsBannedGetV30Response struct {
	//
	Code *int64                                     `json:"code,omitempty"`
	Data *ToolsCommentTermsBannedGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
