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

// AudiencePackageUpdateV2AwemeFanBehaviors
type AudiencePackageUpdateV2AwemeFanBehaviors string

// List of audience_package_update_v2_aweme_fan_behaviors
const (
	LIVE_WATCH_AudiencePackageUpdateV2AwemeFanBehaviors           AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_WATCH"
	FOLLOWED_USER_AudiencePackageUpdateV2AwemeFanBehaviors        AudiencePackageUpdateV2AwemeFanBehaviors = "FOLLOWED_USER"
	LIVE_COMMENT_AudiencePackageUpdateV2AwemeFanBehaviors         AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_GOODS_ORDER_AudiencePackageUpdateV2AwemeFanBehaviors     AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_GOODS_ORDER"
	GOODS_CARTS_CLICK_AudiencePackageUpdateV2AwemeFanBehaviors    AudiencePackageUpdateV2AwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_AudiencePackageUpdateV2AwemeFanBehaviors    AudiencePackageUpdateV2AwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIVE_GOODS_CLICK_AudiencePackageUpdateV2AwemeFanBehaviors     AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_EFFECTIVE_WATCH_AudiencePackageUpdateV2AwemeFanBehaviors AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_AudiencePackageUpdateV2AwemeFanBehaviors     AudiencePackageUpdateV2AwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIKED_USER_AudiencePackageUpdateV2AwemeFanBehaviors           AudiencePackageUpdateV2AwemeFanBehaviors = "LIKED_USER"
	SHARED_USER_AudiencePackageUpdateV2AwemeFanBehaviors          AudiencePackageUpdateV2AwemeFanBehaviors = "SHARED_USER"
	COMMENTED_USER_AudiencePackageUpdateV2AwemeFanBehaviors       AudiencePackageUpdateV2AwemeFanBehaviors = "COMMENTED_USER"
)

// All allowed values of AudiencePackageUpdateV2AwemeFanBehaviors enum
var AllowedAudiencePackageUpdateV2AwemeFanBehaviorsEnumValues = []AudiencePackageUpdateV2AwemeFanBehaviors{
	"LIVE_WATCH",
	"FOLLOWED_USER",
	"LIVE_COMMENT",
	"LIVE_GOODS_ORDER",
	"GOODS_CARTS_CLICK",
	"GOODS_CARTS_ORDER",
	"LIVE_GOODS_CLICK",
	"LIVE_EFFECTIVE_WATCH",
	"LIVE_EXCEPTIONAL",
	"LIKED_USER",
	"SHARED_USER",
	"COMMENTED_USER",
}

// NewAudiencePackageUpdateV2AwemeFanBehaviorsFromValue returns a pointer to a valid AudiencePackageUpdateV2AwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2AwemeFanBehaviorsFromValue(v string) (*AudiencePackageUpdateV2AwemeFanBehaviors, error) {
	ev := AudiencePackageUpdateV2AwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2AwemeFanBehaviors: valid values are %v", v, AllowedAudiencePackageUpdateV2AwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2AwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2AwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_aweme_fan_behaviors value
func (v AudiencePackageUpdateV2AwemeFanBehaviors) Ptr() *AudiencePackageUpdateV2AwemeFanBehaviors {
	return &v
}
