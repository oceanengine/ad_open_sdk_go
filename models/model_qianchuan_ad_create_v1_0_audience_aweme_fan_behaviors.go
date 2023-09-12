/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10AudienceAwemeFanBehaviors
type QianchuanAdCreateV10AudienceAwemeFanBehaviors string

// List of qianchuan_ad_create_v1.0_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_QianchuanAdCreateV10AudienceAwemeFanBehaviors       QianchuanAdCreateV10AudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_QianchuanAdCreateV10AudienceAwemeFanBehaviors        QianchuanAdCreateV10AudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_QianchuanAdCreateV10AudienceAwemeFanBehaviors    QianchuanAdCreateV10AudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_QianchuanAdCreateV10AudienceAwemeFanBehaviors    QianchuanAdCreateV10AudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_QianchuanAdCreateV10AudienceAwemeFanBehaviors           QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_QianchuanAdCreateV10AudienceAwemeFanBehaviors         QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_QianchuanAdCreateV10AudienceAwemeFanBehaviors QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_QianchuanAdCreateV10AudienceAwemeFanBehaviors     QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_QianchuanAdCreateV10AudienceAwemeFanBehaviors     QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_QianchuanAdCreateV10AudienceAwemeFanBehaviors     QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_QianchuanAdCreateV10AudienceAwemeFanBehaviors           QianchuanAdCreateV10AudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_QianchuanAdCreateV10AudienceAwemeFanBehaviors          QianchuanAdCreateV10AudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of QianchuanAdCreateV10AudienceAwemeFanBehaviors enum
var AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsEnumValues = []QianchuanAdCreateV10AudienceAwemeFanBehaviors{
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

// NewQianchuanAdCreateV10AudienceAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanAdCreateV10AudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10AudienceAwemeFanBehaviorsFromValue(v string) (*QianchuanAdCreateV10AudienceAwemeFanBehaviors, error) {
	ev := QianchuanAdCreateV10AudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10AudienceAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10AudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10AudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_audience_aweme_fan_behaviors value
func (v QianchuanAdCreateV10AudienceAwemeFanBehaviors) Ptr() *QianchuanAdCreateV10AudienceAwemeFanBehaviors {
	return &v
}
