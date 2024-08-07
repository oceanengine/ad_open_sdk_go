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

// QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors
type QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors string

// List of qianchuan_ad_detail_get_v1.0_data_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors       QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors        QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors    QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors    QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors           QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors         QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors     QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors     QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors     QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors           QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors          QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors enum
var AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues = []QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors{
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

// NewQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsFromValue(v string) (*QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors, error) {
	ev := QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_audience_aweme_fan_behaviors value
func (v QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors) Ptr() *QianchuanAdDetailGetV10DataAudienceAwemeFanBehaviors {
	return &v
}
