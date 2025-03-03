/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCampaignGetV2OrderField
type ReportCampaignGetV2OrderField string

// List of report_campaign_get_v2_order_field
const (
	ATTRIBUTION_RETENTION_2D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_cost"
	PHONE_CONNECT_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_connect"
	CONVERT_COST_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_cost"
	WECHAT_FIRST_PAY_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_rate"
	PHONE_EFFECTIVE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "phone_effective"
	GAME_ADDICTION_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_rate"
	COUPON_SINGLE_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "coupon_single_page"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_6days"
	DEEP_CONVERT_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "deep_convert"
	LOAN_COMPLETION_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_rate"
	IN_APP_ORDER_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "in_app_order"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_4days"
	LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_cost"
	UNION_ROI_7_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_7"
	LUBAN_ORDER_STAT_AMOUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "luban_order_stat_amount"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_gift_amount"
	PHONE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "phone"
	PLAY_DURATION_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "play_duration"
	BUTTON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "button"
	DOWNLOAD_FINISH_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_7days"
	DISLIKE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "dislike"
	PAY_AMOUNT_ROI_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "pay_amount_roi"
	UNION_ROI_3_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_3"
	REGISTER_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "register"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_rate"
	WECHAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_pay_amount"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_roi"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_game_pay_7d_count"
	IN_APP_DETAIL_UV_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "in_app_detail_uv"
	MAP_SEARCH_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "map_search"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_cost"
	UNION_ROI_0_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_0"
	CLICK_CALL_DY_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_call_dy"
	MESSAGE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "message"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "advanced_creative_coupon_addition"
	ACTIVE_RATE_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_rate"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_5d_cnt"
	ATTRIBUTION_CONVERT_COST_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_cost"
	LOCATION_CLICK_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "location_click"
	PRE_CONVERT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_rate"
	PRE_CONVERT_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "pre_convert_count"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_deep_convert_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_6days"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_8days"
	CUSTOMER_EFFECTIVE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "customer_effective"
	GAME_ADDICTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_addiction"
	REDIRECT_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "redirect"
	VALID_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "valid_play"
	NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_cost"
	ATTRIBUTION_RETENTION_6D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_cost"
	PLAY_DURATION_2S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_2s"
	LIKE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "like"
	PRE_CONVERT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_cost"
	CLICK_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "click"
	CPA_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpa"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_2days"
	AVG_RANK_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "avg_rank"
	WECHAT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "wechat"
	LIVE_COMPONENT_CLICK_RATE_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_5days"
	INTERACT_PER_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "interact_per_cost"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_roi"
	COUPON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "coupon"
	CARD_SHOW_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "card_show"
	IN_APP_CART_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "in_app_cart"
	PLAY_100_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "play_100_feed_break_rate"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCampaignGetV2OrderField                 ReportCampaignGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_1day"
	ACTIVE_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "active"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_4d_cnt"
	INSTALL_FINISH_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_rate"
	CONVERT_RATE_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_rate"
	HOME_VISITED_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "home_visited"
	LOTTERY_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "lottery"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_roi"
	CLICK_LANDING_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "click_landing_page"
	STAT_UNION_LTV_3_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_3"
	CONSULT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "consult"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_amount"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "average_play_time_per_play"
	FIRST_RENTAL_ORDER_COUNT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "first_rental_order_count"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_acitve_pay_count"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_rate"
	VALID_PLAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_cost"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_next_day_open_cnt"
	AVERAGE_VIDEO_PLAY_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "average_video_play"
	ACTIVE_PAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_cost"
	PLAY_DURATION_10S_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_10s"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_active_pay_7d_per_count"
	ATTRIBUTION_RETENTION_5D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_cost"
	LOAN_COMPLETION_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "loan_completion"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_2d_cnt"
	CLICK_DOWNLOAD_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "click_download"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_roi"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCampaignGetV2OrderField  ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	ATTRIBUTION_RETENTION_4D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_cost"
	PRE_LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "pre_loan_credit_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_form_submit"
	PLAY_DURATION_5S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_5s_rate"
	NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_rate"
	FORM_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "form"
	NEXT_DAY_OPEN_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "next_day_open"
	PLAY_100_FEED_BREAK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "play_100_feed_break"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCampaignGetV2OrderField          ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_ltv"
	LUBAN_LIVE_SHARE_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_share_cnt"
	CTR_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "ctr"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_6d_cnt"
	IN_APP_PAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "in_app_pay"
	PLAY_50_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_50_feed_break"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_3days"
	LIVE_COMPONENT_CLICK_COST_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_cost"
	PLAY_75_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_75_feed_break"
	IES_MUSIC_CLICK_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "ies_music_click"
	DOWNLOAD_START_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_cost"
	CONVERT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "convert"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_retention_7d_total_cost"
	REDIRECT_TO_SHOP_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "redirect_to_shop"
	GAME_PAY_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_pay_count"
	STAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "stat_pay_amount"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_retention_7d_sum_cnt"
	CONSULT_EFFECTIVE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "consult_effective"
	WECHAT_LOGIN_COUNT_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "wechat_login_count"
	ATTRIBUTION_RETENTION_7D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_game_pay_7d_cost"
	PLAY_OVER_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "play_over"
	LUBAN_LIVE_COMMENT_CNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_comment_cnt"
	PLAY_DURATION_3S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_3s_rate"
	WIFI_PLAY_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "wifi_play_rate"
	WECHAT_LOGIN_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_login_cost"
	PLAY_DURATION_2S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_2s_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_active_pay_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCampaignGetV2OrderField    ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	GAME_ADDICTION_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_cost"
	LUBAN_LIVE_GIFT_CNT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "luban_live_gift_cnt"
	LUBAN_LIVE_ENTER_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_enter_cnt"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_2days"
	PLAY_OVER_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "play_over_rate"
	ATTRIBUTION_DEEP_CONVERT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_7days"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_ltv"
	AVG_CLICK_COST_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "avg_click_cost"
	DEEP_CONVERT_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_rate"
	DOWNLOAD_FINISH_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "download_finish"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_4days"
	DEEP_CONVERT_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_cost"
	SHARE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "share"
	SUBMIT_CERTIFICATION_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "submit_certification_count"
	CLICK_SHOPWINDOW_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "click_shopwindow"
	GAME_PAY_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_cost"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_phone_click"
	CPC_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpc"
	IN_APP_UV_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "in_app_uv"
	PLAY_DURATION_3S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_3s"
	WECHAT_FIRST_PAY_COST_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_cost"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_slidecart_click_cnt"
	MESSAGE_ACTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "message_action"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_rate"
	LUBAN_ORDER_CNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_cnt"
	AVERAGE_PLAY_PROGRESS_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "average_play_progress"
	TOTAL_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "total_play"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "luban_live_pay_order_count"
	LUBAN_ORDER_ROI_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_roi"
	REPORT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "report"
	FOLLOW_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "follow"
	POI_ADDRESS_CLICK_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "poi_address_click"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_5days"
	DOWNLOAD_START_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_rate"
	AVG_SHOW_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "avg_show_cost"
	WECHAT_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "wechat_first_pay_count"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_pay_order_stat_cost"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "advanced_creative_counsel_click"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCampaignGetV2OrderField      ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	VIEW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "view"
	FIRST_ORDER_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "first_order_count"
	PLAY_DURATION_10S_RATE_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "play_duration_10s_rate"
	SHOW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "show"
	ACTIVE_PAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_rate"
	COMMENT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "comment"
	LOAN_COMPLETION_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_cost"
	APPROVAL_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "approval_count"
	COMMUTE_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "commute_first_pay_count"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_rate"
	CPM_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpm"
	ATTRIBUTION_CONVERT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "attribution_convert"
	LUBAN_LIVE_FOLLOW_CNT_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "luban_live_follow_cnt"
	ACTIVE_COST_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_cost"
	QQ_ReportCampaignGetV2OrderField                                          ReportCampaignGetV2OrderField = "qq"
	CONVERT_SHOW_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "convert_show_rate"
	INSTALL_FINISH_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_cost"
	CLICK_INSTALL_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_install"
	ATTRIBUTION_RETENTION_3D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	PLAY_75_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_8days"
	PAY_COUNT_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "pay_count"
	PLAY_50_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_50_feed_break_rate"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_rate"
	COST_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "cost"
	PLAY_25_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_25_feed_break"
	CLICK_WEBSITE_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_website"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCampaignGetV2OrderField ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_rate"
	DOWNLOAD_START_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "download_start"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	STAT_UNION_LTV_0_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_0"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "live_fans_club_join_cnt"
	PLAY_DURATION_SUM_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_sum"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_3d_cnt"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_ltv"
	ACTIVE_REGISTER_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_cost"
	DOWNLOAD_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "download"
	PRE_LOAN_CREDIT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "pre_loan_credit"
	VALID_PLAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_rate"
	PHONE_CONFIRM_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_confirm"
	DOWNLOAD_FINISH_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_rate"
	WIFI_PLAY_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "wifi_play"
	ACTIVE_REGISTER_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_rate"
	LOAN_CREDIT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_rate"
	POI_COLLECT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "poi_collect"
	VOTE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "vote"
	ACTIVE_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "active_pay_amount"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "attribution_active_pay_7d_count"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "luban_live_click_product_cnt"
	SHOPPING_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "shopping"
	LOAN_CREDIT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "loan_credit"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_7d_cnt"
	STAT_UNION_LTV_7_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_7"
	LIVE_COMPONENT_CLICK_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "live_component_click_count"
	INSTALL_FINISH_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "install_finish"
	IES_CHALLENGE_CLICK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "ies_challenge_click"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "advanced_creative_form_click"
	PLAY_25_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_25_feed_break_rate"
)

// Ptr returns reference to report_campaign_get_v2_order_field value
func (v ReportCampaignGetV2OrderField) Ptr() *ReportCampaignGetV2OrderField {
	return &v
}
