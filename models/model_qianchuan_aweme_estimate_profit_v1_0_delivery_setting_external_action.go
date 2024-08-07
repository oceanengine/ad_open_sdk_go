/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction
type QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction string

// List of qianchuan_aweme_estimate_profit_v1.0_delivery_setting_external_action
const (
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_HEAT_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction                 QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_HEAT"
	AD_CONVERT_TYPE_LIVE_ROI_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction                  QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction  QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction     QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction         QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction          QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction               QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction                  QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction enum
var AllowedQianchuanAwemeEstimateProfitV10DeliverySettingExternalActionEnumValues = []QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction{
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_HEAT",
	"AD_CONVERT_TYPE_LIVE_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanAwemeEstimateProfitV10DeliverySettingExternalActionFromValue returns a pointer to a valid QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeEstimateProfitV10DeliverySettingExternalActionFromValue(v string) (*QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction, error) {
	ev := QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction: valid values are %v", v, AllowedQianchuanAwemeEstimateProfitV10DeliverySettingExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeEstimateProfitV10DeliverySettingExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_estimate_profit_v1.0_delivery_setting_external_action value
func (v QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction) Ptr() *QianchuanAwemeEstimateProfitV10DeliverySettingExternalAction {
	return &v
}
