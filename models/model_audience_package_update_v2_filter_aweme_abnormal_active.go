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

// AudiencePackageUpdateV2FilterAwemeAbnormalActive
type AudiencePackageUpdateV2FilterAwemeAbnormalActive int64

// List of audience_package_update_v2_filter_aweme_abnormal_active
const (
	Enum_0_AudiencePackageUpdateV2FilterAwemeAbnormalActive AudiencePackageUpdateV2FilterAwemeAbnormalActive = 0
	Enum_1_AudiencePackageUpdateV2FilterAwemeAbnormalActive AudiencePackageUpdateV2FilterAwemeAbnormalActive = 1
)

// All allowed values of AudiencePackageUpdateV2FilterAwemeAbnormalActive enum
var AllowedAudiencePackageUpdateV2FilterAwemeAbnormalActiveEnumValues = []AudiencePackageUpdateV2FilterAwemeAbnormalActive{
	0,
	1,
}

// NewAudiencePackageUpdateV2FilterAwemeAbnormalActiveFromValue returns a pointer to a valid AudiencePackageUpdateV2FilterAwemeAbnormalActive
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2FilterAwemeAbnormalActiveFromValue(v int64) (*AudiencePackageUpdateV2FilterAwemeAbnormalActive, error) {
	ev := AudiencePackageUpdateV2FilterAwemeAbnormalActive(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2FilterAwemeAbnormalActive: valid values are %v", v, AllowedAudiencePackageUpdateV2FilterAwemeAbnormalActiveEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2FilterAwemeAbnormalActive) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2FilterAwemeAbnormalActiveEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_filter_aweme_abnormal_active value
func (v AudiencePackageUpdateV2FilterAwemeAbnormalActive) Ptr() *AudiencePackageUpdateV2FilterAwemeAbnormalActive {
	return &v
}
