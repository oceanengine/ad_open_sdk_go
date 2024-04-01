/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoExternalAction
type AdlabGroupCreateV30AdInfoExternalAction string

// List of adlab_group_create_v3.0_ad_info_external_action
const (
	AD_CONVERT_PAGE_VIEW_AdlabGroupCreateV30AdInfoExternalAction                         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_PAGE_VIEW"
	AD_CONVERT_TYPE_ACTIVE_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_ACTIVE"
	AD_CONVERT_TYPE_ACTIVE_REGISTER_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	AD_CONVERT_TYPE_ANCHOR_CLICK_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_ANCHOR_CLICK"
	AD_CONVERT_TYPE_APPLET_CLICK_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APPLET_CLICK"
	AD_CONVERT_TYPE_APP_CART_AdlabGroupCreateV30AdInfoExternalAction                     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APP_CART"
	AD_CONVERT_TYPE_APP_DETAIL_UV_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APP_DETAIL_UV"
	AD_CONVERT_TYPE_APP_ORDER_AdlabGroupCreateV30AdInfoExternalAction                    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APP_ORDER"
	AD_CONVERT_TYPE_APP_PAY_AdlabGroupCreateV30AdInfoExternalAction                      AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APP_PAY"
	AD_CONVERT_TYPE_APP_UV_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_APP_UV"
	AD_CONVERT_TYPE_AUTHORIZATION_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_AUTHORIZATION"
	AD_CONVERT_TYPE_BANKCARD_INFORMATION_AdlabGroupCreateV30AdInfoExternalAction         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_BANKCARD_INFORMATION"
	AD_CONVERT_TYPE_BOOST_AdlabGroupCreateV30AdInfoExternalAction                        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_BOOST"
	AD_CONVERT_TYPE_BUTTON_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_BUTTON"
	AD_CONVERT_TYPE_CERTIFICATION_INFORMATION_AdlabGroupCreateV30AdInfoExternalAction    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CERTIFICATION_INFORMATION"
	AD_CONVERT_TYPE_CLICK_CALL_DY_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_CALL_DY"
	AD_CONVERT_TYPE_CLICK_DOWNLOAD_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_DOWNLOAD"
	AD_CONVERT_TYPE_CLICK_LANDING_PAGE_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_LANDING_PAGE"
	AD_CONVERT_TYPE_CLICK_NUM_AdlabGroupCreateV30AdInfoExternalAction                    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_NUM"
	AD_CONVERT_TYPE_CLICK_SHOPWINDOW_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_SHOPWINDOW"
	AD_CONVERT_TYPE_CLICK_WEBSITE_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLICK_WEBSITE"
	AD_CONVERT_TYPE_CLUE_CONFIRM_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLUE_CONFIRM"
	AD_CONVERT_TYPE_CLUE_HIGH_INTENTION_AdlabGroupCreateV30AdInfoExternalAction          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLUE_HIGH_INTENTION"
	AD_CONVERT_TYPE_CLUE_INTERFLOW_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLUE_INTERFLOW"
	AD_CONVERT_TYPE_CLUE_PAY_SUCCEED_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CLUE_PAY_SUCCEED"
	AD_CONVERT_TYPE_COMMENT_ACTION_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_COMMENT_ACTION"
	AD_CONVERT_TYPE_COMMODITY_CLICK_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_COMMODITY_CLICK"
	AD_CONVERT_TYPE_CONSULT_AdlabGroupCreateV30AdInfoExternalAction                      AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CONSULT"
	AD_CONVERT_TYPE_CONSULT_CLUE_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CONSULT_CLUE"
	AD_CONVERT_TYPE_CONSULT_EFFECTIVE_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CONSULT_EFFECTIVE"
	AD_CONVERT_TYPE_COUPON_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_COUPON"
	AD_CONVERT_TYPE_CREATE_GAMEROLE_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CREATE_GAMEROLE"
	AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_CUSTOMER_EFFECTIVE"
	AD_CONVERT_TYPE_DEEP_PURCHASE_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DEEP_PURCHASE"
	AD_CONVERT_TYPE_DIALBACK_AdlabGroupCreateV30AdInfoExternalAction                     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DIALBACK"
	AD_CONVERT_TYPE_DIALBACK_CONFIRM_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONFIRM"
	AD_CONVERT_TYPE_DIALBACK_CONNECT_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DIALBACK_CONNECT"
	AD_CONVERT_TYPE_DOWNLOAD_FINISH_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_FINISH"
	AD_CONVERT_TYPE_DOWNLOAD_START_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_DOWNLOAD_START"
	AD_CONVERT_TYPE_EFFECTIVE_COPY_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_COPY"
	AD_CONVERT_TYPE_EFFECTIVE_PLAY_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_EFFECTIVE_PLAY"
	AD_CONVERT_TYPE_ENTER_HOMEPAGE_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_ENTER_HOMEPAGE"
	AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FEED_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_FIRST_RENTAL_ORDER_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FIRST_RENTAL_ORDER"
	AD_CONVERT_TYPE_FOLLOW_ACTION_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FOLLOW_ACTION"
	AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT_AdlabGroupCreateV30AdInfoExternalAction         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT"
	AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER"
	AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH_AdlabGroupCreateV30AdInfoExternalAction     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH"
	AD_CONVERT_TYPE_FORM_AdlabGroupCreateV30AdInfoExternalAction                         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FORM"
	AD_CONVERT_TYPE_FORM_ANSWER_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FORM_ANSWER"
	AD_CONVERT_TYPE_FORM_CONNECT_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FORM_CONNECT"
	AD_CONVERT_TYPE_FORM_DEEP_AdlabGroupCreateV30AdInfoExternalAction                    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_FORM_DEEP"
	AD_CONVERT_TYPE_GAMESTATION_DOWNLOAD_DOUPLUS_AdlabGroupCreateV30AdInfoExternalAction AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_GAMESTATION_DOWNLOAD_DOUPLUS"
	AD_CONVERT_TYPE_GAME_ADDICTION_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_GAME_ADDICTION"
	AD_CONVERT_TYPE_HIGH_VALUE_CLUE_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_HIGH_VALUE_CLUE"
	AD_CONVERT_TYPE_IDCARD_INFORMATION_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_IDCARD_INFORMATION"
	AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN_AdlabGroupCreateV30AdInfoExternalAction          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INAPP_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_INSTALL_FINISH_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INSTALL_FINISH"
	AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN_AdlabGroupCreateV30AdInfoExternalAction     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INTELLIGENT_SERVICE_OPEN"
	AD_CONVERT_TYPE_INTENTION_CLUE_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INTENTION_CLUE"
	AD_CONVERT_TYPE_INTERACTION_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INTERACTION"
	AD_CONVERT_TYPE_INVALID_CLUE_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_INVALID_CLUE"
	AD_CONVERT_TYPE_IPU_QUALIFY_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_IPU_QUALIFY"
	AD_CONVERT_TYPE_LIKE_ACTION_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIKE_ACTION"
	AD_CONVERT_TYPE_LINK_ACTION_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LINK_ACTION"
	AD_CONVERT_TYPE_LIVE_APPOINTMENT_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_APPOINTMENT"
	AD_CONVERT_TYPE_LIVE_BUSINESS_FITTING_AdlabGroupCreateV30AdInfoExternalAction        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_BUSINESS_FITTING"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_AdlabGroupCreateV30AdInfoExternalAction    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_AdlabGroupCreateV30AdInfoExternalAction          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK_AdlabGroupCreateV30AdInfoExternalAction         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_COMPONENT_CLICK"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_FANS_ACTION_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_FANS_ACTION"
	AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_FOLLOW_ACITON"
	AD_CONVERT_TYPE_LIVE_GIFT_ACTION_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_GIFT_ACTION"
	AD_CONVERT_TYPE_LIVE_HOMEPAGE_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_HOMEPAGE"
	AD_CONVERT_TYPE_LIVE_JOIN_GROUP_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_JOIN_GROUP"
	AD_CONVERT_TYPE_LIVE_NATIVE_ACITON_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_NATIVE_ACITON"
	AD_CONVERT_TYPE_LIVE_OTO_CLICK_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_CLICK"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY"
	AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_OTO_PAY_CLICK"
	AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION_AdlabGroupCreateV30AdInfoExternalAction  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_SLIDECART_CLICK_ACTION"
	AD_CONVERT_TYPE_LIVE_STAY_TIME_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_STAY_TIME"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_AdlabGroupCreateV30AdInfoExternalAction     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LOAN_COMPLETION_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LOAN_COMPLETION"
	AD_CONVERT_TYPE_LOAN_CREDIT_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LOAN_CREDIT"
	AD_CONVERT_TYPE_LOCATION_ACTION_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LOCATION_ACTION"
	AD_CONVERT_TYPE_LOTTERY_AdlabGroupCreateV30AdInfoExternalAction                      AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LOTTERY"
	AD_CONVERT_TYPE_LT_ROI_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_LT_ROI"
	AD_CONVERT_TYPE_MAP_SEARCH_AdlabGroupCreateV30AdInfoExternalAction                   AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MAP_SEARCH"
	AD_CONVERT_TYPE_MESSAGE_AdlabGroupCreateV30AdInfoExternalAction                      AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE"
	AD_CONVERT_TYPE_MESSAGE_ACTION_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_ACTION"
	AD_CONVERT_TYPE_MESSAGE_CLICK_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLICK"
	AD_CONVERT_TYPE_MESSAGE_CLUE_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_CLUE"
	AD_CONVERT_TYPE_MESSAGE_INTERACTION_AdlabGroupCreateV30AdInfoExternalAction          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_INTERACTION"
	AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP_AdlabGroupCreateV30AdInfoExternalAction           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_JOIN_GROUP"
	AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS_AdlabGroupCreateV30AdInfoExternalAction        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_ORDER_SUCCESS"
	AD_CONVERT_TYPE_MESSAGE_SERVICE_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MESSAGE_SERVICE"
	AD_CONVERT_TYPE_MULTIPLE_AdlabGroupCreateV30AdInfoExternalAction                     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_MULTIPLE"
	AD_CONVERT_TYPE_NATIVE_ACTION_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_NATIVE_ACTION"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_NEXT_DAY_OPEN_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_NEXT_DAY_OPEN"
	AD_CONVERT_TYPE_NOTIFY_DOWNLOAD_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_NOTIFY_DOWNLOAD"
	AD_CONVERT_TYPE_NOTIFY_FORM_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_NOTIFY_FORM"
	AD_CONVERT_TYPE_OTHER_AdlabGroupCreateV30AdInfoExternalAction                        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_OTHER"
	AD_CONVERT_TYPE_OTO_PAY_AdlabGroupCreateV30AdInfoExternalAction                      AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_OTO_PAY"
	AD_CONVERT_TYPE_OTO_STAY_TIME_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_OTO_STAY_TIME"
	AD_CONVERT_TYPE_PAID_CLUE_AdlabGroupCreateV30AdInfoExternalAction                    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PAID_CLUE"
	AD_CONVERT_TYPE_PAY_AdlabGroupCreateV30AdInfoExternalAction                          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PAY"
	AD_CONVERT_TYPE_PERSONAL_INFORMATION_AdlabGroupCreateV30AdInfoExternalAction         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PERSONAL_INFORMATION"
	AD_CONVERT_TYPE_PHONE_AdlabGroupCreateV30AdInfoExternalAction                        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PHONE"
	AD_CONVERT_TYPE_PHONE_CONFIRM_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PHONE_CONFIRM"
	AD_CONVERT_TYPE_PHONE_CONNECT_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PHONE_CONNECT"
	AD_CONVERT_TYPE_PHONE_EFFECTIVE_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PHONE_EFFECTIVE"
	AD_CONVERT_TYPE_POI_ADDRESS_CLICK_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_POI_ADDRESS_CLICK"
	AD_CONVERT_TYPE_POI_COLLECT_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_POI_COLLECT"
	AD_CONVERT_TYPE_POI_MULTIPLE_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_POI_MULTIPLE"
	AD_CONVERT_TYPE_PREMIUM_PAYMENT_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PREMIUM_PAYMENT"
	AD_CONVERT_TYPE_PREMIUM_ROI_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PREMIUM_ROI"
	AD_CONVERT_TYPE_PRE_LOAN_CREDIT_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PRE_LOAN_CREDIT"
	AD_CONVERT_TYPE_PURCHASE_OF_GOODS_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PURCHASE_OF_GOODS"
	AD_CONVERT_TYPE_PURCHASE_ROI_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI"
	AD_CONVERT_TYPE_PURCHASE_ROI_2_D_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_2D"
	AD_CONVERT_TYPE_PURCHASE_ROI_7_D_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_PURCHASE_ROI_7D"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_QQ_AdlabGroupCreateV30AdInfoExternalAction                           AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_QQ"
	AD_CONVERT_TYPE_REDIRECT_AdlabGroupCreateV30AdInfoExternalAction                     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_REDIRECT"
	AD_CONVERT_TYPE_REDIRECT_TO_SHOP_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_SHOP"
	AD_CONVERT_TYPE_REDIRECT_TO_STORE_AdlabGroupCreateV30AdInfoExternalAction            AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_REDIRECT_TO_STORE"
	AD_CONVERT_TYPE_RESERVATION_AdlabGroupCreateV30AdInfoExternalAction                  AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_RESERVATION"
	AD_CONVERT_TYPE_RETENTION_7_D_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_RETENTION_7D"
	AD_CONVERT_TYPE_RETENTION_DAYS_AdlabGroupCreateV30AdInfoExternalAction               AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_RETENTION_DAYS"
	AD_CONVERT_TYPE_RSS_AdlabGroupCreateV30AdInfoExternalAction                          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_RSS"
	AD_CONVERT_TYPE_SALES_LEAD_AdlabGroupCreateV30AdInfoExternalAction                   AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SALES_LEAD"
	AD_CONVERT_TYPE_SHARE_ACTION_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SHARE_ACTION"
	AD_CONVERT_TYPE_SHOPPING_AdlabGroupCreateV30AdInfoExternalAction                     AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SHOPPING"
	AD_CONVERT_TYPE_SHOPPING_ACTION_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SHOPPING_ACTION"
	AD_CONVERT_TYPE_SHOPPING_CART_AdlabGroupCreateV30AdInfoExternalAction                AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SHOPPING_CART"
	AD_CONVERT_TYPE_SHOW_OFF_NUM_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SHOW_OFF_NUM"
	AD_CONVERT_TYPE_STAY_TIME_AdlabGroupCreateV30AdInfoExternalAction                    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_STAY_TIME"
	AD_CONVERT_TYPE_SUBMIT_CERTIFICATION_AdlabGroupCreateV30AdInfoExternalAction         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SUBMIT_CERTIFICATION"
	AD_CONVERT_TYPE_SUCCESSORDER_ACTION_AdlabGroupCreateV30AdInfoExternalAction          AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_UG_ROI_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_UG_ROI"
	AD_CONVERT_TYPE_UPDATE_LEVEL_AdlabGroupCreateV30AdInfoExternalAction                 AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_UPDATE_LEVEL"
	AD_CONVERT_TYPE_VIEW_AdlabGroupCreateV30AdInfoExternalAction                         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_VIEW"
	AD_CONVERT_TYPE_VOTE_AdlabGroupCreateV30AdInfoExternalAction                         AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_VOTE"
	AD_CONVERT_TYPE_WECHAT_AdlabGroupCreateV30AdInfoExternalAction                       AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_WECHAT"
	AD_CONVERT_TYPE_WECHAT_PAY_AdlabGroupCreateV30AdInfoExternalAction                   AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_WECHAT_PAY"
	AD_CONVERT_TYPE_WECHAT_REGISTER_AdlabGroupCreateV30AdInfoExternalAction              AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_WECHAT_REGISTER"
	AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE_AdlabGroupCreateV30AdInfoExternalAction    AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE"
	AD_CONVERT_TYPE_WECHAT_WECOM_ADD_AdlabGroupCreateV30AdInfoExternalAction             AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_WECHAT_WECOM_ADD"
	AD_CONVERT_TYPE_XPATH_AdlabGroupCreateV30AdInfoExternalAction                        AdlabGroupCreateV30AdInfoExternalAction = "AD_CONVERT_TYPE_XPATH"
)

