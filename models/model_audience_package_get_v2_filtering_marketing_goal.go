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

// AudiencePackageGetV2FilteringMarketingGoal
type AudiencePackageGetV2FilteringMarketingGoal string

// List of audience_package_get_v2_filtering_marketing_goal
const (
	LIVE_AudiencePackageGetV2FilteringMarketingGoal            AudiencePackageGetV2FilteringMarketingGoal = "LIVE"
	VIDEO_AND_IMAGE_AudiencePackageGetV2FilteringMarketingGoal AudiencePackageGetV2FilteringMarketingGoal = "VIDEO_AND_IMAGE"
)

// All allowed values of AudiencePackageGetV2FilteringMarketingGoal enum
var AllowedAudiencePackageGetV2FilteringMarketingGoalEnumValues = []AudiencePackageGetV2FilteringMarketingGoal{
	"LIVE",
	"VIDEO_AND_IMAGE",
}

// NewAudiencePackageGetV2FilteringMarketingGoalFromValue returns a pointer to a valid AudiencePackageGetV2FilteringMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageGetV2FilteringMarketingGoalFromValue(v string) (*AudiencePackageGetV2FilteringMarketingGoal, error) {
	ev := AudiencePackageGetV2FilteringMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageGetV2FilteringMarketingGoal: valid values are %v", v, AllowedAudiencePackageGetV2FilteringMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageGetV2FilteringMarketingGoal) IsValid() bool {
	for _, existing := range AllowedAudiencePackageGetV2FilteringMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_get_v2_filtering_marketing_goal value
func (v AudiencePackageGetV2FilteringMarketingGoal) Ptr() *AudiencePackageGetV2FilteringMarketingGoal {
	return &v
}
