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

// BrandAdCreateV30AudienceInfoAges
type BrandAdCreateV30AudienceInfoAges string

// List of brand_ad_create_v3.0_audience_info_ages
const (
	ABOVE50_BrandAdCreateV30AudienceInfoAges      BrandAdCreateV30AudienceInfoAges = "ABOVE50"
	BETWEEN18_23_BrandAdCreateV30AudienceInfoAges BrandAdCreateV30AudienceInfoAges = "BETWEEN18_23"
	BETWEEN24_30_BrandAdCreateV30AudienceInfoAges BrandAdCreateV30AudienceInfoAges = "BETWEEN24_30"
	BETWEEN31_40_BrandAdCreateV30AudienceInfoAges BrandAdCreateV30AudienceInfoAges = "BETWEEN31_40"
	BETWEEN41_49_BrandAdCreateV30AudienceInfoAges BrandAdCreateV30AudienceInfoAges = "BETWEEN41_49"
	UNLIMITED_BrandAdCreateV30AudienceInfoAges    BrandAdCreateV30AudienceInfoAges = "UNLIMITED"
)

// All allowed values of BrandAdCreateV30AudienceInfoAges enum
var AllowedBrandAdCreateV30AudienceInfoAgesEnumValues = []BrandAdCreateV30AudienceInfoAges{
	"ABOVE50",
	"BETWEEN18_23",
	"BETWEEN24_30",
	"BETWEEN31_40",
	"BETWEEN41_49",
	"UNLIMITED",
}

// NewBrandAdCreateV30AudienceInfoAgesFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoAges
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoAgesFromValue(v string) (*BrandAdCreateV30AudienceInfoAges, error) {
	ev := BrandAdCreateV30AudienceInfoAges(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoAges: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoAgesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoAges) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoAgesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_ages value
func (v BrandAdCreateV30AudienceInfoAges) Ptr() *BrandAdCreateV30AudienceInfoAges {
	return &v
}
