/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanBatchCampaignStatusUpdateV10ResponseData
type QianchuanBatchCampaignStatusUpdateV10ResponseData struct {
	//
	Errors []*QianchuanBatchCampaignStatusUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
	//
	Success []int64 `json:"success,omitempty"`
}
