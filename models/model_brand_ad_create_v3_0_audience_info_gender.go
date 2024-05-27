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

// BrandAdCreateV30AudienceInfoGender
type BrandAdCreateV30AudienceInfoGender string

// List of brand_ad_create_v3.0_audience_info_gender
const (
	FEMALE_BrandAdCreateV30AudienceInfoGender    BrandAdCreateV30AudienceInfoGender = "FEMALE"
	MALE_BrandAdCreateV30AudienceInfoGender      BrandAdCreateV30AudienceInfoGender = "MALE"
	UNLIMITED_BrandAdCreateV30AudienceInfoGender BrandAdCreateV30AudienceInfoGender = "UNLIMITED"
)

// All allowed values of BrandAdCreateV30AudienceInfoGender enum
var AllowedBrandAdCreateV30AudienceInfoGenderEnumValues = []BrandAdCreateV30AudienceInfoGender{
	"FEMALE",
	"MALE",
	"UNLIMITED",
}

// NewBrandAdCreateV30AudienceInfoGenderFromValue returns a pointer to a valid BrandAdCreateV30AudienceInfoGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AudienceInfoGenderFromValue(v string) (*BrandAdCreateV30AudienceInfoGender, error) {
	ev := BrandAdCreateV30AudienceInfoGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AudienceInfoGender: valid values are %v", v, AllowedBrandAdCreateV30AudienceInfoGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AudienceInfoGender) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AudienceInfoGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_audience_info_gender value
func (v BrandAdCreateV30AudienceInfoGender) Ptr() *BrandAdCreateV30AudienceInfoGender {
	return &v
}
