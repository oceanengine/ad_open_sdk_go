/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaTemplateGetV2Response struct for DpaTemplateGetV2Response
type DpaTemplateGetV2Response struct {
	//
	Code *int64                        `json:"code,omitempty"`
	Data *DpaTemplateGetV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
