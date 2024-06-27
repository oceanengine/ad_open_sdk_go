/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportMaterialGetV10ResponseDataListInnerFields
type QianchuanReportMaterialGetV10ResponseDataListInnerFields struct {
	//
	AdAllOrderGmvSettleRate14d *float64 `json:"ad_all_order_gmv_settle_rate_14d,omitempty"`
	//
	AdAllOrderGmvSettleRate7d *float64 `json:"ad_all_order_gmv_settle_rate_7d,omitempty"`
	//
	AdLiveOrderSettleCost14d *float64 `json:"ad_live_order_settle_cost_14d,omitempty"`
	//
	AdLiveOrderSettleCost7d *float64 `json:"ad_live_order_settle_cost_7d,omitempty"`
	//
	AdLiveOrderSettleRoi14d *float64 `json:"ad_live_order_settle_roi_14d,omitempty"`
	//
	AdLiveOrderSettleRoi7d *float64 `json:"ad_live_order_settle_roi_7d,omitempty"`
	//
	AveragePlayTimePerPlay *float64 `json:"average_play_time_per_play,omitempty"`
	//
	ClickCnt *float64 `json:"click_cnt,omitempty"`
	//
	ConvertCnt *float64 `json:"convert_cnt,omitempty"`
	//
	ConvertRate *float64 `json:"convert_rate,omitempty"`
	//
	CpaPlatform *float64 `json:"cpa_platform,omitempty"`
	//
	Cpm *float64 `json:"cpm,omitempty"`
	//
	CpmPlatform *float64 `json:"cpm_platform,omitempty"`
	//
	CreateOrderAmount *float64 `json:"create_order_amount,omitempty"`
	//
	CreateOrderCount *float64 `json:"create_order_count,omitempty"`
	//
	CreateOrderCouponAmount *float64 `json:"create_order_coupon_amount,omitempty"`
	//
	CreateOrderRoi *float64 `json:"create_order_roi,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DislikeCnt *float64 `json:"dislike_cnt,omitempty"`
	//
	DyComment *float64 `json:"dy_comment,omitempty"`
	//
	DyFollow *float64 `json:"dy_follow,omitempty"`
	//
	DyLike *float64 `json:"dy_like,omitempty"`
	//
	DyShare *float64 `json:"dy_share,omitempty"`
	//
	EcpConvertCnt *float64 `json:"ecp_convert_cnt,omitempty"`
	//
	EcpConvertRate *float64 `json:"ecp_convert_rate,omitempty"`
	//
	EcpCpaPlatform *float64 `json:"ecp_cpa_platform,omitempty"`
	//
	IndirectOrderUnfinishedEstimateGmv7days *float64 `json:"indirect_order_unfinished_estimate_gmv_7days,omitempty"`
	//
	LiveFansClubJoinCnt *float64 `json:"live_fans_club_join_cnt,omitempty"`
	//
	LiveOrderRefundAmount14d *float64 `json:"live_order_refund_amount_14d,omitempty"`
	//
	LiveOrderRefundAmount7d *float64 `json:"live_order_refund_amount_7d,omitempty"`
	//
	LiveOrderRefundCount14d *int64 `json:"live_order_refund_count_14d,omitempty"`
	//
	LiveOrderRefundCount7d *int64 `json:"live_order_refund_count_7d,omitempty"`
	//
	LiveOrderSettleAmount14d *float64 `json:"live_order_settle_amount_14d,omitempty"`
	//
	LiveOrderSettleAmount7d *float64 `json:"live_order_settle_amount_7d,omitempty"`
	//
	LiveOrderSettleCount14d *int64 `json:"live_order_settle_count_14d,omitempty"`
	//
	LiveOrderSettleCount7d *int64 `json:"live_order_settle_count_7d,omitempty"`
	//
	LiveOrderSettleCountRate14d *float64 `json:"live_order_settle_count_rate_14d,omitempty"`
	//
	LiveOrderSettleCountRate7d *float64 `json:"live_order_settle_count_rate_7d,omitempty"`
	//
	LiveWatchOneMinuteCount *float64 `json:"live_watch_one_minute_count,omitempty"`
	//
	LubanLiveClickProductCnt *float64 `json:"luban_live_click_product_cnt,omitempty"`
	//
	LubanLiveCommentCnt *float64 `json:"luban_live_comment_cnt,omitempty"`
	//
	LubanLiveEnterCnt *float64 `json:"luban_live_enter_cnt,omitempty"`
	//
	LubanLiveGiftAmount *float64 `json:"luban_live_gift_amount,omitempty"`
	//
	LubanLiveGiftCnt *float64 `json:"luban_live_gift_cnt,omitempty"`
	//
	LubanLiveShareCnt *float64 `json:"luban_live_share_cnt,omitempty"`
	//
	LubanLiveSlidecartClickCnt *float64 `json:"luban_live_slidecart_click_cnt,omitempty"`
	//
	MaterialArpu *float64 `json:"material_arpu,omitempty"`
	//
	PayOrderAmount *float64 `json:"pay_order_amount,omitempty"`
	//
	PayOrderCount *float64 `json:"pay_order_count,omitempty"`
	//
	PayOrderCouponAmount *float64 `json:"pay_order_coupon_amount,omitempty"`
	//
	Play25FeedBreak *float64 `json:"play_25_feed_break,omitempty"`
	//
	Play50FeedBreak *float64 `json:"play_50_feed_break,omitempty"`
	//
	Play75FeedBreak *float64 `json:"play_75_feed_break,omitempty"`
	//
	PlayDuration10sRate *float64 `json:"play_duration_10s_rate,omitempty"`
	//
	PlayDuration2sRate *float64 `json:"play_duration_2s_rate,omitempty"`
	//
	PlayDuration3sRate *float64 `json:"play_duration_3s_rate,omitempty"`
	//
	PlayDuration5sRate *float64 `json:"play_duration_5s_rate,omitempty"`
	//
	PlayOver *float64 `json:"play_over,omitempty"`
	//
	PlayOverRate *float64 `json:"play_over_rate,omitempty"`
	//
	PrepayAndPayOrderRoi *float64 `json:"prepay_and_pay_order_roi,omitempty"`
	//
	PrepayOrderAmount *float64 `json:"prepay_order_amount,omitempty"`
	//
	PrepayOrderCount *float64 `json:"prepay_order_count,omitempty"`
	//
	ReportCnt *float64 `json:"report_cnt,omitempty"`
	//
	ShowCnt *float64 `json:"show_cnt,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	TotalPlay *float64 `json:"total_play,omitempty"`
	//
	UnfinishedEstimateOrderGmv *float64 `json:"unfinished_estimate_order_gmv,omitempty"`
	//
	ValidPlay *float64 `json:"valid_play,omitempty"`
	//
	ValidPlayRate *float64 `json:"valid_play_rate,omitempty"`
}
