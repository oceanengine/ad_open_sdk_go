/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFileUploadV2Response struct for ClueFileUploadV2Response
type ClueFileUploadV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *ClueFileUploadV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
