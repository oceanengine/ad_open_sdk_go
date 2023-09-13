/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors
type QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors string

// List of qianchuan_tools_estimate_audience_v1.0_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors       QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors        QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors    QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors    QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors           QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors         QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors     QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors     QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors     QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors           QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors          QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors enum
var AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsEnumValues = []QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors{
	"COMMENTED_USER",
	"FOLLOWED_USER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIKED_USER",
	"LIVE_COMMENT",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIVE_GOODS_CLICK",
	"LIVE_GOODS_ORDER",
	"LIVE_WATCH",
	"SHARED_USER",
}

// NewQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsFromValue(v string) (*QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors, error) {
	ev := QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_tools_estimate_audience_v1.0_audience_aweme_fan_behaviors value
func (v QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors) Ptr() *QianchuanToolsEstimateAudienceV10AudienceAwemeFanBehaviors {
	return &v
}
