/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30NativeSettingIsFeedAndFavSee
type PromotionCreateV30NativeSettingIsFeedAndFavSee string

// List of promotion_create_v3.0_native_setting_is_feed_and_fav_see
const (
	OFF_PromotionCreateV30NativeSettingIsFeedAndFavSee PromotionCreateV30NativeSettingIsFeedAndFavSee = "OFF"
	ON_PromotionCreateV30NativeSettingIsFeedAndFavSee  PromotionCreateV30NativeSettingIsFeedAndFavSee = "ON"
)

// All allowed values of PromotionCreateV30NativeSettingIsFeedAndFavSee enum
var AllowedPromotionCreateV30NativeSettingIsFeedAndFavSeeEnumValues = []PromotionCreateV30NativeSettingIsFeedAndFavSee{
	"OFF",
	"ON",
}

// NewPromotionCreateV30NativeSettingIsFeedAndFavSeeFromValue returns a pointer to a valid PromotionCreateV30NativeSettingIsFeedAndFavSee
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30NativeSettingIsFeedAndFavSeeFromValue(v string) (*PromotionCreateV30NativeSettingIsFeedAndFavSee, error) {
	ev := PromotionCreateV30NativeSettingIsFeedAndFavSee(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30NativeSettingIsFeedAndFavSee: valid values are %v", v, AllowedPromotionCreateV30NativeSettingIsFeedAndFavSeeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30NativeSettingIsFeedAndFavSee) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30NativeSettingIsFeedAndFavSeeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_native_setting_is_feed_and_fav_see value
func (v PromotionCreateV30NativeSettingIsFeedAndFavSee) Ptr() *PromotionCreateV30NativeSettingIsFeedAndFavSee {
	return &v
}
