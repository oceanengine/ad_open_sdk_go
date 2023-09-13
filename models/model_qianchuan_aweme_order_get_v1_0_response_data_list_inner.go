/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderGetV10ResponseDataListInner struct for QianchuanAwemeOrderGetV10ResponseDataListInner
type QianchuanAwemeOrderGetV10ResponseDataListInner struct {
	// 计划id，未支付的订单无计划id
	AdId            *int64                                                         `json:"ad_id,omitempty"`
	AwemeInfo       *QianchuanAwemeOrderGetV10ResponseDataListInnerAwemeInfo       `json:"aweme_info,omitempty"`
	DeliverySetting *QianchuanAwemeOrderGetV10ResponseDataListInnerDeliverySetting `json:"delivery_setting,omitempty"`
	MarketingGoal   *QianchuanAwemeOrderGetV10DataListMarketingGoal                `json:"marketing_goal,omitempty"`
	// 订单创建时间（下单时间）
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// 订单id
	OrderId   *int64                                                   `json:"order_id,omitempty"`
	RoomInfo  *QianchuanAwemeOrderGetV10ResponseDataListInnerRoomInfo  `json:"room_info,omitempty"`
	Status    *QianchuanAwemeOrderGetV10DataListStatus                 `json:"status,omitempty"`
	VideoInfo *QianchuanAwemeOrderGetV10ResponseDataListInnerVideoInfo `json:"video_info,omitempty"`
}
