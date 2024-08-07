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

// AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors
type AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors string

// List of audience_package_get_v2_data_audience_packages_audience_aweme_fan_behaviors
const (
	COMMENTED_USER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors       AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "COMMENTED_USER"
	FOLLOWED_USER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors        AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "FOLLOWED_USER"
	GOODS_CARTS_CLICK_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors    AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "GOODS_CARTS_CLICK"
	GOODS_CARTS_ORDER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors    AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "GOODS_CARTS_ORDER"
	LIKED_USER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors           AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIKED_USER"
	LIVE_COMMENT_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors         AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_COMMENT"
	LIVE_EFFECTIVE_WATCH_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_EFFECTIVE_WATCH"
	LIVE_EXCEPTIONAL_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors     AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_EXCEPTIONAL"
	LIVE_GOODS_CLICK_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors     AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_GOODS_CLICK"
	LIVE_GOODS_ORDER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors     AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_GOODS_ORDER"
	LIVE_WATCH_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors           AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "LIVE_WATCH"
	SHARED_USER_AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors          AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors = "SHARED_USER"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors enum
var AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviorsEnumValues = []AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors{
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

// NewAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviorsFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviorsFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviorsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviorsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_audience_aweme_fan_behaviors value
func (v AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors) Ptr() *AudiencePackageGetV2DataAudiencePackagesAudienceAwemeFanBehaviors {
	return &v
}
