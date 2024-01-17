/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2OrderField
type ReportAdvertiserGetV2OrderField string

// List of report_advertiser_get_v2_order_field
const (
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_2days"
	CPA_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpa"
	COST_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "cost"
	ACTIVE_COST_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "active_cost"
	STAT_UNION_LTV_0_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_0"
	WECHAT_FIRST_PAY_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "wechat_first_pay_rate"
	DEEP_CONVERT_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "deep_convert"
	MAP_SEARCH_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "map_search"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_7d_rate"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "luban_live_slidecart_click_cnt"
	CLICK_SHOPWINDOW_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_CONVERT_COST_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_6days"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdvertiserGetV2OrderField                 ReportAdvertiserGetV2OrderField = "live_watch_one_minute_count"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdvertiserGetV2OrderField             ReportAdvertiserGetV2OrderField = "advanced_creative_counsel_click"
	AVG_RANK_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "avg_rank"
	SHOPPING_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "shopping"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_2d_cost"
	CPC_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpc"
	PLAY_25_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_25_feed_break_rate"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdvertiserGetV2OrderField          ReportAdvertiserGetV2OrderField = "attribution_wechat_login_30d_count"
	WECHAT_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "wechat_pay_amount"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_6days"
	FOLLOW_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "follow"
	UNION_ROI_0_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_0"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_day_acitve_pay_count"
	CLICK_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "click"
	PAY_COUNT_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "pay_count"
	DOWNLOAD_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "download"
	LUBAN_LIVE_ENTER_CNT_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "luban_live_enter_cnt"
	PLAY_75_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_75_feed_break"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "advanced_creative_form_submit"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_0d_ltv"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_8days"
	PLAY_DURATION_2S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_2s_rate"
	LUBAN_LIVE_GIFT_CNT_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "luban_live_gift_cnt"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_3d_cost"
	PHONE_CONNECT_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "phone_connect"
	PHONE_EFFECTIVE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "phone_effective"
	PRE_CONVERT_COST_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "pre_convert_cost"
	WECHAT_FIRST_PAY_COST_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "wechat_first_pay_cost"
	IN_APP_PAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "in_app_pay"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_wechat_pay_30d_amount"
	APPROVAL_COUNT_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "approval_count"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_7d_ltv"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "luban_live_pay_order_count"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_2days"
	INSTALL_FINISH_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "install_finish"
	GAME_PAY_COUNT_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "game_pay_count"
	UNION_ROI_7_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_7"
	FIRST_RENTAL_ORDER_COUNT_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "first_rental_order_count"
	FIRST_ORDER_COUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "first_order_count"
	DOWNLOAD_FINISH_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_0d_roi"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "advanced_creative_phone_click"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "advanced_creative_coupon_addition"
	DOWNLOAD_FINISH_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "download_finish"
	REPORT_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "report"
	LUBAN_ORDER_CNT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "luban_order_cnt"
	CONSULT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "consult"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_3d_ltv"
	CUSTOMER_EFFECTIVE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "customer_effective"
	VALID_PLAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "valid_play"
	COMMENT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "comment"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdvertiserGetV2OrderField   ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "luban_live_pay_order_stat_cost"
	PHONE_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "phone"
	LOAN_CREDIT_COST_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "loan_credit_cost"
	PLAY_50_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_50_feed_break"
	SHOW_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "show"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdvertiserGetV2OrderField         ReportAdvertiserGetV2OrderField = "attribution_retention_7d_total_cost"
	FORM_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "form"
	VALID_PLAY_COST_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "valid_play_cost"
	IN_APP_UV_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "in_app_uv"
	GAME_ADDICTION_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "game_addiction_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_5days"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_5d_rate"
	AVERAGE_PLAY_PROGRESS_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "average_play_progress"
	CTR_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "ctr"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_deep_convert_cost"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_7d_cost"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "luban_live_follow_cnt"
	WECHAT_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "wechat"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdvertiserGetV2OrderField ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	LOAN_CREDIT_RATE_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "loan_credit_rate"
	IN_APP_ORDER_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "in_app_order"
	LOCATION_CLICK_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "location_click"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_8days"
	DOWNLOAD_FINISH_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "download_finish_rate"
	WIFI_PLAY_RATE_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "wifi_play_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_1day"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_4days"
	POI_COLLECT_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "poi_collect"
	DOWNLOAD_START_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "download_start"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_6d_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_rate"
	ACTIVE_PAY_COST_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "active_pay_cost"
	REDIRECT_TO_SHOP_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "redirect_to_shop"
	AVG_CLICK_COST_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "avg_click_cost"
	DOWNLOAD_START_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "download_start_cost"
	PLAY_DURATION_10S_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "play_duration_10s"
	GAME_PAY_COST_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_CONVERT_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "attribution_convert"
	PLAY_25_FEED_BREAK_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "play_25_feed_break"
	PLAY_75_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_75_feed_break_rate"
	LUBAN_ORDER_ROI_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "luban_order_roi"
	GAME_ADDICTION_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "game_addiction"
	PLAY_DURATION_SUM_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "play_duration_sum"
	PLAY_DURATION_5S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_5s_rate"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_7d_cnt"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_game_pay_7d_count"
	PLAY_OVER_RATE_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "play_over_rate"
	ACTIVE_REGISTER_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "active_register_cost"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_5d_cnt"
	CLICK_WEBSITE_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_website"
	MESSAGE_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "message"
	QQ_ReportAdvertiserGetV2OrderField                                          ReportAdvertiserGetV2OrderField = "qq"
	PLAY_50_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "play_50_feed_break_rate"
	HOME_VISITED_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "home_visited"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_7d_roi"
	CONVERT_COST_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "convert_cost"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "live_fans_club_join_cnt"
	ACTIVE_PAY_RATE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "active_pay_rate"
	CLICK_LANDING_PAGE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "click_landing_page"
	NEXT_DAY_OPEN_RATE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "next_day_open_rate"
	VOTE_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "vote"
	ACTIVE_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "active"
	PRE_CONVERT_COUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "pre_convert_count"
	REDIRECT_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "redirect"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdvertiserGetV2OrderField      ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	PLAY_OVER_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "play_over"
	LIKE_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "like"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_4d_rate"
	INSTALL_FINISH_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "install_finish_cost"
	WECHAT_FIRST_PAY_COUNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "wechat_first_pay_count"
	CONSULT_EFFECTIVE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "consult_effective"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "advanced_creative_form_click"
	IES_CHALLENGE_CLICK_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "ies_challenge_click"
	CONVERT_SHOW_RATE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_4d_cost"
	LUBAN_LIVE_SHARE_CNT_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "luban_live_share_cnt"
	IN_APP_DETAIL_UV_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "in_app_detail_uv"
	PHONE_CONFIRM_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "phone_confirm"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_3d_rate"
	COUPON_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "coupon"
	SHARE_ReportAdvertiserGetV2OrderField                                       ReportAdvertiserGetV2OrderField = "share"
	CLICK_DOWNLOAD_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "click_download"
	VALID_PLAY_RATE_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "valid_play_rate"
	PAY_AMOUNT_ROI_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "pay_amount_roi"
	WECHAT_LOGIN_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "wechat_login_cost"
	ACTIVE_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "active_pay_amount"
	LUBAN_LIVE_COMMENT_CNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "luban_live_comment_cnt"
	CLICK_CALL_DY_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_call_dy"
	COMMUTE_FIRST_PAY_COUNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "commute_first_pay_count"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdvertiserGetV2OrderField                     ReportAdvertiserGetV2OrderField = "luban_order_stat_amount"
	VIEW_ReportAdvertiserGetV2OrderField                                        ReportAdvertiserGetV2OrderField = "view"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_game_pay_7d_cost"
	PRE_CONVERT_RATE_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_3d_cnt"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_next_day_open_cost"
	LOTTERY_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "lottery"
	CLICK_INSTALL_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "click_install"
	PLAY_DURATION_2S_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "play_duration_2s"
	CARD_SHOW_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "card_show"
	PLAY_100_FEED_BREAK_RATE_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "play_100_feed_break_rate"
	STAT_PAY_AMOUNT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "stat_pay_amount"
	DOWNLOAD_START_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "download_start_rate"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_5d_cost"
	PRE_LOAN_CREDIT_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "pre_loan_credit_cost"
	SUBMIT_CERTIFICATION_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "submit_certification_count"
	PLAY_DURATION_3S_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_micro_game_3d_roi"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "luban_live_click_product_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdvertiserGetV2OrderField   ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdvertiserGetV2OrderField    ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	COUPON_SINGLE_PAGE_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "coupon_single_page"
	IES_MUSIC_CLICK_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "ies_music_click"
	INTERACT_PER_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_wechat_login_30d_cost"
	LIVE_COMPONENT_CLICK_COST_ReportAdvertiserGetV2OrderField                   ReportAdvertiserGetV2OrderField = "live_component_click_cost"
	PLAY_DURATION_3S_RATE_ReportAdvertiserGetV2OrderField                       ReportAdvertiserGetV2OrderField = "play_duration_3s_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdvertiserGetV2OrderField       ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_cost"
	PLAY_DURATION_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "play_duration"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_7days"
	IN_APP_CART_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "in_app_cart"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "average_play_time_per_play"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdvertiserGetV2OrderField                  ReportAdvertiserGetV2OrderField = "live_component_click_count"
	AVERAGE_VIDEO_PLAY_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "average_video_play"
	CPM_ReportAdvertiserGetV2OrderField                                         ReportAdvertiserGetV2OrderField = "cpm"
	ACTIVE_RATE_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "active_rate"
	PRE_LOAN_CREDIT_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "pre_loan_credit"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdvertiserGetV2OrderField       ReportAdvertiserGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	WIFI_PLAY_ReportAdvertiserGetV2OrderField                                   ReportAdvertiserGetV2OrderField = "wifi_play"
	LOAN_CREDIT_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "loan_credit"
	LIVE_COMPONENT_CLICK_RATE_ReportAdvertiserGetV2OrderField                   ReportAdvertiserGetV2OrderField = "live_component_click_rate"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "luban_live_gift_amount"
	CONVERT_RATE_ReportAdvertiserGetV2OrderField                                ReportAdvertiserGetV2OrderField = "convert_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_game_in_app_ltv_1day"
	REGISTER_ReportAdvertiserGetV2OrderField                                    ReportAdvertiserGetV2OrderField = "register"
	BUTTON_ReportAdvertiserGetV2OrderField                                      ReportAdvertiserGetV2OrderField = "button"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdvertiserGetV2OrderField              ReportAdvertiserGetV2OrderField = "attribution_wechat_pay_30d_roi"
	ACTIVE_REGISTER_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "active_register_rate"
	MESSAGE_ACTION_ReportAdvertiserGetV2OrderField                              ReportAdvertiserGetV2OrderField = "message_action"
	LOAN_COMPLETION_ReportAdvertiserGetV2OrderField                             ReportAdvertiserGetV2OrderField = "loan_completion"
	CONVERT_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "convert"
	NEXT_DAY_OPEN_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "next_day_open"
	DISLIKE_ReportAdvertiserGetV2OrderField                                     ReportAdvertiserGetV2OrderField = "dislike"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_retention_2d_rate"
	LOAN_COMPLETION_COST_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "loan_completion_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdvertiserGetV2OrderField         ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_per_count"
	POI_ADDRESS_CLICK_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "poi_address_click"
	DEEP_CONVERT_RATE_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "deep_convert_rate"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_6d_cnt"
	TOTAL_PLAY_ReportAdvertiserGetV2OrderField                                  ReportAdvertiserGetV2OrderField = "total_play"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdvertiserGetV2OrderField               ReportAdvertiserGetV2OrderField = "attribution_next_day_open_cnt"
	UNION_ROI_3_ReportAdvertiserGetV2OrderField                                 ReportAdvertiserGetV2OrderField = "union_roi_3"
	DEEP_CONVERT_COST_ReportAdvertiserGetV2OrderField                           ReportAdvertiserGetV2OrderField = "deep_convert_cost"
	AVG_SHOW_COST_ReportAdvertiserGetV2OrderField                               ReportAdvertiserGetV2OrderField = "avg_show_cost"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdvertiserGetV2OrderField                ReportAdvertiserGetV2OrderField = "attribution_retention_2d_cnt"
	PLAY_DURATION_10S_RATE_ReportAdvertiserGetV2OrderField                      ReportAdvertiserGetV2OrderField = "play_duration_10s_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdvertiserGetV2OrderField           ReportAdvertiserGetV2OrderField = "attribution_game_in_app_roi_3days"
	ATTRIBUTION_DEEP_CONVERT_ReportAdvertiserGetV2OrderField                    ReportAdvertiserGetV2OrderField = "attribution_deep_convert"
	STAT_UNION_LTV_3_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_3"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdvertiserGetV2OrderField            ReportAdvertiserGetV2OrderField = "attribution_day_active_pay_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdvertiserGetV2OrderField  ReportAdvertiserGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdvertiserGetV2OrderField             ReportAdvertiserGetV2OrderField = "attribution_active_pay_7d_count"
	INSTALL_FINISH_RATE_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "install_finish_rate"
	GAME_ADDICTION_COST_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "game_addiction_cost"
	NEXT_DAY_OPEN_COST_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "next_day_open_cost"
	WECHAT_LOGIN_COUNT_ReportAdvertiserGetV2OrderField                          ReportAdvertiserGetV2OrderField = "wechat_login_count"
	PLAY_100_FEED_BREAK_ReportAdvertiserGetV2OrderField                         ReportAdvertiserGetV2OrderField = "play_100_feed_break"
	STAT_UNION_LTV_7_ReportAdvertiserGetV2OrderField                            ReportAdvertiserGetV2OrderField = "stat_union_ltv_7"
	LOAN_COMPLETION_RATE_ReportAdvertiserGetV2OrderField                        ReportAdvertiserGetV2OrderField = "loan_completion_rate"
)

