/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StardeliveryTaskPostEndTimeUpdateV30Request struct for StardeliveryTaskPostEndTimeUpdateV30Request
type StardeliveryTaskPostEndTimeUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	StarTaskId int64 `json:"star_task_id"`
	//
	StarTaskPostEndTime string `json:"star_task_post_end_time"`
}
