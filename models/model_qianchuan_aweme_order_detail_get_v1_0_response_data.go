/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeOrderDetailGetV10ResponseData
type QianchuanAwemeOrderDetailGetV10ResponseData struct {
	// 计划id，未支付的订单无计划id
	AdId          *int64                                                    `json:"ad_id,omitempty"`
	AddAmountInfo *QianchuanAwemeOrderDetailGetV10ResponseDataAddAmountInfo `json:"add_amount_info,omitempty"`
	Audience      *QianchuanAwemeOrderDetailGetV10ResponseDataAudience      `json:"audience,omitempty"`
	AuditRecord   *QianchuanAwemeOrderDetailGetV10ResponseDataAuditRecord   `json:"audit_record,omitempty"`
	AwemeInfo     *QianchuanAwemeOrderDetailGetV10ResponseDataAwemeInfo     `json:"aweme_info,omitempty"`
	CouponInfo    *QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfo    `json:"coupon_info,omitempty"`
	// 批量优惠券信息
	CouponInfos     []*QianchuanAwemeOrderDetailGetV10ResponseDataCouponInfosInner `json:"coupon_infos,omitempty"`
	DeliverySetting *QianchuanAwemeOrderDetailGetV10ResponseDataDeliverySetting    `json:"delivery_setting,omitempty"`
	MarketingGoal   *QianchuanAwemeOrderDetailGetV10DataMarketingGoal              `json:"marketing_goal,omitempty"`
	// 订单创建时间
	OrderCreateTime *string `json:"order_create_time,omitempty"`
	// 订单ID
	OrderId     *int64                                                  `json:"order_id,omitempty"`
	ProductInfo *QianchuanAwemeOrderDetailGetV10ResponseDataProductInfo `json:"product_info,omitempty"`
	RoomInfo    *QianchuanAwemeOrderDetailGetV10ResponseDataRoomInfo    `json:"room_info,omitempty"`
	Status      *QianchuanAwemeOrderDetailGetV10DataStatus              `json:"status,omitempty"`
	VideoInfo   *QianchuanAwemeOrderDetailGetV10ResponseDataVideoInfo   `json:"video_info,omitempty"`
}
