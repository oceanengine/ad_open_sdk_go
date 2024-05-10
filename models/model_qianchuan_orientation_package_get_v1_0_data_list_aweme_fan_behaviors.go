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

// QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors
type QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors string

// List of qianchuan_orientation_package_get_v1.0_data_list_aweme_fan_behaviors
const (
	COMMENTED_USER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors       QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors        QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors    QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors    QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors           QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors         QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors     QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors     QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors     QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors           QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors          QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors enum
var AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsEnumValues = []QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors{
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

// NewQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsFromValue returns a pointer to a valid QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsFromValue(v string) (*QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors, error) {
	ev := QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors: valid values are %v", v, AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedQianchuanOrientationPackageGetV10DataListAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_orientation_package_get_v1.0_data_list_aweme_fan_behaviors value
func (v QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors) Ptr() *QianchuanOrientationPackageGetV10DataListAwemeFanBehaviors {
	return &v
}
