/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarComponentQueryLinkV2Response struct for StarComponentQueryLinkV2Response
type StarComponentQueryLinkV2Response struct {
	//
	Code *int64                                `json:"code,omitempty"`
	Data *StarComponentQueryLinkV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
