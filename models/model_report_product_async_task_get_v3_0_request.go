/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportProductAsyncTaskGetV30Request struct for ReportProductAsyncTaskGetV30Request
type ReportProductAsyncTaskGetV30Request struct {
	//
	AdvertiserId *int64                                        `json:"advertiser_id,omitempty"`
	Filtering    *ReportProductAsyncTaskGetV30RequestFiltering `json:"filtering,omitempty"`
}
