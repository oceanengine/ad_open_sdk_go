/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2LandingType
type AudiencePackageCreateV2LandingType string

// List of audience_package_create_v2_landing_type
const (
	AWEME_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "AWEME"
	APP_IOS_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "APP_IOS"
	EXTERNAL_AudiencePackageCreateV2LandingType    AudiencePackageCreateV2LandingType = "EXTERNAL"
	SHOP_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "SHOP"
	QUICK_APP_AudiencePackageCreateV2LandingType   AudiencePackageCreateV2LandingType = "QUICK_APP"
	DPA_AudiencePackageCreateV2LandingType         AudiencePackageCreateV2LandingType = "DPA"
	APP_ANDROID_AudiencePackageCreateV2LandingType AudiencePackageCreateV2LandingType = "APP_ANDROID"
	ARTICLE_AudiencePackageCreateV2LandingType     AudiencePackageCreateV2LandingType = "ARTICLE"
	GOODS_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "GOODS"
	STORE_AudiencePackageCreateV2LandingType       AudiencePackageCreateV2LandingType = "STORE"
	LIVE_AudiencePackageCreateV2LandingType        AudiencePackageCreateV2LandingType = "LIVE"
	MICRO_GAME_AudiencePackageCreateV2LandingType  AudiencePackageCreateV2LandingType = "MICRO_GAME"
)

// All allowed values of AudiencePackageCreateV2LandingType enum
var AllowedAudiencePackageCreateV2LandingTypeEnumValues = []AudiencePackageCreateV2LandingType{
	"AWEME",
	"APP_IOS",
	"EXTERNAL",
	"SHOP",
	"QUICK_APP",
	"DPA",
	"APP_ANDROID",
	"ARTICLE",
	"GOODS",
	"STORE",
	"LIVE",
	"MICRO_GAME",
}

// NewAudiencePackageCreateV2LandingTypeFromValue returns a pointer to a valid AudiencePackageCreateV2LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2LandingTypeFromValue(v string) (*AudiencePackageCreateV2LandingType, error) {
	ev := AudiencePackageCreateV2LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2LandingType: valid values are %v", v, AllowedAudiencePackageCreateV2LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2LandingType) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_landing_type value
func (v AudiencePackageCreateV2LandingType) Ptr() *AudiencePackageCreateV2LandingType {
	return &v
}
