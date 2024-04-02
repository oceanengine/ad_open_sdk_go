/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportProductHourlyAsyncTaskCreateV30RequestFiltering
type ReportProductHourlyAsyncTaskCreateV30RequestFiltering struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	DeepExternalAction *string `json:"deep_external_action,omitempty"`
	//
	DpaVideoTid *int64 `json:"dpa_video_tid,omitempty"`
	//
	ExternalAction *string `json:"external_action,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ShowAppName *string `json:"show_app_name,omitempty"`
}