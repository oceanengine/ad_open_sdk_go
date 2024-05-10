/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCreativeGetV10ResponseDataListInner struct for QianchuanReportCreativeGetV10ResponseDataListInner
type QianchuanReportCreativeGetV10ResponseDataListInner struct {
	//
	AdAllOrderGmvSettleRate14d *float64 `json:"ad_all_order_gmv_settle_rate_14d,omitempty"`
	//
	AdAllOrderGmvSettleRate7d *float64 `json:"ad_all_order_gmv_settle_rate_7d,omitempty"`
	//
	AdId *int64 `json:"ad_id,omitempty"`
	//
	AdLiveOrderSettleCost14d *float64 `json:"ad_live_order_settle_cost_14d,omitempty"`
	//
	AdLiveOrderSettleCost7d *float64 `json:"ad_live_order_settle_cost_7d,omitempty"`
	//
	AdLiveOrderSettleRoi14d *float64 `json:"ad_live_order_settle_roi_14d,omitempty"`
	//
	AdLiveOrderSettleRoi7d *float64 `json:"ad_live_order_settle_roi_7d,omitempty"`
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AllOrderCreateRoi7days *float64 `json:"all_order_create_roi_7days,omitempty"`
	//
	AllOrderPayCount7days *float64 `json:"all_order_pay_count_7days,omitempty"`
	//
	AllOrderPayGmv7days *float64 `json:"all_order_pay_gmv_7days,omitempty"`
	//
	AllOrderPayRoi7days *float64 `json:"all_order_pay_roi_7days,omitempty"`
	//
	AttributionConvertCnt *float64 `json:"attribution_convert_cnt,omitempty"`
	//
	AttributionConvertCost *float64 `json:"attribution_convert_cost,omitempty"`
	//
	AttributionConvertRate *float64 `json:"attribution_convert_rate,omitempty"`
	//
	AttributionDeepConvertCnt *float64 `json:"attribution_deep_convert_cnt,omitempty"`
	//
	AttributionDeepConvertCost *float64 `json:"attribution_deep_convert_cost,omitempty"`
	//
	AttributionDeepConvertRate *float64 `json:"attribution_deep_convert_rate,omitempty"`
	//
	ClickCnt *float64 `json:"click_cnt,omitempty"`
	//
	ConvertCnt *float64 `json:"convert_cnt,omitempty"`
	//
	ConvertCost *float64 `json:"convert_cost,omitempty"`
	//
	ConvertRate *float64 `json:"convert_rate,omitempty"`
	//
	CpcPlatform *float64 `json:"cpc_platform,omitempty"`
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
	CreativeId *int64 `json:"creative_id,omitempty"`
	//
	Ctr *float64 `json:"ctr,omitempty"`
	//
	DeepConvertCnt *float64 `json:"deep_convert_cnt,omitempty"`
	//
	DeepConvertCost *float64 `json:"deep_convert_cost,omitempty"`
	//
	DeepConvertRate *float64 `json:"deep_convert_rate,omitempty"`
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
	IndirectOrderCreateCount7days *float64 `json:"indirect_order_create_count_7days,omitempty"`
	//
	IndirectOrderCreateGmv7days *float64 `json:"indirect_order_create_gmv_7days,omitempty"`
	//
	IndirectOrderPayCount7days *float64 `json:"indirect_order_pay_count_7days,omitempty"`
	//
	IndirectOrderPayGmv7days *float64 `json:"indirect_order_pay_gmv_7days,omitempty"`
	//
	IndirectOrderPrepayCount7days *float64 `json:"indirect_order_prepay_count_7days,omitempty"`
	//
	IndirectOrderPrepayGmv7days *float64 `json:"indirect_order_prepay_gmv_7days,omitempty"`
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
	LivePayOrderCostPerOrder *float64 `json:"live_pay_order_cost_per_order,omitempty"`
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
	PayOrderAmount *float64 `json:"pay_order_amount,omitempty"`
	//
	PayOrderCostPerOrder *float64 `json:"pay_order_cost_per_order,omitempty"`
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
	PlayDuration3s *float64 `json:"play_duration_3s,omitempty"`
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
	QianchuanAuthorFirstOrderCnt *float64 `json:"qianchuan_author_first_order_cnt,omitempty"`
	//
	QianchuanAuthorFirstOrderConvertCost *float64 `json:"qianchuan_author_first_order_convert_cost,omitempty"`
	//
	QianchuanAuthorFirstOrderDirectPayGmv *float64 `json:"qianchuan_author_first_order_direct_pay_gmv,omitempty"`
	//
	QianchuanAuthorFirstOrderDirectPayOrderRoi *float64 `json:"qianchuan_author_first_order_direct_pay_order_roi,omitempty"`
	//
	QianchuanAuthorFirstOrderLtv30 *float64 `json:"qianchuan_author_first_order_ltv30,omitempty"`
	//
	QianchuanAuthorFirstOrderRate *float64 `json:"qianchuan_author_first_order_rate,omitempty"`
	//
	QianchuanAuthorFirstOrderRoi30 *float64 `json:"qianchuan_author_first_order_roi30,omitempty"`
	//
	QianchuanBrandFirstOrderCnt *float64 `json:"qianchuan_brand_first_order_cnt,omitempty"`
	//
	QianchuanBrandFirstOrderConvertCost *float64 `json:"qianchuan_brand_first_order_convert_cost,omitempty"`
	//
	QianchuanBrandFirstOrderDirectPayGmv *float64 `json:"qianchuan_brand_first_order_direct_pay_gmv,omitempty"`
	//
	QianchuanBrandFirstOrderDirectPayOrderRoi *float64 `json:"qianchuan_brand_first_order_direct_pay_order_roi,omitempty"`
	//
	QianchuanBrandFirstOrderLtv30 *float64 `json:"qianchuan_brand_first_order_ltv30,omitempty"`
	//
	QianchuanBrandFirstOrderRate *float64 `json:"qianchuan_brand_first_order_rate,omitempty"`
	//
	QianchuanBrandFirstOrderRoi30 *float64 `json:"qianchuan_brand_first_order_roi30,omitempty"`
	//
	QianchuanFirstOrderCnt *float64 `json:"qianchuan_first_order_cnt,omitempty"`
	//
	QianchuanFirstOrderConvertCost *float64 `json:"qianchuan_first_order_convert_cost,omitempty"`
	//
	QianchuanFirstOrderDirectPayGmv *float64 `json:"qianchuan_first_order_direct_pay_gmv,omitempty"`
	//
	QianchuanFirstOrderDirectPayOrderRoi *float64 `json:"qianchuan_first_order_direct_pay_order_roi,omitempty"`
	//
	QianchuanFirstOrderLtv30 *float64 `json:"qianchuan_first_order_ltv30,omitempty"`
	//
	QianchuanFirstOrderRate *float64 `json:"qianchuan_first_order_rate,omitempty"`
	//
	QianchuanFirstOrderRoi30 *float64 `json:"qianchuan_first_order_roi30,omitempty"`
	//
	ShowCnt *float64 `json:"show_cnt,omitempty"`
	//
	StatCost *float64 `json:"stat_cost,omitempty"`
	//
	StatDatetime *string `json:"stat_datetime,omitempty"`
	//
	TotalPlay *float64 `json:"total_play,omitempty"`
	//
	UnfinishedEstimateOrderGmv *float64 `json:"unfinished_estimate_order_gmv,omitempty"`
}
