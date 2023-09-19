/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction
type EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction string

// List of event_manager_optimized_goal_get_v2_v3.0_data_goals_external_action
const (
	AD_CONVERT_PAGE_VIEW_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_TYPE_ACTIVE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                     EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_ARPU0_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_AUTHORIZATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLASS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_CLICK_CALL_DY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                     EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_CLASS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_GAME_ADDICTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                     EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                  EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                     EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                     EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                 EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction              EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                  EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SHARE_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                    EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction               EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                   EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction         EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                        EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                      EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_FIRST_MSG_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_FIRST_MSG"
	AD_CONVERT_TYPE_WECHAT_PAY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                  EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction          EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW"
	AD_CONVERT_TYPE_WECHAT_QRCODE_TRY_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction           EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_TRY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction             EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction            EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction                       EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction enum
var AllowedEventManagerOptimizedGoalGetV2V30DataGoalsExternalActionEnumValues = []EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction{
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_TYPE_ACTIVE",
	"AD_CONVERT_TYPE_ACTIVE_REGISTER",
	"AD_CONVERT_TYPE_ANCHOR_CLICK",
	"AD_CONVERT_TYPE_APPLET_CLICK",
	"AD_CONVERT_TYPE_APP_CART",
	"AD_CONVERT_TYPE_APP_DETAIL_UV",
	"AD_CONVERT_TYPE_APP_ORDER",
	"AD_CONVERT_TYPE_APP_PAY",
	"AD_CONVERT_TYPE_APP_UV",
	"AD_CONVERT_TYPE_ARPU0",
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
	"AD_CONVERT_TYPE_CLASS",
	"AD_CONVERT_TYPE_CLICK_CALL_DY",
	"AD_CONVERT_TYPE_CLICK_DOWNLOAD",
	"AD_CONVERT_TYPE_CLICK_LANDING_PAGE",
	"AD_CONVERT_TYPE_CLICK_NUM",
	"AD_CONVERT_TYPE_CLICK_SHOPWINDOW",
	"AD_CONVERT_TYPE_CLICK_WEBSITE",
	"AD_CONVERT_TYPE_CLUE_CONFIRM",
	"AD_CONVERT_TYPE_CLUE_HIGH_INTENTION",
	"AD_CONVERT_TYPE_CLUE_INTERFLOW",
	"AD_CONVERT_TYPE_CLUE_PAY_SUCCEED",
	"AD_CONVERT_TYPE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_COMMODITY_CLICK",
	"AD_CONVERT_TYPE_CONSULT",
	"AD_CONVERT_TYPE_CONSULT_CLUE",
	"AD_CONVERT_TYPE_CONSULT_EFFECTIVE",
	"AD_CONVERT_TYPE_COUPON",
	"AD_CONVERT_TYPE_CREATE_GAMEROLE",
	"AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE",
	"AD_CONVERT_TYPE_DEEP_PURCHASE",
	"AD_CONVERT_TYPE_DIALBACK",
	"AD_CONVERT_TYPE_DIALBACK_CONFIRM",
	"AD_CONVERT_TYPE_DIALBACK_CONNECT",
	"AD_CONVERT_TYPE_DOWNLOAD_FINISH",
	"AD_CONVERT_TYPE_DOWNLOAD_START",
	"AD_CONVERT_TYPE_EFFECTIVE_COPY",
	"AD_CONVERT_TYPE_EFFECTIVE_PLAY",
	"AD_CONVERT_TYPE_ENTER_HOMEPAGE",
	"AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_FIRST_CLASS",
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_GAME_ADDICTION",
	"AD_CONVERT_TYPE_HIGH_VALUE_CLUE",
	"AD_CONVERT_TYPE_IDCARD_INFORMATION",
	"AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_INSTALL_FINISH",
	"AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN",
	"AD_CONVERT_TYPE_INTENTION_CLUE",
	"AD_CONVERT_TYPE_INTERACTION",
	"AD_CONVERT_TYPE_INVALID_CLUE",
	"AD_CONVERT_TYPE_IPU_QUALIFY",
	"AD_CONVERT_TYPE_LIKE_ACTION",
	"AD_CONVERT_TYPE_LINK_ACTION",
	"AD_CONVERT_TYPE_LIVE_APPOINTMENT",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_FANS_ACTION",
	"AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON",
	"AD_CONVERT_TYPE_LIVE_GIFT_ACTION",
	"AD_CONVERT_TYPE_LIVE_HOMEPAGE",
	"AD_CONVERT_TYPE_LIVE_JOIN_GROUP",
	"AD_CONVERT_TYPE_LIVE_NATIVE_ACITON",
	"AD_CONVERT_TYPE_LIVE_OTO_CLICK",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY",
	"AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK",
	"AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION",
	"AD_CONVERT_TYPE_LIVE_STAY_TIME",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LOAN",
	"AD_CONVERT_TYPE_LOAN_COMPLETION",
	"AD_CONVERT_TYPE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_LOCATION_ACTION",
	"AD_CONVERT_TYPE_LOTTERY",
	"AD_CONVERT_TYPE_LT_ROI",
	"AD_CONVERT_TYPE_MAP_SEARCH",
	"AD_CONVERT_TYPE_MESSAGE",
	"AD_CONVERT_TYPE_MESSAGE_ACTION",
	"AD_CONVERT_TYPE_MESSAGE_CLICK",
	"AD_CONVERT_TYPE_MESSAGE_CLUE",
	"AD_CONVERT_TYPE_MESSAGE_INTERACTION",
	"AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP",
	"AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS",
	"AD_CONVERT_TYPE_MESSAGE_SERVICE",
	"AD_CONVERT_TYPE_MULTIPLE",
	"AD_CONVERT_TYPE_NATIVE_ACTION",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_NEXT_DAY_OPEN",
	"AD_CONVERT_TYPE_NOTIFY_DOWNLOAD",
	"AD_CONVERT_TYPE_NOTIFY_FORM",
	"AD_CONVERT_TYPE_OTHER",
	"AD_CONVERT_TYPE_OTO_PAY",
	"AD_CONVERT_TYPE_OTO_STAY_TIME",
	"AD_CONVERT_TYPE_PAID_CLUE",
	"AD_CONVERT_TYPE_PAY",
	"AD_CONVERT_TYPE_PERSONAL_INFORMATION",
	"AD_CONVERT_TYPE_PHONE",
	"AD_CONVERT_TYPE_PHONE_CONFIRM",
	"AD_CONVERT_TYPE_PHONE_CONNECT",
	"AD_CONVERT_TYPE_PHONE_EFFECTIVE",
	"AD_CONVERT_TYPE_POI_ADDRESS_CLICK",
	"AD_CONVERT_TYPE_POI_COLLECT",
	"AD_CONVERT_TYPE_POI_MULTIPLE",
	"AD_CONVERT_TYPE_PREMIUM_PAYMENT",
	"AD_CONVERT_TYPE_PREMIUM_ROI",
	"AD_CONVERT_TYPE_PREMIUM_UPGEADE",
	"AD_CONVERT_TYPE_PRE_LOAN_CREDIT",
	"AD_CONVERT_TYPE_PURCHASE_OF_GOODS",
	"AD_CONVERT_TYPE_PURCHASE_ROI",
	"AD_CONVERT_TYPE_PURCHASE_ROI_2D",
	"AD_CONVERT_TYPE_PURCHASE_ROI_7D",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_QQ",
	"AD_CONVERT_TYPE_REDIRECT",
	"AD_CONVERT_TYPE_REDIRECT_TO_SHOP",
	"AD_CONVERT_TYPE_REDIRECT_TO_STORE",
	"AD_CONVERT_TYPE_RESERVATION",
	"AD_CONVERT_TYPE_RETENTION_7D",
	"AD_CONVERT_TYPE_RETENTION_DAYS",
	"AD_CONVERT_TYPE_RSS",
	"AD_CONVERT_TYPE_SALES_LEAD",
	"AD_CONVERT_TYPE_SHARE_ACTION",
	"AD_CONVERT_TYPE_SHOPPING",
	"AD_CONVERT_TYPE_SHOPPING_ACTION",
	"AD_CONVERT_TYPE_SHOPPING_CART",
	"AD_CONVERT_TYPE_SHOW_OFF_NUM",
	"AD_CONVERT_TYPE_STAY_TIME",
	"AD_CONVERT_TYPE_SUBMIT_CERTIFICATION",
	"AD_CONVERT_TYPE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_UG_ROI",
	"AD_CONVERT_TYPE_UPDATE_LEVEL",
	"AD_CONVERT_TYPE_VIEW",
	"AD_CONVERT_TYPE_VOTE",
	"AD_CONVERT_TYPE_WECHAT",
	"AD_CONVERT_TYPE_WECHAT_FIRST_MSG",
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW",
	"AD_CONVERT_TYPE_WECHAT_QRCODE_TRY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_XPATH",
}

// NewEventManagerOptimizedGoalGetV2V30DataGoalsExternalActionFromValue returns a pointer to a valid EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventManagerOptimizedGoalGetV2V30DataGoalsExternalActionFromValue(v string) (*EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction, error) {
	ev := EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction: valid values are %v", v, AllowedEventManagerOptimizedGoalGetV2V30DataGoalsExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction) IsValid() bool {
	for _, existing := range AllowedEventManagerOptimizedGoalGetV2V30DataGoalsExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to event_manager_optimized_goal_get_v2_v3.0_data_goals_external_action value
func (v EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction) Ptr() *EventManagerOptimizedGoalGetV2V30DataGoalsExternalAction {
	return &v
}
