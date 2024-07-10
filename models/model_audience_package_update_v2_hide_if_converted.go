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

// AudiencePackageUpdateV2HideIfConverted
type AudiencePackageUpdateV2HideIfConverted string

// List of audience_package_update_v2_hide_if_converted
const (
	AD_AudiencePackageUpdateV2HideIfConverted           AudiencePackageUpdateV2HideIfConverted = "AD"
	APP_AudiencePackageUpdateV2HideIfConverted          AudiencePackageUpdateV2HideIfConverted = "APP"
	CAMPAIGN_AudiencePackageUpdateV2HideIfConverted     AudiencePackageUpdateV2HideIfConverted = "CAMPAIGN"
	NO_EXCLUDE_AudiencePackageUpdateV2HideIfConverted   AudiencePackageUpdateV2HideIfConverted = "NO_EXCLUDE"
	ORGANIZATION_AudiencePackageUpdateV2HideIfConverted AudiencePackageUpdateV2HideIfConverted = "ORGANIZATION"
	ADVERTISER_AudiencePackageUpdateV2HideIfConverted   AudiencePackageUpdateV2HideIfConverted = "ADVERTISER"
	CUSTOMER_AudiencePackageUpdateV2HideIfConverted     AudiencePackageUpdateV2HideIfConverted = "CUSTOMER"
)

// All allowed values of AudiencePackageUpdateV2HideIfConverted enum
var AllowedAudiencePackageUpdateV2HideIfConvertedEnumValues = []AudiencePackageUpdateV2HideIfConverted{
	"AD",
	"APP",
	"CAMPAIGN",
	"NO_EXCLUDE",
	"ORGANIZATION",
	"ADVERTISER",
	"CUSTOMER",
}

// NewAudiencePackageUpdateV2HideIfConvertedFromValue returns a pointer to a valid AudiencePackageUpdateV2HideIfConverted
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAudiencePackageUpdateV2HideIfConvertedFromValue(v string) (*AudiencePackageUpdateV2HideIfConverted, error) {
	ev := AudiencePackageUpdateV2HideIfConverted(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AudiencePackageUpdateV2HideIfConverted: valid values are %v", v, AllowedAudiencePackageUpdateV2HideIfConvertedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AudiencePackageUpdateV2HideIfConverted) IsValid() bool {
	for _, existing := range AllowedAudiencePackageUpdateV2HideIfConvertedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to audience_package_update_v2_hide_if_converted value
func (v AudiencePackageUpdateV2HideIfConverted) Ptr() *AudiencePackageUpdateV2HideIfConverted {
	return &v
}
