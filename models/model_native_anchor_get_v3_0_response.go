/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetV30Response struct for NativeAnchorGetV30Response
type NativeAnchorGetV30Response struct {
	//
	Code *int64                          `json:"code,omitempty"`
	Data *NativeAnchorGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
