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

// CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee
type CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee int64

// List of creative_procedural_creative_update_v2_ad_data_is_feed_and_fav_see
const (
	Enum_0_CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee = 0
	Enum_1_CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee = 1
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues = []CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee{
	0,
	1,
}

// NewCreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSeeFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSeeFromValue(v int64) (*CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSeeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_is_feed_and_fav_see value
func (v CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee) Ptr() *CreativeProceduralCreativeUpdateV2AdDataIsFeedAndFavSee {
	return &v
}
