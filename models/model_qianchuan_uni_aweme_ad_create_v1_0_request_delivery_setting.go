/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniAwemeAdCreateV10RequestDeliverySetting
type QianchuanUniAwemeAdCreateV10RequestDeliverySetting struct {
	//
	Budget float64 `json:"budget"`
	//
	EndTime          *string                                                      `json:"end_time,omitempty"`
	LiveScheduleType *QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	QcpxMode         *QianchuanUniAwemeAdCreateV10DeliverySettingQcpxMode         `json:"qcpx_mode,omitempty"`
	//
	Roi2Goal     *float64                                                `json:"roi2_goal,omitempty"`
	SmartBidType QianchuanUniAwemeAdCreateV10DeliverySettingSmartBidType `json:"smart_bid_type"`
	//
	StartTime         *string                                                       `json:"start_time,omitempty"`
	VideoScheduleType *QianchuanUniAwemeAdCreateV10DeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
