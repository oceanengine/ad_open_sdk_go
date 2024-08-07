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

// BrandAdGetV30AdForm
type BrandAdGetV30AdForm int64

// List of brand_ad_get_v3.0_ad_form
const (
	Enum_1_BrandAdGetV30AdForm BrandAdGetV30AdForm = 1
	Enum_2_BrandAdGetV30AdForm BrandAdGetV30AdForm = 2
	Enum_3_BrandAdGetV30AdForm BrandAdGetV30AdForm = 3
	Enum_4_BrandAdGetV30AdForm BrandAdGetV30AdForm = 4
	Enum_5_BrandAdGetV30AdForm BrandAdGetV30AdForm = 5
	Enum_6_BrandAdGetV30AdForm BrandAdGetV30AdForm = 6
	Enum_7_BrandAdGetV30AdForm BrandAdGetV30AdForm = 7
	Enum_8_BrandAdGetV30AdForm BrandAdGetV30AdForm = 8
)

// All allowed values of BrandAdGetV30AdForm enum
var AllowedBrandAdGetV30AdFormEnumValues = []BrandAdGetV30AdForm{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
}

// NewBrandAdGetV30AdFormFromValue returns a pointer to a valid BrandAdGetV30AdForm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30AdFormFromValue(v int64) (*BrandAdGetV30AdForm, error) {
	ev := BrandAdGetV30AdForm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30AdForm: valid values are %v", v, AllowedBrandAdGetV30AdFormEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30AdForm) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30AdFormEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_ad_form value
func (v BrandAdGetV30AdForm) Ptr() *BrandAdGetV30AdForm {
	return &v
}
