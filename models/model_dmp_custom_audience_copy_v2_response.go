/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpCustomAudienceCopyV2Response struct for DmpCustomAudienceCopyV2Response
type DmpCustomAudienceCopyV2Response struct {
	//
	Code *int64                               `json:"code,omitempty"`
	Data *DmpCustomAudienceCopyV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
