/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// RejectMaterialAiRepairAcceptTaskCreateV30ResponseData
type RejectMaterialAiRepairAcceptTaskCreateV30ResponseData struct {
	//
	AdoptAiRepairIds []int64 `json:"adopt_ai_repair_ids,omitempty"`
	//
	Errors []*RejectMaterialAiRepairAcceptTaskCreateV30ResponseDataErrorsInner `json:"errors,omitempty"`
}