// All allowed values of AdlabGroupCreateV30AdInfoExternalAction enum
var AllowedAdlabGroupCreateV30AdInfoExternalActionEnumValues = []AdlabGroupCreateV30AdInfoExternalAction{
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
	"AD_CONVERT_TYPE_AUTHORIZATION",
	"AD_CONVERT_TYPE_BANKCARD_INFORMATION",
	"AD_CONVERT_TYPE_BOOST",
	"AD_CONVERT_TYPE_BUTTON",
	"AD_CONVERT_TYPE_CERTIFICATION_INFORMATION",
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
	"AD_CONVERT_TYPE_FIRST_RENTAL_ORDER",
	"AD_CONVERT_TYPE_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_FOLLOW_CLICK_PRODUCT",
	"AD_CONVERT_TYPE_FOLLOW_LIVE_ENTER",
	"AD_CONVERT_TYPE_FOLLOW_VIDEO_PLAY_FINISH",
	"AD_CONVERT_TYPE_FORM",
	"AD_CONVERT_TYPE_FORM_ANSWER",
	"AD_CONVERT_TYPE_FORM_CONNECT",
	"AD_CONVERT_TYPE_FORM_DEEP",
	"AD_CONVERT_TYPE_GAMESTATION_DOWNLOAD_DOUPLUS",
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
	"AD_CONVERT_TYPE_LIVE_BUSINESS_FITTING",
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
	"AD_CONVERT_TYPE_WECHAT_PAY",
	"AD_CONVERT_TYPE_WECHAT_REGISTER",
	"AD_CONVERT_TYPE_WECHAT_USER_FIRST_MESSAGE",
	"AD_CONVERT_TYPE_WECHAT_WECOM_ADD",
	"AD_CONVERT_TYPE_XPATH",
}

// NewAdlabGroupCreateV30AdInfoExternalActionFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoExternalActionFromValue(v string) (*AdlabGroupCreateV30AdInfoExternalAction, error) {
	ev := AdlabGroupCreateV30AdInfoExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoExternalAction: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoExternalAction) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_external_action value
func (v AdlabGroupCreateV30AdInfoExternalAction) Ptr() *AdlabGroupCreateV30AdInfoExternalAction {
	return &v
}
