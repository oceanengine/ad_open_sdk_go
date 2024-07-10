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

// BrandAdCreateV30AudienceInfoAc
type BrandAdCreateV30AudienceInfoAc string

// List of brand_ad_create_v3.0_audience_info_ac
const (
	UNLIMITED_BrandAdCreateV30AudienceInfoAc BrandAdCreateV30AudienceInfoAc = "UNLIMITED"
	WIFI_BrandAdCreateV30AudienceInfoAc      BrandAdCreateV30AudienceInfoAc = "WIFI"
)

// All allowed values of BrandAdCreateV30AudienceInfoAc enum
var AllowedBrandAdCreateV30AudienceInfoAcEnumValues = []BrandAdCreateV30AudienceInfoAc{
	"UNLIMITED",
	"WIFI",
}

// NewBrandAdCreateV30AudienceInfoAcFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoAc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoAcFromValue(v string) (*BrandAdCreateV30AudienceInfoAc, error) {
	ev := BrandAdCreateV30AudienceInfoAc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoAc: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoAcEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoAc) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoAcEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_ac value
func (v BrandAdCreateV30AudienceInfoAc) Ptr() *BrandAdCreateV30AudienceInfoAc {
	return &v
}
