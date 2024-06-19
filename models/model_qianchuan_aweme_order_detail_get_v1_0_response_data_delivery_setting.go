/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseDataDeliverySetting 投放设置
type QianchuanAwemeOrderDetailGetV10ResponseDataDeliverySetting struct {
	// 投放金额
	Amount  *int64                                                     `json:"amount,omitempty"`
	BidMode *QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidMode `json:"bid_mode,omitempty"`
	BidType *QianchuanAwemeOrderDetailGetV10DataDeliverySettingBidType `json:"bid_type,omitempty"`
	// 手动出价金额
	CpaBid *float64 `json:"cpa_bid,omitempty"`
	// 期望曝光时长
	DeliveryTime     *float64                                                            `json:"delivery_time,omitempty"`
	ExternalAction   *QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction   `json:"external_action,omitempty"`
	LiveroomHeatMode *QianchuanAwemeOrderDetailGetV10DataDeliverySettingLiveroomHeatMode `json:"liveroom_heat_mode,omitempty"`
	// 实际支付金额
	PaymentAmount *float64 `json:"payment_amount,omitempty"`
	// 支付ROI目标
	RoiGoal *float64 `json:"roi_goal,omitempty"`
}
