/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdScheduleFixedRangeUpdateV10Request struct for QianchuanAdScheduleFixedRangeUpdateV10Request
type QianchuanAdScheduleFixedRangeUpdateV10Request struct {
	//
	AdIds []int64 `json:"ad_ids"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ScheduleFixedRange int64 `json:"schedule_fixed_range"`
}
