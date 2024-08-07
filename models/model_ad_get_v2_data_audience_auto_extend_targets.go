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

// AdGetV2DataAudienceAutoExtendTargets
type AdGetV2DataAudienceAutoExtendTargets string

// List of ad_get_v2_data_audience_auto_extend_targets
const (
	INTEREST_TAG_AdGetV2DataAudienceAutoExtendTargets    AdGetV2DataAudienceAutoExtendTargets = "INTEREST_TAG"
	GENDER_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "GENDER"
	AGE_AdGetV2DataAudienceAutoExtendTargets             AdGetV2DataAudienceAutoExtendTargets = "AGE"
	REGION_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "REGION"
	CUSTOM_AUDIENCE_AdGetV2DataAudienceAutoExtendTargets AdGetV2DataAudienceAutoExtendTargets = "CUSTOM_AUDIENCE"
	AD_TAG_AdGetV2DataAudienceAutoExtendTargets          AdGetV2DataAudienceAutoExtendTargets = "AD_TAG"
)

// All allowed values of AdGetV2DataAudienceAutoExtendTargets enum
var AllowedAdGetV2DataAudienceAutoExtendTargetsEnumValues = []AdGetV2DataAudienceAutoExtendTargets{
	"INTEREST_TAG",
	"GENDER",
	"AGE",
	"REGION",
	"CUSTOM_AUDIENCE",
	"AD_TAG",
}

// NewAdGetV2DataAudienceAutoExtendTargetsFromValue returns a pointer to a valid AdGetV2DataAudienceAutoExtendTargets
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataAudienceAutoExtendTargetsFromValue(v string) (*AdGetV2DataAudienceAutoExtendTargets, error) {
	ev := AdGetV2DataAudienceAutoExtendTargets(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataAudienceAutoExtendTargets: valid values are %v", v, AllowedAdGetV2DataAudienceAutoExtendTargetsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataAudienceAutoExtendTargets) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataAudienceAutoExtendTargetsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_audience_auto_extend_targets value
func (v AdGetV2DataAudienceAutoExtendTargets) Ptr() *AdGetV2DataAudienceAutoExtendTargets {
	return &v
}
