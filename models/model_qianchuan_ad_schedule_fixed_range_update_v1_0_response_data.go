/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdScheduleFixedRangeUpdateV10ResponseData
type QianchuanAdScheduleFixedRangeUpdateV10ResponseData struct {
	//
	AdIds []int64 `json:"ad_ids,omitempty"`
	//
	Errors []*QianchuanAdScheduleFixedRangeUpdateV10ResponseDataErrorsInner `json:"errors,omitempty"`
}
