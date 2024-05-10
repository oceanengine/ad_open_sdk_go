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

// CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee
type CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee int64

// List of creative_custom_creative_update_v2_ad_data_is_feed_and_fav_see
const (
	Enum_0_CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee = 0
	Enum_1_CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee = 1
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee enum
var AllowedCreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues = []CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee{
	0,
	1,
}

// NewCreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSeeFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSeeFromValue(v int64) (*CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_is_feed_and_fav_see value
func (v CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee) Ptr() *CreativeCustomCreativeUpdateV2AdDataIsFeedAndFavSee {
	return &v
}
