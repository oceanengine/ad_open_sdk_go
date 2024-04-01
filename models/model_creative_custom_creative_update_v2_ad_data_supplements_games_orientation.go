/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation
type CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation string

// List of creative_custom_creative_update_v2_ad_data_supplements_games_orientation
const (
	VERTICAL_CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation   CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation = "VERTICAL"
	HORIZONTAL_CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation = "HORIZONTAL"
)

// All allowed values of CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation enum
var AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues = []CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation{
	"VERTICAL",
	"HORIZONTAL",
}

// NewCreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientationFromValue returns a pointer to a valid CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientationFromValue(v string) (*CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation, error) {
	ev := CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation: valid values are %v", v, AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_update_v2_ad_data_supplements_games_orientation value
func (v CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation) Ptr() *CreativeCustomCreativeUpdateV2AdDataSupplementsGamesOrientation {
	return &v
}
