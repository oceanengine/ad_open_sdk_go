/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentCreateCommonComponentV2Response struct for StarComponentCreateCommonComponentV2Response
type StarComponentCreateCommonComponentV2Response struct {
	//
	Code *int64                                            `json:"code,omitempty"`
	Data *StarComponentCreateCommonComponentV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}