// All allowed values of ReportAdvertiserGetV2OrderField enum
var AllowedReportAdvertiserGetV2OrderFieldEnumValues = []ReportAdvertiserGetV2OrderField{
	"attribution_game_in_app_ltv_3days",
	"attribution_game_in_app_ltv_2days",
	"cpa",
	"cost",
	"active_cost",
	"stat_union_ltv_0",
	"wechat_first_pay_rate",
	"deep_convert",
	"map_search",
	"attribution_retention_7d_rate",
	"luban_live_slidecart_click_cnt",
	"click_shopwindow",
	"attribution_convert_cost",
	"attribution_next_day_open_rate",
	"attribution_game_in_app_roi_6days",
	"live_watch_one_minute_count",
	"advanced_creative_counsel_click",
	"avg_rank",
	"shopping",
	"attribution_retention_2d_cost",
	"cpc",
	"play_25_feed_break_rate",
	"attribution_wechat_login_30d_count",
	"wechat_pay_amount",
	"attribution_game_in_app_ltv_6days",
	"follow",
	"union_roi_0",
	"attribution_day_acitve_pay_count",
	"click",
	"pay_count",
	"download",
	"luban_live_enter_cnt",
	"play_75_feed_break",
	"advanced_creative_form_submit",
	"attribution_micro_game_0d_ltv",
	"attribution_game_in_app_roi_8days",
	"play_duration_2s_rate",
	"luban_live_gift_cnt",
	"attribution_retention_3d_cost",
	"phone_connect",
	"phone_effective",
	"pre_convert_cost",
	"wechat_first_pay_cost",
	"in_app_pay",
	"attribution_wechat_pay_30d_amount",
	"approval_count",
	"attribution_micro_game_7d_ltv",
	"luban_live_pay_order_count",
	"attribution_game_in_app_roi_2days",
	"install_finish",
	"game_pay_count",
	"union_roi_7",
	"first_rental_order_count",
	"first_order_count",
	"download_finish_cost",
	"attribution_micro_game_0d_roi",
	"advanced_creative_phone_click",
	"advanced_creative_coupon_addition",
	"download_finish",
	"report",
	"luban_order_cnt",
	"consult",
	"attribution_micro_game_3d_ltv",
	"customer_effective",
	"valid_play",
	"comment",
	"attribution_game_in_app_ltv_4days",
	"attribution_active_pay_intra_one_day_cost",
	"luban_live_pay_order_stat_cost",
	"phone",
	"loan_credit_cost",
	"play_50_feed_break",
	"show",
	"attribution_retention_7d_total_cost",
	"form",
	"valid_play_cost",
	"in_app_uv",
	"game_addiction_rate",
	"attribution_game_in_app_roi_5days",
	"attribution_retention_4d_cnt",
	"attribution_retention_5d_rate",
	"average_play_progress",
	"ctr",
	"attribution_game_in_app_ltv_5days",
	"attribution_deep_convert_cost",
	"attribution_retention_7d_cost",
	"luban_live_follow_cnt",
	"wechat",
	"attribution_active_pay_intra_one_day_amount",
	"loan_credit_rate",
	"in_app_order",
	"location_click",
	"attribution_game_in_app_ltv_8days",
	"download_finish_rate",
	"wifi_play_rate",
	"attribution_game_in_app_roi_1day",
	"attribution_game_in_app_roi_4days",
	"poi_collect",
	"download_start",
	"attribution_retention_6d_cost",
	"attribution_active_pay_7d_rate",
	"active_pay_cost",
	"redirect_to_shop",
	"avg_click_cost",
	"download_start_cost",
	"play_duration_10s",
	"game_pay_cost",
	"attribution_convert",
	"play_25_feed_break",
	"play_75_feed_break_rate",
	"luban_order_roi",
	"game_addiction",
	"play_duration_sum",
	"play_duration_5s_rate",
	"attribution_retention_7d_cnt",
	"attribution_game_pay_7d_count",
	"play_over_rate",
	"active_register_cost",
	"attribution_retention_5d_cnt",
	"click_website",
	"message",
	"qq",
	"play_50_feed_break_rate",
	"home_visited",
	"attribution_micro_game_7d_roi",
	"convert_cost",
	"live_fans_club_join_cnt",
	"active_pay_rate",
	"click_landing_page",
	"next_day_open_rate",
	"vote",
	"active",
	"pre_convert_count",
	"redirect",
	"attribution_wechat_first_pay_30d_count",
	"play_over",
	"like",
	"attribution_retention_4d_rate",
	"install_finish_cost",
	"wechat_first_pay_count",
	"consult_effective",
	"advanced_creative_form_click",
	"ies_challenge_click",
	"convert_show_rate",
	"attribution_retention_4d_cost",
	"luban_live_share_cnt",
	"in_app_detail_uv",
	"phone_confirm",
	"attribution_retention_3d_rate",
	"coupon",
	"share",
	"click_download",
	"valid_play_rate",
	"pay_amount_roi",
	"wechat_login_cost",
	"active_pay_amount",
	"luban_live_comment_cnt",
	"click_call_dy",
	"commute_first_pay_count",
	"luban_order_stat_amount",
	"view",
	"attribution_game_pay_7d_cost",
	"pre_convert_rate",
	"attribution_retention_3d_cnt",
	"attribution_retention_7d_sum_cnt",
	"attribution_next_day_open_cost",
	"lottery",
	"click_install",
	"play_duration_2s",
	"card_show",
	"play_100_feed_break_rate",
	"stat_pay_amount",
	"download_start_rate",
	"attribution_retention_5d_cost",
	"pre_loan_credit_cost",
	"submit_certification_count",
	"play_duration_3s",
	"attribution_micro_game_3d_roi",
	"luban_live_click_product_cnt",
	"attribution_active_pay_intra_one_day_rate",
	"attribution_active_pay_intra_one_day_roi",
	"coupon_single_page",
	"ies_music_click",
	"interact_per_cost",
	"attribution_wechat_login_30d_cost",
	"live_component_click_cost",
	"play_duration_3s_rate",
	"attribution_wechat_first_pay_30d_cost",
	"attribution_active_pay_7d_cost",
	"play_duration",
	"attribution_game_in_app_ltv_7days",
	"attribution_retention_6d_rate",
	"attribution_game_in_app_roi_7days",
	"in_app_cart",
	"average_play_time_per_play",
	"live_component_click_count",
	"average_video_play",
	"cpm",
	"active_rate",
	"pre_loan_credit",
	"attribution_wechat_first_pay_30d_rate",
	"wifi_play",
	"loan_credit",
	"live_component_click_rate",
	"luban_live_gift_amount",
	"convert_rate",
	"attribution_game_in_app_ltv_1day",
	"register",
	"button",
	"attribution_wechat_pay_30d_roi",
	"active_register_rate",
	"message_action",
	"loan_completion",
	"convert",
	"next_day_open",
	"dislike",
	"attribution_retention_2d_rate",
	"loan_completion_cost",
	"attribution_active_pay_7d_per_count",
	"poi_address_click",
	"deep_convert_rate",
	"attribution_retention_6d_cnt",
	"total_play",
	"attribution_next_day_open_cnt",
	"union_roi_3",
	"deep_convert_cost",
	"avg_show_cost",
	"attribution_retention_2d_cnt",
	"play_duration_10s_rate",
	"attribution_game_in_app_roi_3days",
	"attribution_deep_convert",
	"stat_union_ltv_3",
	"attribution_day_active_pay_count",
	"attribution_active_pay_intra_one_day_count",
	"attribution_active_pay_7d_count",
	"install_finish_rate",
	"game_addiction_cost",
	"next_day_open_cost",
	"wechat_login_count",
	"play_100_feed_break",
	"stat_union_ltv_7",
	"loan_completion_rate",
}

// NewReportAdvertiserGetV2OrderFieldFromValue returns a pointer to a valid ReportAdvertiserGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2OrderFieldFromValue(v string) (*ReportAdvertiserGetV2OrderField, error) {
	ev := ReportAdvertiserGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2OrderField: valid values are %v", v, AllowedReportAdvertiserGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_order_field value
func (v ReportAdvertiserGetV2OrderField) Ptr() *ReportAdvertiserGetV2OrderField {
	return &v
}
