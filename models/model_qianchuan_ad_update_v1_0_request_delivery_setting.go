/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestDeliverySetting
type QianchuanAdUpdateV10RequestDeliverySetting struct {
	//
	Budget float64 `json:"budget"`
	//
	CpaBid *float64 `json:"cpa_bid,omitempty"`
	//
	EndTime          *string                                              `json:"end_time,omitempty"`
	LiveScheduleType *QianchuanAdUpdateV10DeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	// 新品链路
	ProductNewOpen *bool                                        `json:"product_new_open,omitempty"`
	QcpxMode       *QianchuanAdUpdateV10DeliverySettingQcpxMode `json:"qcpx_mode,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
	//
	ScheduleFixedRange *int64 `json:"schedule_fixed_range,omitempty"`
	//
	ScheduleTime      *string                                               `json:"schedule_time,omitempty"`
	VideoScheduleType *QianchuanAdUpdateV10DeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
