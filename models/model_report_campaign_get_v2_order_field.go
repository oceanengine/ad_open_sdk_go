/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportCampaignGetV2OrderField
type ReportCampaignGetV2OrderField string

// List of report_campaign_get_v2_order_field
const (
	CLICK_DOWNLOAD_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "click_download"
	IN_APP_ORDER_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "in_app_order"
	NEXT_DAY_OPEN_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "next_day_open"
	LOAN_COMPLETION_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_cost"
	PHONE_CONNECT_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_connect"
	PLAY_DURATION_3S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_CONVERT_COST_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_convert_cost"
	ACTIVE_REGISTER_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_rate"
	ATTRIBUTION_DEEP_CONVERT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_deep_convert"
	PLAY_75_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_75_feed_break_rate"
	COST_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "cost"
	ATTRIBUTION_RETENTION_4D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_cost"
	GAME_PAY_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_pay_count"
	CPC_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpc"
	SUBMIT_CERTIFICATION_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "submit_certification_count"
	STAT_UNION_LTV_0_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_0"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_game_pay_7d_count"
	LUBAN_ORDER_STAT_AMOUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "luban_order_stat_amount"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_phone_click"
	CLICK_SHOPWINDOW_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_roi"
	CTR_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "ctr"
	REDIRECT_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "redirect"
	BUTTON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "button"
	AVG_SHOW_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "avg_show_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_form_submit"
	QQ_ReportCampaignGetV2OrderField                                          ReportCampaignGetV2OrderField = "qq"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_roi"
	LUBAN_LIVE_FOLLOW_CNT_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "luban_live_follow_cnt"
	HOME_VISITED_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "home_visited"
	LOAN_CREDIT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_rate"
	AVG_RANK_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "avg_rank"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "average_play_time_per_play"
	MESSAGE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "message"
	IES_MUSIC_CLICK_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "ies_music_click"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_rate"
	CONVERT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "convert"
	CONSULT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "consult"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_4days"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCampaignGetV2OrderField                 ReportCampaignGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_rate"
	INSTALL_FINISH_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_rate"
	LUBAN_LIVE_SHARE_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_share_cnt"
	WECHAT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "wechat"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_pay_order_stat_cost"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_6d_cnt"
	CLICK_LANDING_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "click_landing_page"
	COUPON_SINGLE_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "coupon_single_page"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	PLAY_25_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_25_feed_break"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	PLAY_DURATION_5S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_5s_rate"
	NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_rate"
	LOAN_CREDIT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "loan_credit"
	DOWNLOAD_FINISH_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_cost"
	PAY_COUNT_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "pay_count"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_cost"
	ATTRIBUTION_RETENTION_6D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_8days"
	SHOW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "show"
	GAME_PAY_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCampaignGetV2OrderField  ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_RETENTION_3D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_cost"
	GAME_ADDICTION_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_rate"
	ATTRIBUTION_RETENTION_2D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_cost"
	PLAY_100_FEED_BREAK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "play_100_feed_break"
	PLAY_OVER_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "play_over"
	LUBAN_LIVE_ENTER_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_enter_cnt"
	DISLIKE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "dislike"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "attribution_active_pay_7d_count"
	INTERACT_PER_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_roi"
	STAT_UNION_LTV_3_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_3"
	CLICK_INSTALL_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_install"
	ACTIVE_COST_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_6days"
	POI_COLLECT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "poi_collect"
	LUBAN_ORDER_CNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_cnt"
	VALID_PLAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_rate"
	WECHAT_LOGIN_COUNT_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "wechat_login_count"
	SHARE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "share"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_ltv"
	DOWNLOAD_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "download"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCampaignGetV2OrderField ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	VALID_PLAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_cost"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_acitve_pay_count"
	LUBAN_LIVE_GIFT_CNT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "luban_live_gift_cnt"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCampaignGetV2OrderField    ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "luban_live_pay_order_count"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_cost"
	IN_APP_CART_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "in_app_cart"
	PLAY_50_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_RETENTION_5D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_8days"
	CONVERT_COST_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_cost"
	LIVE_COMPONENT_CLICK_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "live_component_click_count"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_5days"
	LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_cost"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "advanced_creative_form_click"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "advanced_creative_counsel_click"
	WIFI_PLAY_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "wifi_play_rate"
	PRE_LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "pre_loan_credit_cost"
	FIRST_RENTAL_ORDER_COUNT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "first_rental_order_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_2d_cnt"
	DOWNLOAD_START_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "download_start"
	PHONE_CONFIRM_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_confirm"
	UNION_ROI_3_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_3"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCampaignGetV2OrderField      ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_2days"
	TOTAL_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "total_play"
	PLAY_DURATION_10S_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_10s"
	CUSTOMER_EFFECTIVE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "customer_effective"
	CLICK_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "click"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_5d_cnt"
	WECHAT_LOGIN_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_login_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_next_day_open_cnt"
	PHONE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "phone"
	IES_CHALLENGE_CLICK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "ies_challenge_click"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_3d_cnt"
	REGISTER_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "register"
	WECHAT_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "wechat_first_pay_count"
	INSTALL_FINISH_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "install_finish"
	WECHAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_pay_amount"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_ltv"
	PLAY_100_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "play_100_feed_break_rate"
	LIVE_COMPONENT_CLICK_COST_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_cost"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_amount"
	VALID_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "valid_play"
	LUBAN_LIVE_COMMENT_CNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_comment_cnt"
	VIEW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "view"
	IN_APP_DETAIL_UV_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "in_app_detail_uv"
	PLAY_OVER_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "play_over_rate"
	ACTIVE_RATE_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_rate"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCampaignGetV2OrderField          ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_count"
	FOLLOW_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "follow"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_active_pay_count"
	GAME_ADDICTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_addiction"
	CLICK_WEBSITE_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_website"
	AVG_CLICK_COST_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "avg_click_cost"
	PRE_LOAN_CREDIT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "pre_loan_credit"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_1day"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_rate"
	LOTTERY_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "lottery"
	PLAY_DURATION_SUM_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_sum"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_3days"
	ACTIVE_PAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_rate"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_slidecart_click_cnt"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_rate"
	LIKE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "like"
	PLAY_DURATION_3S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_3s_rate"
	WECHAT_FIRST_PAY_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_rate"
	STAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "stat_pay_amount"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "advanced_creative_coupon_addition"
	PLAY_50_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_50_feed_break"
	MESSAGE_ACTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "message_action"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_deep_convert_cost"
	FORM_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "form"
	DEEP_CONVERT_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_cost"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_roi"
	MAP_SEARCH_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "map_search"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_7days"
	CONVERT_RATE_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_rate"
	UNION_ROI_7_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_7"
	COUPON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "coupon"
	PRE_CONVERT_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "pre_convert_count"
	PAY_AMOUNT_ROI_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "pay_amount_roi"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_retention_7d_sum_cnt"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_gift_amount"
	CONVERT_SHOW_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_2days"
	WECHAT_FIRST_PAY_COST_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_cost"
	PLAY_75_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_75_feed_break"
	CONSULT_EFFECTIVE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "consult_effective"
	DEEP_CONVERT_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "deep_convert"
	FIRST_ORDER_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "first_order_count"
	PLAY_25_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_25_feed_break_rate"
	IN_APP_UV_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "in_app_uv"
	COMMENT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "comment"
	ACTIVE_PAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_cost"
	INSTALL_FINISH_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_1day"
	GAME_ADDICTION_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_cost"
	UNION_ROI_0_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_0"
	SHOPPING_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "shopping"
	CLICK_CALL_DY_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_call_dy"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_rate"
	POI_ADDRESS_CLICK_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "poi_address_click"
	AVERAGE_VIDEO_PLAY_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "average_video_play"
	LIVE_COMPONENT_CLICK_RATE_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_ltv"
	REDIRECT_TO_SHOP_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "redirect_to_shop"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "live_fans_club_join_cnt"
	VOTE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "vote"
	COMMUTE_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "commute_first_pay_count"
	DOWNLOAD_FINISH_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_rate"
	LUBAN_ORDER_ROI_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_roi"
	PHONE_EFFECTIVE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "phone_effective"
	STAT_UNION_LTV_7_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_7"
	AVERAGE_PLAY_PROGRESS_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "average_play_progress"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	DOWNLOAD_START_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_rate"
	IN_APP_PAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "in_app_pay"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_5days"
	PRE_CONVERT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_cost"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_7d_cnt"
	DEEP_CONVERT_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_rate"
	PLAY_DURATION_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "play_duration"
	PLAY_DURATION_2S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_2s"
	ACTIVE_REGISTER_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_cost"
	CPM_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpm"
	PLAY_DURATION_10S_RATE_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "play_duration_10s_rate"
	CPA_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpa"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_retention_7d_total_cost"
	ACTIVE_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "active_pay_amount"
	ATTRIBUTION_RETENTION_7D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_cost"
	NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_cost"
	LOCATION_CLICK_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "location_click"
	WIFI_PLAY_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "wifi_play"
	DOWNLOAD_START_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_3days"
	DOWNLOAD_FINISH_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "download_finish"
	LOAN_COMPLETION_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "loan_completion"
	CARD_SHOW_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "card_show"
	APPROVAL_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "approval_count"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_6days"
	PLAY_DURATION_2S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_2s_rate"
	ACTIVE_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "active"
	REPORT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "report"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_active_pay_7d_per_count"
	LOAN_COMPLETION_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_rate"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "luban_live_click_product_cnt"
	PRE_CONVERT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_CONVERT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "attribution_convert"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_4days"
)

