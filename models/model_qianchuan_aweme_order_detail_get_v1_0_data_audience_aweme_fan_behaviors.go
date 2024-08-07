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

// QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors
type QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors string

// List of qianchuan_aweme_order_detail_get_v1.0_data_audience_aweme_fan_behaviors
const (
	ALL_QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors           QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors = "ALL"
	FOLLOWED_USER_QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_SHARE_QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors   QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors = "GOODS_SHARE"
	LIKED_USER_QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors    QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_WATCH_QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors    QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors = "LIVE_WATCH"
)

// All allowed values of QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors enum
var AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues = []QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors{
	"ALL",
	"FOLLOWED_USER",
	"GOODS_SHARE",
	"LIKED_USER",
	"LIVE_WATCH",
}

// NewQianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviorsFromValue(v string) (*QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors, error) {
	ev := QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_aweme_order_detail_get_v1.0_data_audience_aweme_fan_behaviors value
func (v QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors) Ptr() *QianchuanAwemeOrderDetailGetV10DataAudienceAwemeFanBehaviors {
	return &v
}
