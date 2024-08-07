/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskCreateV30Request struct for ReportCustomAsyncTaskCreateV30Request
type ReportCustomAsyncTaskCreateV30Request struct {
	//
	AdvertiserId int64                                   `json:"advertiser_id"`
	DataTopic    ReportCustomAsyncTaskCreateV30DataTopic `json:"data_topic"`
	//
	Dimensions []string `json:"dimensions"`
	//
	EndTime string `json:"end_time"`
	//
	Filters []*ReportCustomAsyncTaskCreateV30RequestFiltersInner `json:"filters"`
	//
	Force *bool `json:"force,omitempty"`
	//
	Metrics []string `json:"metrics"`
	//
	OrderBy []*ReportCustomAsyncTaskCreateV30RequestOrderByInner `json:"order_by"`
	//
	StartTime string `json:"start_time"`
	//
	TaskName string `json:"task_name"`
}
