/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderCreateV10RequestDeliverySetting 投放设置
type QianchuanAwemeOrderCreateV10RequestDeliverySetting struct {
	//
	Amount  int64                                               `json:"amount"`
	BidMode QianchuanAwemeOrderCreateV10DeliverySettingBidMode  `json:"bid_mode"`
	BidType *QianchuanAwemeOrderCreateV10DeliverySettingBidType `json:"bid_type,omitempty"`
	//
	CpaBid *float64 `json:"cpa_bid,omitempty"`
	//
	DeliveryTime     float64                                                      `json:"delivery_time"`
	ExternalAction   QianchuanAwemeOrderCreateV10DeliverySettingExternalAction    `json:"external_action"`
	LiveroomHeatMode *QianchuanAwemeOrderCreateV10DeliverySettingLiveroomHeatMode `json:"liveroom_heat_mode,omitempty"`
	//
	RoiGoal *float64 `json:"roi_goal,omitempty"`
}
