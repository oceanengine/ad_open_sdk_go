/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAudienceAwemeListV2Response struct for ReportAudienceAwemeListV2Response
type ReportAudienceAwemeListV2Response struct {
	//
	Code *int64                                 `json:"code,omitempty"`
	Data *ReportAudienceAwemeListV2ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
