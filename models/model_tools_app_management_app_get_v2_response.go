/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementAppGetV2Response struct for ToolsAppManagementAppGetV2Response
type ToolsAppManagementAppGetV2Response struct {
	//
	Code *int64                                  `json:"code,omitempty"`
	Data *ToolsAppManagementAppGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
