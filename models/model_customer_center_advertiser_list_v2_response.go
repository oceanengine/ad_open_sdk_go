/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CustomerCenterAdvertiserListV2Response struct for CustomerCenterAdvertiserListV2Response
type CustomerCenterAdvertiserListV2Response struct {
	//
	Code *int64                                      `json:"code,omitempty"`
	Data *CustomerCenterAdvertiserListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
