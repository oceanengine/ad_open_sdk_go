/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// RejectMaterialAiRepairGetV30Response struct for RejectMaterialAiRepairGetV30Response
type RejectMaterialAiRepairGetV30Response struct {
	//
	Code *int64                                    `json:"code,omitempty"`
	Data *RejectMaterialAiRepairGetV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
