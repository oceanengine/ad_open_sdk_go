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

// AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType
type AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType string

// List of adlab_group_create_v3.0_ad_info_audience_superior_popularity_type
const (
	GAME_AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType = "GAME"
	NONE_AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType = "NONE"
)

// All allowed values of AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType enum
var AllowedAdlabGroupCreateV30AdInfoAudienceSuperiorPopularityTypeEnumValues = []AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType{
	"GAME",
	"NONE",
}

// NewAdlabGroupCreateV30AdInfoAudienceSuperiorPopularityTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoAudienceSuperiorPopularityTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType, error) {
	ev := AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoAudienceSuperiorPopularityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoAudienceSuperiorPopularityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_audience_superior_popularity_type value
func (v AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType) Ptr() *AdlabGroupCreateV30AdInfoAudienceSuperiorPopularityType {
	return &v
}
