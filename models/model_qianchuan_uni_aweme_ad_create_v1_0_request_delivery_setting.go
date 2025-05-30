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
	DailyDeliveryTime *float64 `json:"daily_delivery_time,omitempty"`
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	EstimateConvert *int64 `json:"estimate_convert,omitempty"`
	//
	EstimateRoiGoal  *float64                                                     `json:"estimate_roi_goal,omitempty"`
	LiveScheduleType *QianchuanUniAwemeAdCreateV10DeliverySettingLiveScheduleType `json:"live_schedule_type,omitempty"`
	//
	MinEstimateConvert *int64 `json:"min_estimate_convert,omitempty"`
	//
	MinEstimateRoiGoal *float64                                             `json:"min_estimate_roi_goal,omitempty"`
	QcpxMode           *QianchuanUniAwemeAdCreateV10DeliverySettingQcpxMode `json:"qcpx_mode,omitempty"`
	//
	Roi2Goal     *float64                                                `json:"roi2_goal,omitempty"`
	SmartBidType QianchuanUniAwemeAdCreateV10DeliverySettingSmartBidType `json:"smart_bid_type"`
	//
	StartTime         *string                                                       `json:"start_time,omitempty"`
	VideoScheduleType *QianchuanUniAwemeAdCreateV10DeliverySettingVideoScheduleType `json:"video_schedule_type,omitempty"`
}
