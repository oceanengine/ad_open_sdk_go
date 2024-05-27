/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction
type QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction string

// List of qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction     QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction           QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_EFFECTIVE_EFFECTIVE_VIEW"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction             QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_HEAT_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction                     QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_HEAT"
	AD_CONVERT_TYPE_LIVE_ROI_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction                      QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction      QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction         QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction             QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction              QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction                   QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction                      QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction enum
var AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalActionEnumValues = []QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction{
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

// NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalActionFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_delivery_setting_external_action value
func (v QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction) Ptr() *QianchuanAwemeOrderDetailGetV10DataDeliverySettingExternalAction {
	return &v
}
