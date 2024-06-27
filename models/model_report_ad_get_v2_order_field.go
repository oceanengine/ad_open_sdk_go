/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2OrderField
type ReportAdGetV2OrderField string

// List of report_ad_get_v2_order_field
const (
	ATTRIBUTION_RETENTION_3D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_cost"
	GAME_PAY_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "game_pay_cost"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_phone_click"
	WIFI_PLAY_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "wifi_play_rate"
	LUBAN_LIVE_GIFT_CNT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "luban_live_gift_cnt"
	APPROVAL_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "approval_count"
	LOAN_COMPLETION_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_cost"
	NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_cost"
	DOWNLOAD_FINISH_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_cost"
	IN_APP_ORDER_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "in_app_order"
	PHONE_CONNECT_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_connect"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_active_pay_count"
	CUSTOMER_EFFECTIVE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "customer_effective"
	PLAY_100_FEED_BREAK_RATE_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	CLICK_INSTALL_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_install"
	PAY_AMOUNT_ROI_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "pay_amount_roi"
	CONVERT_RATE_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_rate"
	VIEW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "view"
	QQ_ReportAdGetV2OrderField                                          ReportAdGetV2OrderField = "qq"
	REPORT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "report"
	STAT_UNION_LTV_0_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_0"
	MESSAGE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "message"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_roi"
	WECHAT_LOGIN_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_login_cost"
	PHONE_EFFECTIVE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "phone_effective"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_slidecart_click_cnt"
	PLAY_DURATION_2S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_2s"
	DOWNLOAD_START_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_rate"
	LOTTERY_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "lottery"
	ACTIVE_COST_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_cost"
	IN_APP_DETAIL_UV_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "in_app_detail_uv"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_active_pay_7d_per_count"
	VOTE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "vote"
	PLAY_50_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_50_feed_break"
	AVG_RANK_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "avg_rank"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_cost"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_rate"
	LOAN_CREDIT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "loan_credit"
	PRE_LOAN_CREDIT_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "pre_loan_credit_cost"
	LUBAN_LIVE_ENTER_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_enter_cnt"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "live_component_click_count"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "luban_live_click_product_cnt"
	PLAY_25_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_25_feed_break"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_6days"
	GAME_ADDICTION_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_rate"
	LUBAN_LIVE_COMMENT_CNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_comment_cnt"
	CONSULT_EFFECTIVE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "consult_effective"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	MAP_SEARCH_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "map_search"
	PLAY_DURATION_10S_RATE_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "play_duration_10s_rate"
	DOWNLOAD_FINISH_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "download_finish"
	MESSAGE_ACTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "message_action"
	GAME_PAY_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_pay_count"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdGetV2OrderField          ReportAdGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_7days"
	TOTAL_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "total_play"
	SHARE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "share"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_ltv"
	DEEP_CONVERT_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_cost"
	LOAN_CREDIT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_rate"
	LUBAN_ORDER_CNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdGetV2OrderField ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	INSTALL_FINISH_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_rate"
	VALID_PLAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_rate"
	WECHAT_FIRST_PAY_COST_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_form_submit"
	FOLLOW_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "follow"
	LUBAN_LIVE_SHARE_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_share_cnt"
	DEEP_CONVERT_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_rate"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_cost"
	HOME_VISITED_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "home_visited"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_retention_7d_total_cost"
	PLAY_DURATION_SUM_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_sum"
	CLICK_LANDING_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "click_landing_page"
	LOCATION_CLICK_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "location_click"
	PLAY_25_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_25_feed_break_rate"
	UNION_ROI_7_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_7"
	BUTTON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "button"
	PHONE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "phone"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_roi"
	COMMENT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "comment"
	WECHAT_FIRST_PAY_COUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "wechat_first_pay_count"
	REDIRECT_TO_SHOP_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "redirect_to_shop"
	DOWNLOAD_START_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_cost"
	LOAN_COMPLETION_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_rate"
	ATTRIBUTION_CONVERT_COST_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_convert_cost"
	IN_APP_PAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "in_app_pay"
	CLICK_DOWNLOAD_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "click_download"
	DOWNLOAD_START_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "download_start"
	CONSULT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "consult"
	PLAY_100_FEED_BREAK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "play_100_feed_break"
	PLAY_OVER_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "play_over"
	ACTIVE_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "active"
	LOAN_CREDIT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_cost"
	POI_ADDRESS_CLICK_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "poi_address_click"
	DEEP_CONVERT_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "deep_convert"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_deep_convert_cost"
	STAT_PAY_AMOUNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "stat_pay_amount"
	ACTIVE_PAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_rate"
	PLAY_DURATION_2S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_2s_rate"
	LUBAN_ORDER_ROI_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_roi"
	AVG_CLICK_COST_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "avg_click_cost"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_5d_cnt"
	STAT_UNION_LTV_7_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_7"
	LOAN_COMPLETION_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "loan_completion"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_6d_cnt"
	COUPON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "coupon"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_7d_cnt"
	AVERAGE_VIDEO_PLAY_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "average_video_play"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdGetV2OrderField  ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	CPA_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpa"
	PLAY_DURATION_5S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_5s_rate"
	CLICK_SHOPWINDOW_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_2days"
	ACTIVE_REGISTER_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_cost"
	CPM_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpm"
	PRE_CONVERT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_8days"
	COST_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "cost"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "advanced_creative_counsel_click"
	IES_MUSIC_CLICK_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "ies_music_click"
	PLAY_DURATION_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "play_duration"
	INTERACT_PER_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "interact_per_cost"
	SUBMIT_CERTIFICATION_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "submit_certification_count"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_3days"
	DISLIKE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "dislike"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_cost"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "live_fans_club_join_cnt"
	UNION_ROI_0_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_0"
	UNION_ROI_3_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_3"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_4days"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_2days"
	INSTALL_FINISH_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_3d_cnt"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "advanced_creative_form_click"
	CLICK_CALL_DY_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_call_dy"
	VALID_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "valid_play"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_cost"
	POI_COLLECT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "poi_collect"
	NEXT_DAY_OPEN_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "next_day_open"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_rate"
	ACTIVE_PAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_cost"
	PHONE_CONFIRM_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_confirm"
	WECHAT_FIRST_PAY_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_rate"
	NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_rate"
	CARD_SHOW_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "card_show"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_acitve_pay_count"
	INSTALL_FINISH_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "install_finish"
	PLAY_DURATION_3S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_3s_rate"
	PLAY_DURATION_10S_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_10s"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_8days"
	FIRST_RENTAL_ORDER_COUNT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "first_rental_order_count"
	PRE_CONVERT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_cost"
	ACTIVE_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "active_pay_amount"
	PRE_CONVERT_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "pre_convert_count"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_login_30d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdGetV2OrderField      ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "attribution_active_pay_7d_count"
	WECHAT_LOGIN_COUNT_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "wechat_login_count"
	COUPON_SINGLE_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "coupon_single_page"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_cost"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "luban_live_follow_cnt"
	ACTIVE_RATE_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_rate"
	COMMUTE_FIRST_PAY_COUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "commute_first_pay_count"
	PLAY_DURATION_3S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_3s"
	PLAY_75_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "luban_live_pay_order_count"
	CLICK_WEBSITE_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_website"
	AVG_SHOW_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "avg_show_cost"
	CTR_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "ctr"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_2d_cnt"
	CONVERT_SHOW_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "convert_show_rate"
	WIFI_PLAY_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "wifi_play"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "luban_order_stat_amount"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_ltv_1day"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdGetV2OrderField                 ReportAdGetV2OrderField = "live_watch_one_minute_count"
	LIKE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "like"
	LIVE_COMPONENT_CLICK_COST_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdGetV2OrderField    ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_wechat_pay_30d_roi"
	PRE_LOAN_CREDIT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "pre_loan_credit"
	REDIRECT_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "redirect"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_rate"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "average_play_time_per_play"
	LIVE_COMPONENT_CLICK_RATE_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_rate"
	FORM_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "form"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_rate"
	IN_APP_UV_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "in_app_uv"
	CONVERT_COST_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_cost"
	DOWNLOAD_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "download"
	PLAY_75_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_75_feed_break"
	STAT_UNION_LTV_3_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_3"
	SHOPPING_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "shopping"
	PLAY_50_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_50_feed_break_rate"
	GAME_ADDICTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_addiction"
	AVERAGE_PLAY_PROGRESS_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "average_play_progress"
	GAME_ADDICTION_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_CONVERT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "attribution_convert"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_6days"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_ltv"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "advanced_creative_coupon_addition"
	CLICK_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "click"
	CPC_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpc"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_roi"
	VALID_PLAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_cost"
	WECHAT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "wechat"
	IES_CHALLENGE_CLICK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "ies_challenge_click"
	CONVERT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "convert"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_7days"
	WECHAT_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_pay_amount"
	FIRST_ORDER_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "first_order_count"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_5days"
	IN_APP_CART_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "in_app_cart"
	SHOW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "show"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ACTIVE_REGISTER_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_roi_1day"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_pay_order_stat_cost"
	ATTRIBUTION_DEEP_CONVERT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_deep_convert"
	REGISTER_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "register"
	PAY_COUNT_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "pay_count"
	DOWNLOAD_FINISH_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_rate"
	PLAY_OVER_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "play_over_rate"
)

