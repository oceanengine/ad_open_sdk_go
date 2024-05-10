/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderCreateV10DeliverySettingExternalAction
type QianchuanAwemeOrderCreateV10DeliverySettingExternalAction string

// List of qianchuan_aweme_order_create_v1.0_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction     QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction           QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction             QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_HEAT_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction                     QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_HEAT"
	AD_CONVERT_TYPE_LIVE_ROI_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction                      QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction      QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction         QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction             QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction              QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction                   QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAwemeOrderCreateV10DeliverySettingExternalAction                      QianchuanAwemeOrderCreateV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAwemeOrderCreateV10DeliverySettingExternalAction enum
var AllowedQianchuanAwemeOrderCreateV10DeliverySettingExternalActionEnumValues = []QianchuanAwemeOrderCreateV10DeliverySettingExternalAction{
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_HEAT",
	"AD_CONVERT_TYPE_LIVE_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanAwemeOrderCreateV10DeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanAwemeOrderCreateV10DeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderCreateV10DeliverySettingExternalActionFromValue(v string) (*QianchuanAwemeOrderCreateV10DeliverySettingExternalAction, error) {
	ev := QianchuanAwemeOrderCreateV10DeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderCreateV10DeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanAwemeOrderCreateV10DeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderCreateV10DeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderCreateV10DeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_create_v1.0_delivery_setting_external_action value
func (v QianchuanAwemeOrderCreateV10DeliverySettingExternalAction) Ptr() *QianchuanAwemeOrderCreateV10DeliverySettingExternalAction {
	return &v
}
