/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportStardeliveryTaskVideoDataGetV30ResponseDataListInner struct for ReportStardeliveryTaskVideoDataGetV30ResponseDataListInner
type ReportStardeliveryTaskVideoDataGetV30ResponseDataListInner struct {
	//
	StarTaskId      *int64                                                        `json:"star_task_id,omitempty"`
	StarVideoStatus *ReportStardeliveryTaskVideoDataGetV30DataListStarVideoStatus `json:"star_video_status,omitempty"`
}
