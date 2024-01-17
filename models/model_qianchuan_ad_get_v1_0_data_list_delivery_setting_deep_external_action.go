/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdGetV10DataListDeliverySettingDeepExternalAction
type QianchuanAdGetV10DataListDeliverySettingDeepExternalAction string

// List of qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction                  QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction    QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction          QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction            QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction                 QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction     QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction        QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction     QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction            QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction             QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction                  QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAdGetV10DataListDeliverySettingDeepExternalAction                     QianchuanAdGetV10DataListDeliverySettingDeepExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAdGetV10DataListDeliverySettingDeepExternalAction enum
var AllowedQianchuanAdGetV10DataListDeliverySettingDeepExternalActionEnumValues = []QianchuanAdGetV10DataListDeliverySettingDeepExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_SETTLE",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanAdGetV10DataListDeliverySettingDeepExternalActionFromValue returns a pointer to a valid QianchuanAdGetV10DataListDeliverySettingDeepExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdGetV10DataListDeliverySettingDeepExternalActionFromValue(v string) (*QianchuanAdGetV10DataListDeliverySettingDeepExternalAction, error) {
	ev := QianchuanAdGetV10DataListDeliverySettingDeepExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdGetV10DataListDeliverySettingDeepExternalAction: valid values are %v", v, AllowedQianchuanAdGetV10DataListDeliverySettingDeepExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdGetV10DataListDeliverySettingDeepExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAdGetV10DataListDeliverySettingDeepExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_get_v1.0_data_list_delivery_setting_deep_external_action value
func (v QianchuanAdGetV10DataListDeliverySettingDeepExternalAction) Ptr() *QianchuanAdGetV10DataListDeliverySettingDeepExternalAction {
	return &v
}
