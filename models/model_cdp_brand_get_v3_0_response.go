/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CdpBrandGetV30Response struct for CdpBrandGetV30Response
type CdpBrandGetV30Response struct {
	//
	Code *int64                      `json:"code,omitempty"`
	Data *CdpBrandGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
