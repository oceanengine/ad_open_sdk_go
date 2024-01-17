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

// AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets
type AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets string

// List of adlab_group_detail_v3.0_data_data_ad_info_audience_auto_extend_targets
const (
	AD_TAG_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets          AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "AD_TAG"
	AGE_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets             AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "AGE"
	CUSTOM_AUDIENCE_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	GENDER_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets          AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "GENDER"
	INTEREST_TAG_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets    AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "INTEREST_TAG"
	REGION_AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets          AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets = "REGION"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets enum
var AllowedAdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargetsEnumValues = []AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets{
	"AD_TAG",
	"AGE",
	"CUSTOM_AUDIENCE",
	"GENDER",
	"INTEREST_TAG",
	"REGION",
}

// NewAdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargetsFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargetsFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_audience_auto_extend_targets value
func (v AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets) Ptr() *AdlabGroupDetailV30DataDataAdInfoAudienceAutoExtendTargets {
	return &v
}
