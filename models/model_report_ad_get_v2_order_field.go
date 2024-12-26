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
	LIVE_FANS_CLUB_JOIN_CNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "live_fans_club_join_cnt"
	DEEP_CONVERT_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_rate"
	WIFI_PLAY_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "wifi_play"
	UNION_ROI_7_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_7"
	LUBAN_ORDER_CNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_4days"
	TOTAL_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "total_play"
	ADVANCED_CREATIVE_PHONE_CLICK_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_phone_click"
	WECHAT_FIRST_PAY_COST_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_AMOUNT_ReportAdGetV2OrderField ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_amount"
	LUBAN_LIVE_COMMENT_CNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_comment_cnt"
	ATTRIBUTION_GAME_IN_APP_LTV_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_6days"
	ATTRIBUTION_RETENTION_3D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_rate"
	ATTRIBUTION_NEXT_DAY_OPEN_CNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_next_day_open_cnt"
	ATTRIBUTION_RETENTION_5D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_5d_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_roi_1day"
	PLAY_75_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_75_feed_break"
	CONVERT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "convert"
	GAME_ADDICTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_addiction"
	LOAN_COMPLETION_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_rate"
	CUSTOMER_EFFECTIVE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "customer_effective"
	DOWNLOAD_FINISH_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_cost"
	ATTRIBUTION_ACTIVE_PAY_7D_PER_COUNT_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_active_pay_7d_per_count"
	AVERAGE_PLAY_TIME_PER_PLAY_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "average_play_time_per_play"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COUNT_ReportAdGetV2OrderField      ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_count"
	ACTIVE_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "active_pay_amount"
	PHONE_EFFECTIVE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "phone_effective"
	INSTALL_FINISH_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "install_finish"
	IN_APP_CART_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "in_app_cart"
	FORM_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "form"
	REPORT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "report"
	INSTALL_FINISH_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_cost"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COUNT_ReportAdGetV2OrderField  ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_count"
	UNION_ROI_0_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_0"
	ATTRIBUTION_GAME_IN_APP_LTV_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_8days"
	GAME_PAY_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "game_pay_count"
	PLAY_DURATION_2S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_2s"
	ATTRIBUTION_MICRO_GAME_3D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_ltv"
	ATTRIBUTION_GAME_IN_APP_ROI_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_5days"
	PLAY_DURATION_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "play_duration"
	DEEP_CONVERT_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "deep_convert_cost"
	ATTRIBUTION_RETENTION_3D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_3d_cost"
	ATTRIBUTION_WECHAT_LOGIN_30D_COUNT_ReportAdGetV2OrderField          ReportAdGetV2OrderField = "attribution_wechat_login_30d_count"
	ATTRIBUTION_ACTIVE_PAY_7D_COUNT_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "attribution_active_pay_7d_count"
	HOME_VISITED_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "home_visited"
	LUBAN_LIVE_SLIDECART_CLICK_CNT_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_slidecart_click_cnt"
	LIVE_COMPONENT_CLICK_COST_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_cost"
	ATTRIBUTION_MICRO_GAME_7D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_ltv"
	CTR_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "ctr"
	MAP_SEARCH_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "map_search"
	AVG_CLICK_COST_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "avg_click_cost"
	DOWNLOAD_FINISH_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "download_finish"
	ATTRIBUTION_DAY_ACTIVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_active_pay_count"
	ATTRIBUTION_WECHAT_PAY_30D_ROI_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_wechat_pay_30d_roi"
	GAME_ADDICTION_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_cost"
	ATTRIBUTION_GAME_PAY_7D_COST_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_game_pay_7d_cost"
	ATTRIBUTION_WECHAT_PAY_30D_AMOUNT_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_pay_30d_amount"
	PLAY_DURATION_SUM_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_sum"
	PLAY_DURATION_10S_RATE_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "play_duration_10s_rate"
	PLAY_DURATION_3S_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "play_duration_3s"
	ATTRIBUTION_ACTIVE_PAY_7D_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_cost"
	DOWNLOAD_START_COST_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_cost"
	PRE_CONVERT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_cost"
	PLAY_75_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_75_feed_break_rate"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_RATE_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_rate"
	IN_APP_DETAIL_UV_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "in_app_detail_uv"
	ATTRIBUTION_RETENTION_6D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_6d_cnt"
	ATTRIBUTION_GAME_PAY_7D_COUNT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_game_pay_7d_count"
	ATTRIBUTION_GAME_IN_APP_LTV_1DAY_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_game_in_app_ltv_1day"
	ATTRIBUTION_WECHAT_FIRST_PAY_30D_COST_ReportAdGetV2OrderField       ReportAdGetV2OrderField = "attribution_wechat_first_pay_30d_cost"
	ACTIVE_RATE_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_rate"
	CPA_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpa"
	ATTRIBUTION_MICRO_GAME_7D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_7d_roi"
	SUBMIT_CERTIFICATION_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "submit_certification_count"
	ATTRIBUTION_GAME_IN_APP_ROI_6DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_6days"
	SHOPPING_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "shopping"
	WIFI_PLAY_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "wifi_play_rate"
	INSTALL_FINISH_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "install_finish_rate"
	BUTTON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "button"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_RATE_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_rate"
	LIKE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "like"
	PLAY_25_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_25_feed_break_rate"
	STAT_UNION_LTV_0_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_0"
	INTERACT_PER_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "interact_per_cost"
	LUBAN_LIVE_CLICK_PRODUCT_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "luban_live_click_product_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_7days"
	ACTIVE_REGISTER_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_cost"
	ACTIVE_COST_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "active_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_5DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_5days"
	STAT_UNION_LTV_7_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_7"
	WECHAT_PAY_AMOUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_pay_amount"
	CLICK_WEBSITE_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_website"
	VALID_PLAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_cost"
	ATTRIBUTION_GAME_IN_APP_ROI_4DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_4days"
	PLAY_DURATION_5S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_5s_rate"
	APPROVAL_COUNT_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "approval_count"
	LOTTERY_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "lottery"
	PLAY_DURATION_2S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_2s_rate"
	PHONE_CONNECT_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_connect"
	WECHAT_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "wechat"
	ATTRIBUTION_RETENTION_4D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_cost"
	NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_rate"
	LUBAN_LIVE_GIFT_CNT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "luban_live_gift_cnt"
	ATTRIBUTION_GAME_IN_APP_ROI_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_3days"
	DOWNLOAD_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "download"
	ATTRIBUTION_MICRO_GAME_0D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_roi"
	ATTRIBUTION_WECHAT_LOGIN_30D_COST_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_wechat_login_30d_cost"
	COST_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "cost"
	ATTRIBUTION_RETENTION_7D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_7d_cnt"
	CPC_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpc"
	CLICK_LANDING_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "click_landing_page"
	VALID_PLAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "valid_play_rate"
	PAY_AMOUNT_ROI_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "pay_amount_roi"
	MESSAGE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "message"
	IES_MUSIC_CLICK_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "ies_music_click"
	LUBAN_LIVE_ENTER_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_enter_cnt"
	AVG_RANK_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "avg_rank"
	ATTRIBUTION_CONVERT_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "attribution_convert"
	QQ_ReportAdGetV2OrderField                                          ReportAdGetV2OrderField = "qq"
	LOAN_CREDIT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_rate"
	FOLLOW_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "follow"
	LUBAN_ORDER_ROI_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "luban_order_roi"
	ATTRIBUTION_RETENTION_7D_SUM_CNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_retention_7d_sum_cnt"
	ACTIVE_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "active"
	ATTRIBUTION_RETENTION_6D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_rate"
	COUPON_ReportAdGetV2OrderField                                      ReportAdGetV2OrderField = "coupon"
	SHARE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "share"
	ATTRIBUTION_NEXT_DAY_OPEN_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_rate"
	PLAY_DURATION_3S_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "play_duration_3s_rate"
	PRE_CONVERT_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "pre_convert_count"
	LOAN_COMPLETION_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "loan_completion_cost"
	STAT_UNION_LTV_3_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "stat_union_ltv_3"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_ROI_ReportAdGetV2OrderField    ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_roi"
	PRE_LOAN_CREDIT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "pre_loan_credit"
	DOWNLOAD_START_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "download_start_rate"
	CONSULT_EFFECTIVE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "consult_effective"
	ATTRIBUTION_RETENTION_7D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_cost"
	CPM_ReportAdGetV2OrderField                                         ReportAdGetV2OrderField = "cpm"
	ATTRIBUTION_MICRO_GAME_3D_ROI_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_3d_roi"
	CLICK_INSTALL_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_install"
	PLAY_25_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_25_feed_break"
	LIVE_COMPONENT_CLICK_RATE_ReportAdGetV2OrderField                   ReportAdGetV2OrderField = "live_component_click_rate"
	VALID_PLAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "valid_play"
	PHONE_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "phone"
	ADVANCED_CREATIVE_COUPON_ADDITION_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "advanced_creative_coupon_addition"
	LUBAN_LIVE_FOLLOW_CNT_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "luban_live_follow_cnt"
	ATTRIBUTION_RETENTION_2D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_cost"
	ATTRIBUTION_GAME_IN_APP_LTV_7DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_7days"
	STAT_PAY_AMOUNT_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "stat_pay_amount"
	ATTRIBUTION_RETENTION_5D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_cost"
	ATTRIBUTION_RETENTION_2D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_2d_rate"
	ATTRIBUTION_ACTIVE_PAY_7D_RATE_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_active_pay_7d_rate"
	LUBAN_LIVE_SHARE_CNT_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "luban_live_share_cnt"
	LUBAN_LIVE_PAY_ORDER_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "luban_live_pay_order_count"
	REGISTER_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "register"
	PRE_LOAN_CREDIT_COST_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "pre_loan_credit_cost"
	CONVERT_RATE_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_rate"
	PLAY_50_FEED_BREAK_RATE_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "play_50_feed_break_rate"
	COMMENT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "comment"
	UNION_ROI_3_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "union_roi_3"
	REDIRECT_ReportAdGetV2OrderField                                    ReportAdGetV2OrderField = "redirect"
	ACTIVE_REGISTER_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "active_register_rate"
	IN_APP_ORDER_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "in_app_order"
	LUBAN_LIVE_PAY_ORDER_STAT_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "luban_live_pay_order_stat_cost"
	WECHAT_LOGIN_COUNT_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "wechat_login_count"
	DOWNLOAD_FINISH_RATE_ReportAdGetV2OrderField                        ReportAdGetV2OrderField = "download_finish_rate"
	WECHAT_FIRST_PAY_COUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "wechat_first_pay_count"
	WECHAT_FIRST_PAY_RATE_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "wechat_first_pay_rate"
	COUPON_SINGLE_PAGE_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "coupon_single_page"
	ADVANCED_CREATIVE_FORM_SUBMIT_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "advanced_creative_form_submit"
	VIEW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "view"
	NEXT_DAY_OPEN_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "next_day_open"
	LOAN_CREDIT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "loan_credit"
	DEEP_CONVERT_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "deep_convert"
	VOTE_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "vote"
	GAME_ADDICTION_RATE_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "game_addiction_rate"
	SHOW_ReportAdGetV2OrderField                                        ReportAdGetV2OrderField = "show"
	ATTRIBUTION_RETENTION_6D_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_6d_cost"
	FIRST_RENTAL_ORDER_COUNT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "first_rental_order_count"
	LIVE_COMPONENT_CLICK_COUNT_ReportAdGetV2OrderField                  ReportAdGetV2OrderField = "live_component_click_count"
	IN_APP_PAY_ReportAdGetV2OrderField                                  ReportAdGetV2OrderField = "in_app_pay"
	ATTRIBUTION_ACTIVE_PAY_INTRA_ONE_DAY_COST_ReportAdGetV2OrderField   ReportAdGetV2OrderField = "attribution_active_pay_intra_one_day_cost"
	ATTRIBUTION_RETENTION_7D_TOTAL_COST_ReportAdGetV2OrderField         ReportAdGetV2OrderField = "attribution_retention_7d_total_cost"
	ATTRIBUTION_DAY_ACITVE_PAY_COUNT_ReportAdGetV2OrderField            ReportAdGetV2OrderField = "attribution_day_acitve_pay_count"
	PLAY_DURATION_10S_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "play_duration_10s"
	CLICK_CALL_DY_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "click_call_dy"
	CLICK_ReportAdGetV2OrderField                                       ReportAdGetV2OrderField = "click"
	COMMUTE_FIRST_PAY_COUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "commute_first_pay_count"
	PRE_CONVERT_RATE_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "pre_convert_rate"
	PHONE_CONFIRM_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "phone_confirm"
	CLICK_SHOPWINDOW_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "click_shopwindow"
	LUBAN_LIVE_GIFT_AMOUNT_ReportAdGetV2OrderField                      ReportAdGetV2OrderField = "luban_live_gift_amount"
	GAME_PAY_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "game_pay_cost"
	ATTRIBUTION_RETENTION_4D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_4d_cnt"
	CONSULT_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "consult"
	ATTRIBUTION_GAME_IN_APP_ROI_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_2days"
	AVERAGE_VIDEO_PLAY_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "average_video_play"
	ADVANCED_CREATIVE_FORM_CLICK_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "advanced_creative_form_click"
	LOAN_COMPLETION_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "loan_completion"
	CLICK_DOWNLOAD_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "click_download"
	CONVERT_SHOW_RATE_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "convert_show_rate"
	AVERAGE_PLAY_PROGRESS_ReportAdGetV2OrderField                       ReportAdGetV2OrderField = "average_play_progress"
	ATTRIBUTION_RETENTION_7D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_7d_rate"
	PLAY_100_FEED_BREAK_RATE_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "play_100_feed_break_rate"
	ATTRIBUTION_GAME_IN_APP_LTV_3DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_3days"
	ATTRIBUTION_DEEP_CONVERT_COST_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_deep_convert_cost"
	ATTRIBUTION_MICRO_GAME_0D_LTV_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_micro_game_0d_ltv"
	WECHAT_LOGIN_COST_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "wechat_login_cost"
	LOAN_CREDIT_COST_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "loan_credit_cost"
	CONVERT_COST_ReportAdGetV2OrderField                                ReportAdGetV2OrderField = "convert_cost"
	ATTRIBUTION_CONVERT_COST_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_convert_cost"
	PLAY_50_FEED_BREAK_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "play_50_feed_break"
	DISLIKE_ReportAdGetV2OrderField                                     ReportAdGetV2OrderField = "dislike"
	NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField                          ReportAdGetV2OrderField = "next_day_open_cost"
	CARD_SHOW_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "card_show"
	REDIRECT_TO_SHOP_ReportAdGetV2OrderField                            ReportAdGetV2OrderField = "redirect_to_shop"
	LUBAN_ORDER_STAT_AMOUNT_ReportAdGetV2OrderField                     ReportAdGetV2OrderField = "luban_order_stat_amount"
	ATTRIBUTION_NEXT_DAY_OPEN_COST_ReportAdGetV2OrderField              ReportAdGetV2OrderField = "attribution_next_day_open_cost"
	ACTIVE_PAY_COST_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_cost"
	ATTRIBUTION_DEEP_CONVERT_ReportAdGetV2OrderField                    ReportAdGetV2OrderField = "attribution_deep_convert"
	FIRST_ORDER_COUNT_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "first_order_count"
	PLAY_OVER_RATE_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "play_over_rate"
	ATTRIBUTION_RETENTION_4D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_4d_rate"
	MESSAGE_ACTION_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "message_action"
	PAY_COUNT_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "pay_count"
	ATTRIBUTION_RETENTION_2D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_2d_cnt"
	PLAY_OVER_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "play_over"
	AVG_SHOW_COST_ReportAdGetV2OrderField                               ReportAdGetV2OrderField = "avg_show_cost"
	ATTRIBUTION_RETENTION_5D_RATE_ReportAdGetV2OrderField               ReportAdGetV2OrderField = "attribution_retention_5d_rate"
	POI_COLLECT_ReportAdGetV2OrderField                                 ReportAdGetV2OrderField = "poi_collect"
	ADVANCED_CREATIVE_COUNSEL_CLICK_ReportAdGetV2OrderField             ReportAdGetV2OrderField = "advanced_creative_counsel_click"
	PLAY_100_FEED_BREAK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "play_100_feed_break"
	LOCATION_CLICK_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "location_click"
	ATTRIBUTION_RETENTION_3D_CNT_ReportAdGetV2OrderField                ReportAdGetV2OrderField = "attribution_retention_3d_cnt"
	ACTIVE_PAY_RATE_ReportAdGetV2OrderField                             ReportAdGetV2OrderField = "active_pay_rate"
	IES_CHALLENGE_CLICK_ReportAdGetV2OrderField                         ReportAdGetV2OrderField = "ies_challenge_click"
	IN_APP_UV_ReportAdGetV2OrderField                                   ReportAdGetV2OrderField = "in_app_uv"
	DOWNLOAD_START_ReportAdGetV2OrderField                              ReportAdGetV2OrderField = "download_start"
	POI_ADDRESS_CLICK_ReportAdGetV2OrderField                           ReportAdGetV2OrderField = "poi_address_click"
	ATTRIBUTION_GAME_IN_APP_LTV_2DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_ltv_2days"
	LIVE_WATCH_ONE_MINUTE_COUNT_ReportAdGetV2OrderField                 ReportAdGetV2OrderField = "live_watch_one_minute_count"
	ATTRIBUTION_GAME_IN_APP_ROI_8DAYS_ReportAdGetV2OrderField           ReportAdGetV2OrderField = "attribution_game_in_app_roi_8days"
)

// Ptr returns reference to report_ad_get_v2_order_field value
func (v ReportAdGetV2OrderField) Ptr() *ReportAdGetV2OrderField {
	return &v
}
