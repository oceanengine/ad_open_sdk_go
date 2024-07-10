/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportLiveGetV10ResponseData
type QianchuanReportLiveGetV10ResponseData struct {
	//
	AdAllOrderGmvSettleRate14d *float64 `json:"ad_all_order_gmv_settle_rate_14d,omitempty"`
	//
	AdAllOrderGmvSettleRate7d *float64 `json:"ad_all_order_gmv_settle_rate_7d,omitempty"`
	//
	AdLiveCreateOrderRate *float64 `json:"ad_live_create_order_rate,omitempty"`
	//
	AdLiveOrderSettleCost14d *float64 `json:"ad_live_order_settle_cost_14d,omitempty"`
	//
	AdLiveOrderSettleCost7d *float64 `json:"ad_live_order_settle_cost_7d,omitempty"`
	//
	AdLiveOrderSettleRoi14d *float64 `json:"ad_live_order_settle_roi_14d,omitempty"`
	//
	AdLiveOrderSettleRoi7d *float64 `json:"ad_live_order_settle_roi_7d,omitempty"`
	//
	AdLivePayOrderGmvAvg *float64 `json:"ad_live_pay_order_gmv_avg,omitempty"`
	//
	AdLivePayOrderRate *float64 `json:"ad_live_pay_order_rate,omitempty"`
	//
	AdLivePrepayAndPayOrderGmvRoi *float64 `json:"ad_live_prepay_and_pay_order_gmv_roi,omitempty"`
	//
	ClickCnt *int64 `json:"click_cnt,omitempty"`
	//
	ConvertCnt *int64 `json:"convert_cnt,omitempty"`
	//
	ConvertRate *float64 `json:"convert_rate,omitempty"`
	//
	CpaPlatform *float64 `json:"cpa_platform,omitempty"`
	//
	CpcPlatform *float64 `json:"cpc_platform,omitempty"`
	//
	CpmPlatform *float64 `json:"cpm_platform,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	EcpConvertCnt *float64 `json:"ecp_convert_cnt,omitempty"`
	//
	EcpConvertRate *float64 `json:"ecp_convert_rate,omitempty"`
	//
	EcpCpaPlatform *float64 `json:"ecp_cpa_platform,omitempty"`
	//
	LiveClickCartCountAlias *int64 `json:"live_click_cart_count_alias,omitempty"`
	//
	LiveClickProductCountAlias *int64 `json:"live_click_product_count_alias,omitempty"`
	//
	LiveCreateOrderCountAlias *int64 `json:"live_create_order_count_alias,omitempty"`
	//
	LiveCreateOrderRate *float64 `json:"live_create_order_rate,omitempty"`
	//
	LiveOrderPayCouponAmount *float64 `json:"live_order_pay_coupon_amount,omitempty"`
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
	LivePayOrderCountAlias *int64 `json:"live_pay_order_count_alias,omitempty"`
	//
	LivePayOrderGmvAlias *float64 `json:"live_pay_order_gmv_alias,omitempty"`
	//
	LivePayOrderGmvAvg *float64 `json:"live_pay_order_gmv_avg,omitempty"`
	//
	LivePayOrderGmvRoi *float64 `json:"live_pay_order_gmv_roi,omitempty"`
	//
	LivePayOrderRate *float64 `json:"live_pay_order_rate,omitempty"`
	//
	LivePrepayOrderCountAlias *float64 `json:"live_prepay_order_count_alias,omitempty"`
	//
	LivePrepayOrderGmvAlias *float64 `json:"live_prepay_order_gmv_alias,omitempty"`
	//
	LiveWatchOneMinuteCount *int64 `json:"live_watch_one_minute_count,omitempty"`
	//
	LubanLiveOrderCount *int64 `json:"luban_live_order_count,omitempty"`
	//
	LubanLivePayOrderCount *int64 `json:"luban_live_pay_order_count,omitempty"`
	//
	LubanLivePayOrderGmv *float64 `json:"luban_live_pay_order_gmv,omitempty"`
	//
	LubanLivePayOrderGpm *float64 `json:"luban_live_pay_order_gpm,omitempty"`
	//
	LubanLivePrepayOrderCount *int64 `json:"luban_live_prepay_order_count,omitempty"`
	//
	LubanLivePrepayOrderGmv *float64 `json:"luban_live_prepay_order_gmv,omitempty"`
	//
	QualificationStatCost *float64 `json:"qualification_stat_cost,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	TotalLiveCommentCnt *int64 `json:"total_live_comment_cnt,omitempty"`
	//
	TotalLiveFansClubJoinCnt *int64 `json:"total_live_fans_club_join_cnt,omitempty"`
	//
	TotalLiveFollowCnt *int64 `json:"total_live_follow_cnt,omitempty"`
	//
	TotalLiveGiftAmount *float64 `json:"total_live_gift_amount,omitempty"`
	//
	TotalLiveGiftCnt *int64 `json:"total_live_gift_cnt,omitempty"`
	//
	TotalLivePayOrderGpm *float64 `json:"total_live_pay_order_gpm,omitempty"`
	//
	TotalLiveShareCnt *int64 `json:"total_live_share_cnt,omitempty"`
	//
	TotalLiveWatchCnt *int64 `json:"total_live_watch_cnt,omitempty"`
}
