/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAbTestDeleteV2ResponseData
type ToolsAbTestDeleteV2ResponseData struct {
	//
	Errors []*ToolsAbTestDeleteV2ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	Success []int64 `json:"success,omitempty"`
}
