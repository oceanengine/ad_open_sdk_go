/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceListGetV10Filtering
type QianchuanAudienceListGetV10Filtering struct {
	//
	AudienceId []int64 `json:"audience_id,omitempty"`
	//
	AudienceName *string `json:"audience_name,omitempty"`
	//
	AudienceType []*QianchuanAudienceListGetV10FilteringAudienceType `json:"audience_type,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
	//
	Status []int32 `json:"status,omitempty"`
}
