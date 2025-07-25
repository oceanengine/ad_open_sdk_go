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
	DEEP_CONVERT_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_rate"
	STAT_UNION_LTV_7_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_7"
	ATTRIBUTION_RETENTION_4D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_3days"
	VIEW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "view"
	VALID_PLAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_rate"
	PLAY_DURATION_SUM_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_sum"
	PLAY_DURATION_10S_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "play_duration_10s"
	CONSULT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "consult"
	COMMUTE_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "commute_first_pay_count"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportCampaignGetV2OrderField ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "attribution_active_pay_7d_count"
	ATTRIBUTION_RETENTION_3D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_3d_cnt"
	ATTRIBUTION_RETENTION_5D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_cost"
	GAME_ADDICTION_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_cost"
	ATTRIBUTION_RETENTION_3D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_rate"
	LIVE_FANS_CLUB_JOIN_CNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "live_fans_club_join_cnt"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	WECHAT_FIRST_PAY_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_7days"
	FIRST_RENTAL_ORDER_COUNT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "first_rental_order_count"
	PLAY_DURATION_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "play_duration"
	PLAY_OVER_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "play_over"
	PLAY_DURATION_3S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_4days"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_3days"
	ACTIVE_PAY_RATE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_rate"
	PLAY_DURATION_2S_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "play_duration_2s"
	ACTIVE_COST_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_cost"
	PRE_LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "pre_loan_credit_cost"
	AVG_RANK_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "avg_rank"
	MESSAGE_ACTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "message_action"
	WIFI_PLAY_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "wifi_play_rate"
	WECHAT_FIRST_PAY_COST_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "wechat_first_pay_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_6days"
	REPORT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "report"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_cost"
	GAME_PAY_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "game_pay_cost"
	GAME_ADDICTION_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "game_addiction_rate"
	PLAY_75_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_75_feed_break"
	AVERAGE_PLAY_PROGRESS_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "average_play_progress"
	WECHAT_LOGIN_COUNT_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "wechat_login_count"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_6days"
	WECHAT_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "wechat"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_next_day_open_cnt"
	WIFI_PLAY_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "wifi_play"
	LUBAN_ORDER_STAT_AMOUNT_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "luban_order_stat_amount"
	PRE_CONVERT_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "pre_convert_count"
	LUBAN_LIVE_FOLLOW_CNT_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "luban_live_follow_cnt"
	GAME_PAY_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_pay_count"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_amount"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_ltv"
	AVG_CLICK_COST_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "avg_click_cost"
	IN_APP_UV_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "in_app_uv"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_roi"
	PLAY_DURATION_10S_RATE_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "play_duration_10s_rate"
	CLICK_CALL_DY_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_call_dy"
	ATTRIBUTION_RETENTION_2D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_rate"
	AVERAGE_VIDEO_PLAY_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "average_video_play"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_5days"
	LOAN_COMPLETION_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "loan_completion"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	PLAY_25_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_25_feed_break"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_pay_order_stat_cost"
	LIKE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "like"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportCampaignGetV2OrderField  ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	CONVERT_COST_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_cost"
	ACTIVE_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "active"
	PHONE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "phone"
	ATTRIBUTION_RETENTION_7D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_7d_cnt"
	DOWNLOAD_START_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "download_start"
	ATTRIBUTION_RETENTION_4D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_4d_rate"
	LIVE_COMPONENT_CLICK_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "live_component_click_count"
	CPA_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpa"
	PLAY_DURATION_3S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_3s_rate"
	ATTRIBUTION_RETENTION_5D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_5d_rate"
	ACTIVE_PAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "active_pay_cost"
	REGISTER_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "register"
	ACTIVE_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "active_pay_amount"
	QQ_ReportCampaignGetV2OrderField                                          ReportCampaignGetV2OrderField = "qq"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportCampaignGetV2OrderField                 ReportCampaignGetV2OrderField = "live_watch_one_minute_count"
	PLAY_25_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_25_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_roi"
	INSTALL_FINISH_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "install_finish"
	INSTALL_FINISH_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_rate"
	ADVANCED_CREATIVE_FORM_CLICK_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "advanced_creative_form_click"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_4days"
	ATTRIBUTION_RETENTION_6D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_rate"
	ATTRIBUTION_RETENTION_2D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_2d_cost"
	APPROVAL_COUNT_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "approval_count"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "luban_live_slidecart_click_cnt"
	LOAN_CREDIT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "loan_credit"
	LUBAN_ORDER_ROI_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_roi"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_1day"
	ATTRIBUTION_RETENTION_5D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_5d_cnt"
	TOTAL_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "total_play"
	CLICK_LANDING_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "click_landing_page"
	ATTRIBUTION_RETENTION_6D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportCampaignGetV2OrderField      ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	ATTRIBUTION_CONVERT_COST_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_convert_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_8days"
	CUSTOMER_EFFECTIVE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "customer_effective"
	BUTTON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "button"
	CPC_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpc"
	ATTRIBUTION_RETENTION_6D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_6d_cost"
	SHOPPING_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "shopping"
	COUPON_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "coupon"
	VALID_PLAY_COST_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "valid_play_cost"
	SHARE_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "share"
	STAT_UNION_LTV_0_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_0"
	INTERACT_PER_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "interact_per_cost"
	DOWNLOAD_START_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_cost"
	PLAY_100_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportCampaignGetV2OrderField       ReportCampaignGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	POI_COLLECT_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "poi_collect"
	ACTIVE_RATE_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "active_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_7days"
	PRE_LOAN_CREDIT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "pre_loan_credit"
	UNION_ROI_3_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_3"
	PLAY_OVER_RATE_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "play_over_rate"
	LOAN_CREDIT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_cost"
	IN_APP_PAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "in_app_pay"
	CONVERT_SHOW_RATE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "convert_show_rate"
	LIVE_COMPONENT_CLICK_COST_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_cost"
	CLICK_DOWNLOAD_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "click_download"
	ACTIVE_REGISTER_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_rate"
	PHONE_CONNECT_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_connect"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_0d_roi"
	REDIRECT_TO_SHOP_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "redirect_to_shop"
	LUBAN_LIVE_ENTER_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_enter_cnt"
	LOAN_COMPLETION_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_rate"
	DOWNLOAD_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "download"
	MESSAGE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "message"
	LUBAN_LIVE_GIFT_CNT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "luban_live_gift_cnt"
	CLICK_INSTALL_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_install"
	DEEP_CONVERT_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "deep_convert"
	LUBAN_LIVE_GIFT_AMOUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_gift_amount"
	DOWNLOAD_FINISH_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "download_finish"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_phone_click"
	DEEP_CONVERT_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "deep_convert_cost"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_7d_ltv"
	GAME_ADDICTION_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "game_addiction"
	UNION_ROI_7_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_7"
	HOME_VISITED_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "home_visited"
	WECHAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_pay_amount"
	INSTALL_FINISH_COST_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "install_finish_cost"
	IN_APP_DETAIL_UV_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "in_app_detail_uv"
	WECHAT_FIRST_PAY_COUNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "wechat_first_pay_count"
	NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_active_pay_7d_per_count"
	FOLLOW_ReportCampaignGetV2OrderField                                      ReportCampaignGetV2OrderField = "follow"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_wechat_pay_30d_roi"
	DOWNLOAD_FINISH_RATE_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_rate"
	IN_APP_CART_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "in_app_cart"
	MAP_SEARCH_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "map_search"
	FORM_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "form"
	ACTIVE_REGISTER_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "active_register_cost"
	COUPON_SINGLE_PAGE_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "coupon_single_page"
	PRE_CONVERT_COST_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_rate"
	NEXT_DAY_OPEN_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "next_day_open"
	LOAN_COMPLETION_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "loan_completion_cost"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_2days"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "average_play_time_per_play"
	ATTRIBUTION_RETENTION_7D_RATE_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_rate"
	COST_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "cost"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_next_day_open_rate"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_game_pay_7d_cost"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "luban_live_pay_order_count"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_deep_convert_cost"
	IES_CHALLENGE_CLICK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "ies_challenge_click"
	PRE_CONVERT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "pre_convert_rate"
	ATTRIBUTION_RETENTION_3D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportCampaignGetV2OrderField    ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_acitve_pay_count"
	STAT_PAY_AMOUNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "stat_pay_amount"
	PLAY_50_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_50_feed_break_rate"
	DOWNLOAD_START_RATE_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "download_start_rate"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "luban_live_click_product_cnt"
	LOAN_CREDIT_RATE_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "loan_credit_rate"
	NEXT_DAY_OPEN_COST_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "next_day_open_cost"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "advanced_creative_form_submit"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_micro_game_3d_ltv"
	CONSULT_EFFECTIVE_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "consult_effective"
	SHOW_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "show"
	REDIRECT_ReportCampaignGetV2OrderField                                    ReportCampaignGetV2OrderField = "redirect"
	CLICK_WEBSITE_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "click_website"
	CTR_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "ctr"
	IES_MUSIC_CLICK_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "ies_music_click"
	PLAY_DURATION_5S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_5s_rate"
	CONVERT_RATE_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "convert_rate"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_day_active_pay_count"
	PAY_AMOUNT_ROI_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "pay_amount_roi"
	LUBAN_LIVE_COMMENT_CNT_ReportCampaignGetV2OrderField                      ReportCampaignGetV2OrderField = "luban_live_comment_cnt"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportCampaignGetV2OrderField          ReportCampaignGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_2d_cnt"
	PHONE_EFFECTIVE_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "phone_effective"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_roi_5days"
	CPM_ReportCampaignGetV2OrderField                                         ReportCampaignGetV2OrderField = "cpm"
	PLAY_75_FEED_BREAK_RATE_ReportCampaignGetV2OrderField                     ReportCampaignGetV2OrderField = "play_75_feed_break_rate"
	PLAY_100_FEED_BREAK_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "play_100_feed_break"
	COMMENT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "comment"
	LIVE_COMPONENT_CLICK_RATE_ReportCampaignGetV2OrderField                   ReportCampaignGetV2OrderField = "live_component_click_rate"
	POI_ADDRESS_CLICK_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "poi_address_click"
	LOCATION_CLICK_ReportCampaignGetV2OrderField                              ReportCampaignGetV2OrderField = "location_click"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportCampaignGetV2OrderField             ReportCampaignGetV2OrderField = "advanced_creative_counsel_click"
	PAY_COUNT_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "pay_count"
	LUBAN_ORDER_CNT_ReportCampaignGetV2OrderField                             ReportCampaignGetV2OrderField = "luban_order_cnt"
	DISLIKE_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "dislike"
	SUBMIT_CERTIFICATION_COUNT_ReportCampaignGetV2OrderField                  ReportCampaignGetV2OrderField = "submit_certification_count"
	CARD_SHOW_ReportCampaignGetV2OrderField                                   ReportCampaignGetV2OrderField = "card_show"
	AVG_SHOW_COST_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "avg_show_cost"
	ATTRIBUTION_CONVERT_ReportCampaignGetV2OrderField                         ReportCampaignGetV2OrderField = "attribution_convert"
	CLICK_SHOPWINDOW_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "click_shopwindow"
	ATTRIBUTION_RETENTION_4D_CNT_ReportCampaignGetV2OrderField                ReportCampaignGetV2OrderField = "attribution_retention_4d_cnt"
	LUBAN_LIVE_SHARE_CNT_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "luban_live_share_cnt"
	DOWNLOAD_FINISH_COST_ReportCampaignGetV2OrderField                        ReportCampaignGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportCampaignGetV2OrderField         ReportCampaignGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportCampaignGetV2OrderField   ReportCampaignGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	VOTE_ReportCampaignGetV2OrderField                                        ReportCampaignGetV2OrderField = "vote"
	CONVERT_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "convert"
	ATTRIBUTION_RETENTION_7D_COST_ReportCampaignGetV2OrderField               ReportCampaignGetV2OrderField = "attribution_retention_7d_cost"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "advanced_creative_coupon_addition"
	UNION_ROI_0_ReportCampaignGetV2OrderField                                 ReportCampaignGetV2OrderField = "union_roi_0"
	IN_APP_ORDER_ReportCampaignGetV2OrderField                                ReportCampaignGetV2OrderField = "in_app_order"
	VALID_PLAY_ReportCampaignGetV2OrderField                                  ReportCampaignGetV2OrderField = "valid_play"
	WECHAT_LOGIN_COST_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "wechat_login_cost"
	FIRST_ORDER_COUNT_ReportCampaignGetV2OrderField                           ReportCampaignGetV2OrderField = "first_order_count"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_1day"
	LOTTERY_ReportCampaignGetV2OrderField                                     ReportCampaignGetV2OrderField = "lottery"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_8days"
	STAT_UNION_LTV_3_ReportCampaignGetV2OrderField                            ReportCampaignGetV2OrderField = "stat_union_ltv_3"
	CLICK_ReportCampaignGetV2OrderField                                       ReportCampaignGetV2OrderField = "click"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportCampaignGetV2OrderField              ReportCampaignGetV2OrderField = "attribution_active_pay_7d_cost"
	PLAY_50_FEED_BREAK_ReportCampaignGetV2OrderField                          ReportCampaignGetV2OrderField = "play_50_feed_break"
	PLAY_DURATION_2S_RATE_ReportCampaignGetV2OrderField                       ReportCampaignGetV2OrderField = "play_duration_2s_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportCampaignGetV2OrderField           ReportCampaignGetV2OrderField = "attribution_game_in_app_ltv_2days"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportCampaignGetV2OrderField            ReportCampaignGetV2OrderField = "attribution_retention_7d_sum_cnt"
	PHONE_CONFIRM_ReportCampaignGetV2OrderField                               ReportCampaignGetV2OrderField = "phone_confirm"
	ATTRIBUTION_DEEP_CONVERT_ReportCampaignGetV2OrderField                    ReportCampaignGetV2OrderField = "attribution_deep_convert"
)

// Ptr returns reference to report_campaign_get_v2_order_field value
func (v ReportCampaignGetV2OrderField) Ptr() *ReportCampaignGetV2OrderField {
	return &v
}
