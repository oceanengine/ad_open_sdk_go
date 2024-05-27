/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AudiencePackageGetV2DataAudiencePackagesMarketingGoal
type AudiencePackageGetV2DataAudiencePackagesMarketingGoal string

// List of audience_package_get_v2_data_audience_packages_marketing_goal
const (
	LIVE_AudiencePackageGetV2DataAudiencePackagesMarketingGoal            AudiencePackageGetV2DataAudiencePackagesMarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_AudiencePackageGetV2DataAudiencePackagesMarketingGoal AudiencePackageGetV2DataAudiencePackagesMarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of AudiencePackageGetV2DataAudiencePackagesMarketingGoal enum
var AllowedAudiencePackageGetV2DataAudiencePackagesMarketingGoalEnumValues = []AudiencePackageGetV2DataAudiencePackagesMarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewAudiencePackageGetV2DataAudiencePackagesMarketingGoalFromValue returns a pointer to a valid AudiencePackageGetV2DataAudiencePackagesMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2DataAudiencePackagesMarketingGoalFromValue(v string) (*AudiencePackageGetV2DataAudiencePackagesMarketingGoal, error) {
	ev := AudiencePackageGetV2DataAudiencePackagesMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2DataAudiencePackagesMarketingGoal: valid values are %v", v, AllowedAudiencePackageGetV2DataAudiencePackagesMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2DataAudiencePackagesMarketingGoal) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2DataAudiencePackagesMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_data_audience_packages_marketing_goal value
func (v AudiencePackageGetV2DataAudiencePackagesMarketingGoal) Ptr() *AudiencePackageGetV2DataAudiencePackagesMarketingGoal {
	return &v
}
