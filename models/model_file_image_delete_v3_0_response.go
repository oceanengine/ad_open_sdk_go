/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageDeleteV30Response struct for FileImageDeleteV30Response
type FileImageDeleteV30Response struct {
	//
	Code *int64                          `json:"code,omitempty"`
	Data *FileImageDeleteV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
