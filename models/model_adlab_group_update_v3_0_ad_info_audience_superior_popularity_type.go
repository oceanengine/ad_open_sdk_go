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

// AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType
type AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType string

// List of adlab_group_update_v3.0_ad_info_audience_superior_popularity_type
const (
	GAME_AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType = "GAME"
	NONE_AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType enum
var AllowedAdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityTypeEnumValues = []AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityTypeFromValue(v string) (*AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType, error) {
	ev := AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType: valid values are %v", v, AllowedAdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_update_v3.0_ad_info_audience_superior_popularity_type value
func (v AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType) Ptr() *AdlabGroupUpdateV30AdInfoAudienceSuperiorPopularityType {
	return &v
}
