/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10ExternalAction
type QianchuanToolsEstimateAudienceV10ExternalAction string

// List of qianchuan_tools_estimate_audience_v1.0_external_action
const (
	AD_CONVERT_TYPE_CARD_ACTIVE_QianchuanToolsEstimateAudienceV10ExternalAction                  QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_CARD_ACTIVE"
	AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction    QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION"
	AD_CONVERT_TYPE_LIVE_COMMENT_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction          QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_COMMENT_ACTION"
	AD_CONVERT_TYPE_LIVE_ENTER_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction            QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_ENTER_ACTION"
	AD_CONVERT_TYPE_LIVE_PAY_ROI_QianchuanToolsEstimateAudienceV10ExternalAction                 QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_PAY_ROI"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction     QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_QianchuanToolsEstimateAudienceV10ExternalAction        QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY"
	AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7_DAYS_QianchuanToolsEstimateAudienceV10ExternalAction QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS"
	AD_CONVERT_TYPE_NEW_FOLLOW_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction            QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_NEW_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_FOLLOW_ACTION_QianchuanToolsEstimateAudienceV10ExternalAction             QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_QC_FOLLOW_ACTION"
	AD_CONVERT_TYPE_QC_MUST_BUY_QianchuanToolsEstimateAudienceV10ExternalAction                  QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_QC_MUST_BUY"
	AD_CONVERT_TYPE_SHOPPING_QianchuanToolsEstimateAudienceV10ExternalAction                     QianchuanToolsEstimateAudienceV10ExternalAction = "AD_CONVERT_TYPE_SHOPPING"
)

// All allowed values of QianchuanToolsEstimateAudienceV10ExternalAction enum
var AllowedQianchuanToolsEstimateAudienceV10ExternalActionEnumValues = []QianchuanToolsEstimateAudienceV10ExternalAction{
	"AD_CONVERT_TYPE_CARD_ACTIVE",
	"AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION",
	"AD_CONVERT_TYPE_LIVE_COMMENT_ACTION",
	"AD_CONVERT_TYPE_LIVE_ENTER_ACTION",
	"AD_CONVERT_TYPE_LIVE_PAY_ROI",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY",
	"AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS",
	"AD_CONVERT_TYPE_NEW_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_FOLLOW_ACTION",
	"AD_CONVERT_TYPE_QC_MUST_BUY",
	"AD_CONVERT_TYPE_SHOPPING",
}

// NewQianchuanToolsEstimateAudienceV10ExternalActionFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10ExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10ExternalActionFromValue(v string) (*QianchuanToolsEstimateAudienceV10ExternalAction, error) {
	ev := QianchuanToolsEstimateAudienceV10ExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10ExternalAction: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10ExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10ExternalAction) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10ExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_external_action value
func (v QianchuanToolsEstimateAudienceV10ExternalAction) Ptr() *QianchuanToolsEstimateAudienceV10ExternalAction {
	return &v
}
