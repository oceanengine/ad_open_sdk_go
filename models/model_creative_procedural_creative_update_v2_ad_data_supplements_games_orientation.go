/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation
type CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation string

// List of creative_procedural_creative_update_v2_ad_data_supplements_games_orientation
const (
	VERTICAL_CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation   CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation = "VERTICAL"
	HORIZONTAL_CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation = "HORIZONTAL"
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues = []CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation{
	"VERTICAL",
	"HORIZONTAL",
}

// NewCreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientationFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientationFromValue(v string) (*CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_supplements_games_orientation value
func (v CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation) Ptr() *CreativeProceduralCreativeUpdateV2AdDataSupplementsGamesOrientation {
	return &v
}
