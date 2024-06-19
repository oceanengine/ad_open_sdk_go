/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AssetsCreativeComponentGetV2FilteringStatus
type AssetsCreativeComponentGetV2FilteringStatus string

// List of assets_creative_component_get_v2_filtering_status
const (
	PASS_AssetsCreativeComponentGetV2FilteringStatus   AssetsCreativeComponentGetV2FilteringStatus = "PASS"
	REJECT_AssetsCreativeComponentGetV2FilteringStatus AssetsCreativeComponentGetV2FilteringStatus = "REJECT"
	UNDER_AssetsCreativeComponentGetV2FilteringStatus  AssetsCreativeComponentGetV2FilteringStatus = "UNDER"
)

// All allowed values of AssetsCreativeComponentGetV2FilteringStatus enum
var AllowedAssetsCreativeComponentGetV2FilteringStatusEnumValues = []AssetsCreativeComponentGetV2FilteringStatus{
	"PASS",
	"REJECT",
	"UNDER",
}

// NewAssetsCreativeComponentGetV2FilteringStatusFromValue returns a pointer to a valid AssetsCreativeComponentGetV2FilteringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetsCreativeComponentGetV2FilteringStatusFromValue(v string) (*AssetsCreativeComponentGetV2FilteringStatus, error) {
	ev := AssetsCreativeComponentGetV2FilteringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetsCreativeComponentGetV2FilteringStatus: valid values are %v", v, AllowedAssetsCreativeComponentGetV2FilteringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetsCreativeComponentGetV2FilteringStatus) IsValid() bool {
	for _, existing := range AllowedAssetsCreativeComponentGetV2FilteringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to assets_creative_component_get_v2_filtering_status value
func (v AssetsCreativeComponentGetV2FilteringStatus) Ptr() *AssetsCreativeComponentGetV2FilteringStatus {
	return &v
}
