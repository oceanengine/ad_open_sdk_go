/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskGetV30ResponseData
type ReportCustomAsyncTaskGetV30ResponseData struct {
	//
	List []*ReportCustomAsyncTaskGetV30ResponseDataListInner `json:"list,omitempty"`
	//
	PageInfo map[string]interface{} `json:"page_info,omitempty"`
}
