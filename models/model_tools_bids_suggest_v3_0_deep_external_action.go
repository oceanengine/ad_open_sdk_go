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

// ToolsBidsSuggestV30DeepExternalAction
type ToolsBidsSuggestV30DeepExternalAction string

// List of tools_bids_suggest_v3.0_deep_external_action
const (
	AD_CONVERT_PAGE_VIEW_ToolsBidsSuggestV30DeepExternalAction                        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_PHONE_CONNECT_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_PHONE_CONNECT"
	AD_CONVERT_TYPE_ACTIVE_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_ToolsBidsSuggestV30DeepExternalAction                   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_ToolsBidsSuggestV30DeepExternalAction                     ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_ARPU0_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_ARPU0"
	AD_CONVERT_TYPE_AUTHORIZATION_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_ToolsBidsSuggestV30DeepExternalAction        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_ToolsBidsSuggestV30DeepExternalAction   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLASS_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLASS"
	AD_CONVERT_TYPE_CLICK_CALL_DY_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_ToolsBidsSuggestV30DeepExternalAction                   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_ToolsBidsSuggestV30DeepExternalAction         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_ToolsBidsSuggestV30DeepExternalAction                     ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_CLASS_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FIRST_CLASS"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_ToolsBidsSuggestV30DeepExternalAction        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_ToolsBidsSuggestV30DeepExternalAction    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_ToolsBidsSuggestV30DeepExternalAction                        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_ToolsBidsSuggestV30DeepExternalAction                   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_GAME_ADDICTION_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_ToolsBidsSuggestV30DeepExternalAction         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_ToolsBidsSuggestV30DeepExternalAction    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_ToolsBidsSuggestV30DeepExternalAction   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_ToolsBidsSuggestV30DeepExternalAction         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_ToolsBidsSuggestV30DeepExternalAction        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_ToolsBidsSuggestV30DeepExternalAction ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_ToolsBidsSuggestV30DeepExternalAction    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_ToolsBidsSuggestV30DeepExternalAction       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LOAN_ToolsBidsSuggestV30DeepExternalAction                        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LOAN"
	AD_CONVERT_TYPE_LOAN_COMPLETION_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_ToolsBidsSuggestV30DeepExternalAction                     ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_ToolsBidsSuggestV30DeepExternalAction                  ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_ToolsBidsSuggestV30DeepExternalAction                     ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_ToolsBidsSuggestV30DeepExternalAction         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_ToolsBidsSuggestV30DeepExternalAction       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_ToolsBidsSuggestV30DeepExternalAction                     ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_ToolsBidsSuggestV30DeepExternalAction                   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_ToolsBidsSuggestV30DeepExternalAction                         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_ToolsBidsSuggestV30DeepExternalAction        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PREMIUM_UPGEADE_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PREMIUM_UPGEADE"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_ToolsBidsSuggestV30DeepExternalAction                          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_ToolsBidsSuggestV30DeepExternalAction                 ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_ToolsBidsSuggestV30DeepExternalAction              ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_ToolsBidsSuggestV30DeepExternalAction                         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_ToolsBidsSuggestV30DeepExternalAction                  ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SHARE_ACTION_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_ToolsBidsSuggestV30DeepExternalAction                    ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_ToolsBidsSuggestV30DeepExternalAction               ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_ToolsBidsSuggestV30DeepExternalAction                   ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_ToolsBidsSuggestV30DeepExternalAction        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_ToolsBidsSuggestV30DeepExternalAction         ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_ToolsBidsSuggestV30DeepExternalAction                ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_ToolsBidsSuggestV30DeepExternalAction                        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_ToolsBidsSuggestV30DeepExternalAction                        ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_ToolsBidsSuggestV30DeepExternalAction                      ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_FIRST_MSG_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_FIRST_MSG"
	AD_CONVERT_TYPE_WECHAT_PAY_ToolsBidsSuggestV30DeepExternalAction                  ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW_ToolsBidsSuggestV30DeepExternalAction          ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_SHOW"
	AD_CONVERT_TYPE_WECHAT_QRCODE_TRY_ToolsBidsSuggestV30DeepExternalAction           ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_QRCODE_TRY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_ToolsBidsSuggestV30DeepExternalAction             ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_ToolsBidsSuggestV30DeepExternalAction            ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_ToolsBidsSuggestV30DeepExternalAction                       ToolsBidsSuggestV30DeepExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of ToolsBidsSuggestV30DeepExternalAction enum
var AllowedToolsBidsSuggestV30DeepExternalActionEnumValues = []ToolsBidsSuggestV30DeepExternalAction{
	"AD_CONVERT_PAGE_VIEW",
	"AD_CONVERT_PHONE_CONNECT",
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

// NewToolsBidsSuggestV30DeepExternalActionFromValue returns a pointer to a valid ToolsBidsSuggestV30DeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsBidsSuggestV30DeepExternalActionFromValue(v string) (*ToolsBidsSuggestV30DeepExternalAction, error) {
	ev := ToolsBidsSuggestV30DeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsBidsSuggestV30DeepExternalAction: valid values are %v", v, AllowedToolsBidsSuggestV30DeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsBidsSuggestV30DeepExternalAction) IsValid() bool {
	for _, existing := range AllowedToolsBidsSuggestV30DeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_bids_suggest_v3.0_deep_external_action value
func (v ToolsBidsSuggestV30DeepExternalAction) Ptr() *ToolsBidsSuggestV30DeepExternalAction {
	return &v
}
