/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDeleteV2Response struct for ClueFormDeleteV2Response
type ClueFormDeleteV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *ClueFormDeleteV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
