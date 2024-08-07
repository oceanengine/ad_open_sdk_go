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

// QianchuanAdUpdateV10AudienceAwemeFanBehaviors
type QianchuanAdUpdateV10AudienceAwemeFanBehaviors string

// List of qianchuan_ad_update_v1.0_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors       QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors        QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_QianchuanAdUpdateV10AudienceAwemeFanBehaviors    QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors    QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors           QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_QianchuanAdUpdateV10AudienceAwemeFanBehaviors         QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_QianchuanAdUpdateV10AudienceAwemeFanBehaviors QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_QianchuanAdUpdateV10AudienceAwemeFanBehaviors     QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_QianchuanAdUpdateV10AudienceAwemeFanBehaviors     QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors     QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_QianchuanAdUpdateV10AudienceAwemeFanBehaviors           QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_QianchuanAdUpdateV10AudienceAwemeFanBehaviors          QianchuanAdUpdateV10AudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of QianchuanAdUpdateV10AudienceAwemeFanBehaviors enum
var AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsEnumValues = []QianchuanAdUpdateV10AudienceAwemeFanBehaviors{
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

// NewQianchuanAdUpdateV10AudienceAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanAdUpdateV10AudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10AudienceAwemeFanBehaviorsFromValue(v string) (*QianchuanAdUpdateV10AudienceAwemeFanBehaviors, error) {
	ev := QianchuanAdUpdateV10AudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10AudienceAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10AudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10AudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_audience_aweme_fan_behaviors value
func (v QianchuanAdUpdateV10AudienceAwemeFanBehaviors) Ptr() *QianchuanAdUpdateV10AudienceAwemeFanBehaviors {
	return &v
}
