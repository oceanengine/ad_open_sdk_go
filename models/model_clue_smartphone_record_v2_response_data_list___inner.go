/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneRecordV2ResponseDataListInner struct for ClueSmartphoneRecordV2ResponseDataListInner
type ClueSmartphoneRecordV2ResponseDataListInner struct {
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	CallDuration *int64 `json:"call_duration,omitempty"`
	//
	ClueId *int64 `json:"clue_id,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	InstanceId *int64 `json:"instance_id,omitempty"`
	//
	RealDuration *int64 `json:"real_duration,omitempty"`
	//
	SiteId *int64 `json:"site_id,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
	//
	VirtualNumber *string `json:"virtual_number,omitempty"`
}
