/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsCountryInfoV2Response struct for ToolsCountryInfoV2Response
type ToolsCountryInfoV2Response struct {
	//
	Code *int64                          `json:"code,omitempty"`
	Data *ToolsCountryInfoV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