// All allowed values of ReportCampaignGetV2OrderField enum
var AllowedReportCampaignGetV2OrderFieldEnumValues = []ReportCampaignGetV2OrderField{
	"click_download",
	"in_app_order",
	"next_day_open",
	"loan_completion_cost",
	"phone_connect",
	"play_duration_3s",
	"attribution_convert_cost",
	"active_register_rate",
	"attribution_deep_convert",
	"play_75_feed_break_rate",
	"cost",
	"attribution_retention_4d_cost",
	"game_pay_count",
	"cpc",
	"submit_certification_count",
	"stat_union_ltv_0",
	"attribution_game_pay_7d_count",
	"luban_order_stat_amount",
	"advanced_creative_phone_click",
	"click_shopwindow",
	"attribution_micro_game_0d_roi",
	"ctr",
	"redirect",
	"button",
	"avg_show_cost",
	"advanced_creative_form_submit",
	"qq",
	"attribution_micro_game_7d_roi",
	"luban_live_follow_cnt",
	"home_visited",
	"loan_credit_rate",
	"avg_rank",
	"average_play_time_per_play",
	"message",
	"ies_music_click",
	"attribution_retention_5d_rate",
	"convert",
	"consult",
	"attribution_next_day_open_rate",
	"attribution_game_in_app_roi_4days",
	"live_watch_one_minute_count",
	"attribution_retention_7d_rate",
	"install_finish_rate",
	"luban_live_share_cnt",
	"wechat",
	"luban_live_pay_order_stat_cost",
	"attribution_retention_4d_cnt",
	"attribution_retention_6d_cnt",
	"click_landing_page",
	"coupon_single_page",
	"attribution_wechat_first_pay_30d_cost",
	"play_25_feed_break",
	"attribution_active_pay_intra_one_day_rate",
	"play_duration_5s_rate",
	"next_day_open_rate",
	"loan_credit",
	"download_finish_cost",
	"pay_count",
	"attribution_wechat_login_30d_cost",
	"attribution_retention_6d_cost",
	"attribution_game_in_app_ltv_8days",
	"show",
	"game_pay_cost",
	"attribution_active_pay_intra_one_day_count",
	"attribution_retention_3d_cost",
	"game_addiction_rate",
	"attribution_retention_2d_cost",
	"play_100_feed_break",
	"play_over",
	"luban_live_enter_cnt",
	"dislike",
	"attribution_active_pay_7d_count",
	"interact_per_cost",
	"attribution_micro_game_3d_roi",
	"stat_union_ltv_3",
	"click_install",
	"active_cost",
	"attribution_active_pay_7d_cost",
	"attribution_game_in_app_ltv_6days",
	"poi_collect",
	"luban_order_cnt",
	"valid_play_rate",
	"wechat_login_count",
	"share",
	"attribution_micro_game_0d_ltv",
	"download",
	"attribution_retention_2d_rate",
	"attribution_active_pay_intra_one_day_amount",
	"valid_play_cost",
	"attribution_day_acitve_pay_count",
	"luban_live_gift_cnt",
	"attribution_active_pay_intra_one_day_roi",
	"luban_live_pay_order_count",
	"attribution_active_pay_7d_rate",
	"attribution_next_day_open_cost",
	"in_app_cart",
	"play_50_feed_break_rate",
	"attribution_retention_5d_cost",
	"attribution_game_in_app_roi_8days",
	"convert_cost",
	"live_component_click_count",
	"attribution_game_in_app_roi_5days",
	"loan_credit_cost",
	"advanced_creative_form_click",
	"advanced_creative_counsel_click",
	"wifi_play_rate",
	"pre_loan_credit_cost",
	"first_rental_order_count",
	"attribution_retention_2d_cnt",
	"download_start",
	"phone_confirm",
	"union_roi_3",
	"attribution_wechat_first_pay_30d_count",
	"attribution_game_in_app_roi_2days",
	"total_play",
	"play_duration_10s",
	"customer_effective",
	"click",
	"attribution_retention_5d_cnt",
	"wechat_login_cost",
	"attribution_next_day_open_cnt",
	"phone",
	"ies_challenge_click",
	"attribution_retention_3d_cnt",
	"register",
	"wechat_first_pay_count",
	"install_finish",
	"wechat_pay_amount",
	"attribution_micro_game_7d_ltv",
	"play_100_feed_break_rate",
	"live_component_click_cost",
	"attribution_wechat_pay_30d_amount",
	"valid_play",
	"luban_live_comment_cnt",
	"view",
	"in_app_detail_uv",
	"play_over_rate",
	"active_rate",
	"attribution_wechat_login_30d_count",
	"follow",
	"attribution_day_active_pay_count",
	"game_addiction",
	"click_website",
	"avg_click_cost",
	"pre_loan_credit",
	"attribution_game_in_app_roi_1day",
	"attribution_retention_4d_rate",
	"lottery",
	"play_duration_sum",
	"attribution_game_in_app_roi_3days",
	"active_pay_rate",
	"luban_live_slidecart_click_cnt",
	"attribution_retention_6d_rate",
	"like",
	"play_duration_3s_rate",
	"wechat_first_pay_rate",
	"stat_pay_amount",
	"advanced_creative_coupon_addition",
	"play_50_feed_break",
	"message_action",
	"attribution_deep_convert_cost",
	"form",
	"deep_convert_cost",
	"attribution_wechat_pay_30d_roi",
	"map_search",
	"attribution_game_in_app_roi_7days",
	"convert_rate",
	"union_roi_7",
	"coupon",
	"pre_convert_count",
	"pay_amount_roi",
	"attribution_wechat_first_pay_30d_rate",
	"attribution_retention_7d_sum_cnt",
	"luban_live_gift_amount",
	"convert_show_rate",
	"attribution_game_in_app_ltv_2days",
	"wechat_first_pay_cost",
	"play_75_feed_break",
	"consult_effective",
	"deep_convert",
	"first_order_count",
	"play_25_feed_break_rate",
	"in_app_uv",
	"comment",
	"active_pay_cost",
	"install_finish_cost",
	"attribution_game_in_app_ltv_1day",
	"game_addiction_cost",
	"union_roi_0",
	"shopping",
	"click_call_dy",
	"attribution_retention_3d_rate",
	"poi_address_click",
	"average_video_play",
	"live_component_click_rate",
	"attribution_micro_game_3d_ltv",
	"redirect_to_shop",
	"live_fans_club_join_cnt",
	"vote",
	"commute_first_pay_count",
	"download_finish_rate",
	"luban_order_roi",
	"phone_effective",
	"stat_union_ltv_7",
	"average_play_progress",
	"attribution_active_pay_intra_one_day_cost",
	"download_start_rate",
	"in_app_pay",
	"attribution_game_in_app_ltv_5days",
	"pre_convert_cost",
	"attribution_retention_7d_cnt",
	"deep_convert_rate",
	"play_duration",
	"play_duration_2s",
	"active_register_cost",
	"cpm",
	"play_duration_10s_rate",
	"cpa",
	"attribution_retention_7d_total_cost",
	"active_pay_amount",
	"attribution_retention_7d_cost",
	"next_day_open_cost",
	"location_click",
	"wifi_play",
	"download_start_cost",
	"attribution_game_in_app_ltv_3days",
	"download_finish",
	"loan_completion",
	"card_show",
	"approval_count",
	"attribution_game_pay_7d_cost",
	"attribution_game_in_app_ltv_7days",
	"attribution_game_in_app_roi_6days",
	"play_duration_2s_rate",
	"active",
	"report",
	"attribution_active_pay_7d_per_count",
	"loan_completion_rate",
	"luban_live_click_product_cnt",
	"pre_convert_rate",
	"attribution_convert",
	"attribution_game_in_app_ltv_4days",
}

// NewReportCampaignGetV2OrderFieldFromValue returns a pointer to a valid ReportCampaignGetV2OrderField
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCampaignGetV2OrderFieldFromValue(v string) (*ReportCampaignGetV2OrderField, error) {
	ev := ReportCampaignGetV2OrderField(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCampaignGetV2OrderField: valid values are %v", v, AllowedReportCampaignGetV2OrderFieldEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCampaignGetV2OrderField) IsValid() bool {
	for _, existing := range AllowedReportCampaignGetV2OrderFieldEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_campaign_get_v2_order_field value
func (v ReportCampaignGetV2OrderField) Ptr() *ReportCampaignGetV2OrderField {
	return &v
}
