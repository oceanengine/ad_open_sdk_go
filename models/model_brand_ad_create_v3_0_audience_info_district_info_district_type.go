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

// BrandAdCreateV30AudienceInfoDistrictInfoDistrictType
type BrandAdCreateV30AudienceInfoDistrictInfoDistrictType string

// List of brand_ad_create_v3.0_audience_info_district_info_district_type
const (
	REGION_BrandAdCreateV30AudienceInfoDistrictInfoDistrictType    BrandAdCreateV30AudienceInfoDistrictInfoDistrictType = "REGION"
	UNLIMITED_BrandAdCreateV30AudienceInfoDistrictInfoDistrictType BrandAdCreateV30AudienceInfoDistrictInfoDistrictType = "UNLIMITED"
)

// All allowed values of BrandAdCreateV30AudienceInfoDistrictInfoDistrictType enum
var AllowedBrandAdCreateV30AudienceInfoDistrictInfoDistrictTypeEnumValues = []BrandAdCreateV30AudienceInfoDistrictInfoDistrictType{
	"REGION",
	"UNLIMITED",
}

// NewBrandAdCreateV30AudienceInfoDistrictInfoDistrictTypeFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoDistrictInfoDistrictType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoDistrictInfoDistrictTypeFromValue(v string) (*BrandAdCreateV30AudienceInfoDistrictInfoDistrictType, error) {
	ev := BrandAdCreateV30AudienceInfoDistrictInfoDistrictType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoDistrictInfoDistrictType: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoDistrictInfoDistrictTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoDistrictInfoDistrictType) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoDistrictInfoDistrictTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_district_info_district_type value
func (v BrandAdCreateV30AudienceInfoDistrictInfoDistrictType) Ptr() *BrandAdCreateV30AudienceInfoDistrictInfoDistrictType {
	return &v
}