// All allowed values of ReportAdGetV2OrderField enum
var AllowedReportAdGetV2OrderFieldEnumValues = []ReportAdGetV2OrderField{
	"attribution_retention_3d_cost",
	"attribution_retention_5d_cost",
	"game_pay_cost",
	"advanced_creative_phone_click",
	"wifi_play_rate",
	"luban_live_gift_cnt",
	"approval_count",
	"loan_completion_cost",
	"next_day_open_cost",
	"download_finish_cost",
	"in_app_order",
	"phone_connect",
	"attribution_retention_4d_rate",
	"attribution_game_in_app_roi_4days",
	"attribution_day_active_pay_count",
	"customer_effective",
	"play_100_feed_break_rate",
	"attribution_active_pay_intra_one_day_rate",
	"click_install",
	"pay_amount_roi",
	"convert_rate",
	"view",
	"qq",
	"report",
	"stat_union_ltv_0",
	"message",
	"attribution_retention_4d_cnt",
	"attribution_active_pay_7d_cost",
	"attribution_wechat_first_pay_30d_cost",
	"attribution_micro_game_0d_roi",
	"wechat_login_cost",
	"phone_effective",
	"luban_live_slidecart_click_cnt",
	"play_duration_2s",
	"download_start_rate",
	"lottery",
	"active_cost",
	"in_app_detail_uv",
	"attribution_game_pay_7d_count",
	"attribution_active_pay_7d_per_count",
	"vote",
	"play_50_feed_break",
	"avg_rank",
	"attribution_retention_6d_rate",
	"attribution_retention_4d_cost",
	"attribution_retention_3d_rate",
	"loan_credit",
	"pre_loan_credit_cost",
	"luban_live_enter_cnt",
	"live_component_click_count",
	"luban_live_click_product_cnt",
	"play_25_feed_break",
	"attribution_game_in_app_ltv_6days",
	"game_addiction_rate",
	"luban_live_comment_cnt",
	"consult_effective",
	"attribution_active_pay_intra_one_day_cost",
	"map_search",
	"play_duration_10s_rate",
	"download_finish",
	"message_action",
	"game_pay_count",
	"attribution_wechat_login_30d_count",
	"attribution_retention_7d_rate",
	"attribution_game_in_app_ltv_3days",
	"attribution_game_in_app_roi_7days",
	"total_play",
	"share",
	"attribution_micro_game_0d_ltv",
	"deep_convert_cost",
	"loan_credit_rate",
	"luban_order_cnt",
	"attribution_active_pay_intra_one_day_amount",
	"install_finish_rate",
	"valid_play_rate",
	"wechat_first_pay_cost",
	"advanced_creative_form_submit",
	"follow",
	"luban_live_share_cnt",
	"deep_convert_rate",
	"attribution_retention_7d_cost",
	"home_visited",
	"attribution_retention_7d_total_cost",
	"play_duration_sum",
	"click_landing_page",
	"location_click",
	"play_25_feed_break_rate",
	"union_roi_7",
	"button",
	"phone",
	"attribution_micro_game_7d_roi",
	"comment",
	"wechat_first_pay_count",
	"redirect_to_shop",
	"download_start_cost",
	"loan_completion_rate",
	"attribution_convert_cost",
	"in_app_pay",
	"click_download",
	"download_start",
	"consult",
	"play_100_feed_break",
	"play_over",
	"active",
	"loan_credit_cost",
	"poi_address_click",
	"deep_convert",
	"attribution_deep_convert_cost",
	"stat_pay_amount",
	"active_pay_rate",
	"play_duration_2s_rate",
	"luban_order_roi",
	"avg_click_cost",
	"attribution_retention_5d_cnt",
	"stat_union_ltv_7",
	"loan_completion",
	"attribution_retention_2d_rate",
	"attribution_retention_6d_cnt",
	"coupon",
	"attribution_retention_7d_cnt",
	"average_video_play",
	"attribution_active_pay_intra_one_day_count",
	"cpa",
	"play_duration_5s_rate",
	"click_shopwindow",
	"attribution_game_in_app_ltv_2days",
	"active_register_cost",
	"cpm",
	"pre_convert_rate",
	"attribution_game_in_app_roi_8days",
	"cost",
	"advanced_creative_counsel_click",
	"ies_music_click",
	"play_duration",
	"interact_per_cost",
	"submit_certification_count",
	"attribution_game_in_app_roi_3days",
	"dislike",
	"attribution_retention_2d_cost",
	"live_fans_club_join_cnt",
	"union_roi_0",
	"union_roi_3",
	"attribution_game_in_app_ltv_4days",
	"luban_live_gift_amount",
	"attribution_game_in_app_roi_2days",
	"install_finish_cost",
	"attribution_retention_3d_cnt",
	"advanced_creative_form_click",
	"click_call_dy",
	"valid_play",
	"attribution_game_pay_7d_cost",
	"attribution_next_day_open_cost",
	"poi_collect",
	"next_day_open",
	"attribution_active_pay_7d_rate",
	"active_pay_cost",
	"phone_confirm",
	"wechat_first_pay_rate",
	"next_day_open_rate",
	"card_show",
	"attribution_day_acitve_pay_count",
	"install_finish",
	"play_duration_3s_rate",
	"play_duration_10s",
	"attribution_retention_7d_sum_cnt",
	"attribution_game_in_app_ltv_8days",
	"first_rental_order_count",
	"pre_convert_cost",
	"active_pay_amount",
	"pre_convert_count",
	"attribution_wechat_login_30d_cost",
	"attribution_game_in_app_ltv_5days",
	"attribution_wechat_first_pay_30d_count",
	"attribution_active_pay_7d_count",
	"wechat_login_count",
	"coupon_single_page",
	"attribution_retention_6d_cost",
	"luban_live_follow_cnt",
	"active_rate",
	"commute_first_pay_count",
	"play_duration_3s",
	"play_75_feed_break_rate",
	"attribution_micro_game_3d_ltv",
	"attribution_wechat_first_pay_30d_rate",
	"luban_live_pay_order_count",
	"click_website",
	"avg_show_cost",
	"ctr",
	"attribution_retention_2d_cnt",
	"convert_show_rate",
	"wifi_play",
	"luban_order_stat_amount",
	"attribution_game_in_app_ltv_1day",
	"live_watch_one_minute_count",
	"like",
	"live_component_click_cost",
	"attribution_active_pay_intra_one_day_roi",
	"attribution_next_day_open_cnt",
	"attribution_wechat_pay_30d_roi",
	"pre_loan_credit",
	"redirect",
	"attribution_next_day_open_rate",
	"average_play_time_per_play",
	"live_component_click_rate",
	"form",
	"attribution_retention_5d_rate",
	"in_app_uv",
	"convert_cost",
	"download",
	"play_75_feed_break",
	"stat_union_ltv_3",
	"shopping",
	"play_50_feed_break_rate",
	"game_addiction",
	"average_play_progress",
	"game_addiction_cost",
	"attribution_convert",
	"attribution_game_in_app_roi_6days",
	"attribution_micro_game_7d_ltv",
	"advanced_creative_coupon_addition",
	"click",
	"cpc",
	"attribution_micro_game_3d_roi",
	"valid_play_cost",
	"wechat",
	"ies_challenge_click",
	"convert",
	"attribution_game_in_app_ltv_7days",
	"wechat_pay_amount",
	"first_order_count",
	"attribution_game_in_app_roi_5days",
	"in_app_cart",
	"show",
	"attribution_wechat_pay_30d_amount",
	"active_register_rate",
	"attribution_game_in_app_roi_1day",
	"luban_live_pay_order_stat_cost",
	"attribution_deep_convert",
	"register",
	"pay_count",
	"download_finish_rate",
	"play_over_rate",
}

// NewReportAdGetV2OrderFieldFromValue returns a pointer to a valid ReportAdGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2OrderFieldFromValue(v string) (*ReportAdGetV2OrderField, error) {
	ev := ReportAdGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2OrderField: valid values are %v", v, AllowedReportAdGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_order_field value
func (v ReportAdGetV2OrderField) Ptr() *ReportAdGetV2OrderField {
	return &v
}
