/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestCreateV2Response struct for ToolsAbTestCreateV2Response
type ToolsAbTestCreateV2Response struct {
	//
	Code *int64                           `json:"code,omitempty"`
	Data *ToolsAbTestCreateV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
