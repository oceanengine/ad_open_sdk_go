/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdGetV2OrderField
type ReportAdGetV2OrderField string

// List of report_ad_get_v2_order_field
const (
	CLICK_CALL_DY_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_call_dy"
	INSTALL_FINISH_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_rate"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_7days"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_roi"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_rate"
	WECHAT_FIRST_PAY_COST_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_cost"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "live_component_click_count"
	WECHAT_LOGIN_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_login_cost"
	DOWNLOAD_START_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_cost"
	COMMENT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "comment"
	IES_CHALLENGE_CLICK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "ies_challenge_click"
	FIRST_ORDER_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "first_order_count"
	ATTRIBUTION_CONVERT_COST_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_convert_cost"
	SHOW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "show"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_rate"
	PLAY_DURATION_3S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_3s"
	PLAY_25_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_25_feed_break_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_8days"
	PRE_LOAN_CREDIT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "pre_loan_credit"
	GAME_PAY_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_pay_count"
	CLICK_SHOPWINDOW_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "click_shopwindow"
	AVERAGE_VIDEO_PLAY_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "average_video_play"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	SHARE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "share"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_5d_cnt"
	PHONE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "phone"
	ATTRIBUTION_CONVERT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "attribution_convert"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_ltv"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_8days"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	BUTTON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "button"
	HOME_VISITED_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "home_visited"
	PLAY_100_FEED_BREAK_RATE_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_ltv"
	WECHAT_FIRST_PAY_COUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "wechat_first_pay_count"
	LUBAN_LIVE_SHARE_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_share_cnt"
	CONVERT_SHOW_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "convert_show_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_rate"
	ACTIVE_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "active"
	GAME_ADDICTION_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_rate"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_rate"
	DOWNLOAD_START_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_rate"
	PLAY_OVER_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "play_over"
	SHOPPING_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "shopping"
	VALID_PLAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_rate"
	STAT_UNION_LTV_0_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_0"
	GAME_ADDICTION_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_cost"
	CPA_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpa"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_cost"
	QQ_ReportAdGetV2OrderField                                          ReportAdGetV2OrderField = "qq"
	CLICK_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "click"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "average_play_time_per_play"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_5days"
	DOWNLOAD_FINISH_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_cost"
	REGISTER_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "register"
	FORM_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "form"
	AVERAGE_PLAY_PROGRESS_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "average_play_progress"
	NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_rate"
	PLAY_DURATION_5S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_5s_rate"
	LOCATION_CLICK_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "location_click"
	LIVE_COMPONENT_CLICK_RATE_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_rate"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_game_pay_7d_cost"
	ACTIVE_PAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_cost"
	TOTAL_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "total_play"
	PLAY_DURATION_10S_RATE_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "play_duration_10s_rate"
	REDIRECT_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "redirect"
	PHONE_CONNECT_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_connect"
	IN_APP_PAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "in_app_pay"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_cost"
	LOAN_CREDIT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_cost"
	AVG_CLICK_COST_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "avg_click_cost"
	PHONE_CONFIRM_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_confirm"
	SUBMIT_CERTIFICATION_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "submit_certification_count"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_3days"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_phone_click"
	INSTALL_FINISH_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "install_finish"
	STAT_UNION_LTV_3_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_3"
	IN_APP_UV_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "in_app_uv"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "attribution_active_pay_7d_count"
	DEEP_CONVERT_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_cost"
	AVG_RANK_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "avg_rank"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_4days"
	COST_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "cost"
	WIFI_PLAY_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "wifi_play"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdGetV2OrderField          ReportAdGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_rate"
	PHONE_EFFECTIVE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "phone_effective"
	CONVERT_RATE_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_rate"
	COMMUTE_FIRST_PAY_COUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "commute_first_pay_count"
	WECHAT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "wechat"
	ACTIVE_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "active_pay_amount"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_pay_order_stat_cost"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_ltv"
	COUPON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "coupon"
	DOWNLOAD_FINISH_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_rate"
	WIFI_PLAY_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "wifi_play_rate"
	ACTIVE_REGISTER_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_rate"
	IN_APP_CART_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "in_app_cart"
	INSTALL_FINISH_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_cost"
	DISLIKE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "dislike"
	PLAY_50_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_50_feed_break_rate"
	FIRST_RENTAL_ORDER_COUNT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "first_rental_order_count"
	GAME_ADDICTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_addiction"
	ACTIVE_COST_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdGetV2OrderField  ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "luban_live_click_product_cnt"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_cost"
	CPM_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpm"
	CARD_SHOW_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "card_show"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "luban_live_follow_cnt"
	UNION_ROI_7_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_7"
	PRE_CONVERT_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "pre_convert_count"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_active_pay_7d_per_count"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_7d_cnt"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_pay_30d_amount"
	VOTE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "vote"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_roi"
	ACTIVE_RATE_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_rate"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_cost"
	CONSULT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "consult"
	PLAY_50_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_50_feed_break"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "luban_live_pay_order_count"
	PLAY_OVER_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_5days"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_2days"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_6days"
	LUBAN_ORDER_ROI_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_roi"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdGetV2OrderField      ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	PLAY_DURATION_SUM_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_sum"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	FOLLOW_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "follow"
	DOWNLOAD_FINISH_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "download_finish"
	NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_cost"
	UNION_ROI_3_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_3"
	LIKE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "like"
	PLAY_25_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_25_feed_break"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdGetV2OrderField                 ReportAdGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_2d_cnt"
	STAT_PAY_AMOUNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "stat_pay_amount"
	LUBAN_LIVE_ENTER_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_enter_cnt"
	MAP_SEARCH_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "map_search"
	LOAN_CREDIT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_rate"
	COUPON_SINGLE_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "coupon_single_page"
	REPORT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "report"
	WECHAT_FIRST_PAY_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_rate"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_rate"
	LOAN_COMPLETION_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_cost"
	CLICK_DOWNLOAD_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "click_download"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "advanced_creative_coupon_addition"
	AVG_SHOW_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "avg_show_cost"
	CTR_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "ctr"
	INTERACT_PER_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "interact_per_cost"
	REDIRECT_TO_SHOP_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "redirect_to_shop"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_form_submit"
	CONVERT_COST_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_cost"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_3d_cnt"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "advanced_creative_form_click"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdGetV2OrderField ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	PRE_CONVERT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_rate"
	WECHAT_LOGIN_COUNT_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "wechat_login_count"
	GAME_PAY_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "game_pay_cost"
	VIEW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "view"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_acitve_pay_count"
	LUBAN_LIVE_COMMENT_CNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_comment_cnt"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_slidecart_click_cnt"
	CLICK_LANDING_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "click_landing_page"
	CLICK_WEBSITE_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_website"
	LOAN_CREDIT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "loan_credit"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_login_30d_cost"
	CONVERT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "convert"
	MESSAGE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "message"
	LUBAN_LIVE_GIFT_CNT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "luban_live_gift_cnt"
	PRE_LOAN_CREDIT_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "pre_loan_credit_cost"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_retention_7d_sum_cnt"
	POI_COLLECT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "poi_collect"
	POI_ADDRESS_CLICK_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "poi_address_click"
	DOWNLOAD_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "download"
	MESSAGE_ACTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "message_action"
	CPC_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpc"
	PLAY_75_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_75_feed_break"
	PRE_CONVERT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_cost"
	DOWNLOAD_START_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "download_start"
	ATTRIBUTION_DEEP_CONVERT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_deep_convert"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_4days"
	IN_APP_DETAIL_UV_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "in_app_detail_uv"
	DEEP_CONVERT_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "deep_convert"
	DEEP_CONVERT_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_rate"
	NEXT_DAY_OPEN_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "next_day_open"
	LOAN_COMPLETION_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "loan_completion"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "advanced_creative_counsel_click"
	VALID_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "valid_play"
	PLAY_DURATION_10S_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_10s"
	LOAN_COMPLETION_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_rate"
	IN_APP_ORDER_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "in_app_order"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_next_day_open_cnt"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_gift_amount"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_wechat_pay_30d_roi"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdGetV2OrderField    ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	PLAY_DURATION_3S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_3s_rate"
	CLICK_INSTALL_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_install"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_cost"
	STAT_UNION_LTV_7_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_7"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_deep_convert_cost"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "luban_order_stat_amount"
	WECHAT_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_pay_amount"
	PLAY_DURATION_2S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_2s"
	LUBAN_ORDER_CNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_retention_7d_total_cost"
	CUSTOMER_EFFECTIVE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "customer_effective"
	PAY_COUNT_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "pay_count"
	IES_MUSIC_CLICK_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "ies_music_click"
	APPROVAL_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "approval_count"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_game_pay_7d_count"
	LOTTERY_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "lottery"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_cost"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_cost"
	PAY_AMOUNT_ROI_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "pay_amount_roi"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_active_pay_count"
	PLAY_100_FEED_BREAK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "play_100_feed_break"
	PLAY_DURATION_2S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_2s_rate"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_rate"
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "live_fans_club_join_cnt"
	UNION_ROI_0_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_0"
	ACTIVE_REGISTER_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_cost"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_roi"
	LIVE_COMPONENT_CLICK_COST_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_cost"
	PLAY_DURATION_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "play_duration"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_7days"
	ACTIVE_PAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_rate"
	CONSULT_EFFECTIVE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "consult_effective"
	VALID_PLAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_cost"
	PLAY_75_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_4d_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_roi_1day"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_2days"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_6days"
)

// Ptr returns reference to report_ad_get_v2_order_field value
func (v ReportAdGetV2OrderField) Ptr() *ReportAdGetV2OrderField {
	return &v
}
