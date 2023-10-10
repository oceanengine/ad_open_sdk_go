/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageCreateV2Age
type AudiencePackageCreateV2Age string

// List of audience_package_create_v2_age
const (
	AGE_BELOW_18_AudiencePackageCreateV2Age      AudiencePackageCreateV2Age = "AGE_BELOW_18"
	AGE_BETWEEN_18_19_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_18_19"
	AGE_ABOVE_50_AudiencePackageCreateV2Age      AudiencePackageCreateV2Age = "AGE_ABOVE_50"
	AGE_BETWEEN_41_45_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_41_45"
	AGE_BETWEEN_46_50_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_46_50"
	AGE_BETWEEN_36_40_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_36_40"
	AGE_BETWEEN_41_49_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_41_49"
	AGE_BETWEEN_24_30_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_24_30"
	AGE_BETWEEN_51_55_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_51_55"
	AGE_BETWEEN_20_23_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_20_23"
	AGE_BETWEEN_31_35_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_31_35"
	AGE_BETWEEN_18_23_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_18_23"
	AGE_BETWEEN_31_40_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_31_40"
	AGE_BETWEEN_56_59_AudiencePackageCreateV2Age AudiencePackageCreateV2Age = "AGE_BETWEEN_56_59"
	AGE_ABOVE_60_AudiencePackageCreateV2Age      AudiencePackageCreateV2Age = "AGE_ABOVE_60"
)

// All allowed values of AudiencePackageCreateV2Age enum
var AllowedAudiencePackageCreateV2AgeEnumValues = []AudiencePackageCreateV2Age{
	"AGE_BELOW_18",
	"AGE_BETWEEN_18_19",
	"AGE_ABOVE_50",
	"AGE_BETWEEN_41_45",
	"AGE_BETWEEN_46_50",
	"AGE_BETWEEN_36_40",
	"AGE_BETWEEN_41_49",
	"AGE_BETWEEN_24_30",
	"AGE_BETWEEN_51_55",
	"AGE_BETWEEN_20_23",
	"AGE_BETWEEN_31_35",
	"AGE_BETWEEN_18_23",
	"AGE_BETWEEN_31_40",
	"AGE_BETWEEN_56_59",
	"AGE_ABOVE_60",
}

// NewAudiencePackageCreateV2AgeFromValue returns a pointer to a valid AudiencePackageCreateV2Age
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageCreateV2AgeFromValue(v string) (*AudiencePackageCreateV2Age, error) {
	ev := AudiencePackageCreateV2Age(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageCreateV2Age: valid values are %v", v, AllowedAudiencePackageCreateV2AgeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageCreateV2Age) IsValid() bool {
	for _, existing := range AllowedAudiencePackageCreateV2AgeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_create_v2_age value
func (v AudiencePackageCreateV2Age) Ptr() *AudiencePackageCreateV2Age {
	return &v
}
