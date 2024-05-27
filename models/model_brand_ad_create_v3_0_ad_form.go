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

// BrandAdCreateV30AdForm
type BrandAdCreateV30AdForm string

// List of brand_ad_create_v3.0_ad_form
const (
	CONTENT_RECOMMEND_BrandAdCreateV30AdForm BrandAdCreateV30AdForm = "CONTENT_RECOMMEND"
	CONTENT_SERVICE_BrandAdCreateV30AdForm   BrandAdCreateV30AdForm = "CONTENT_SERVICE"
)

// All allowed values of BrandAdCreateV30AdForm enum
var AllowedBrandAdCreateV30AdFormEnumValues = []BrandAdCreateV30AdForm{
	"CONTENT_RECOMMEND",
	"CONTENT_SERVICE",
}

// NewBrandAdCreateV30AdFormFromValue returns a pointer to a valid BrandAdCreateV30AdForm
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30AdFormFromValue(v string) (*BrandAdCreateV30AdForm, error) {
	ev := BrandAdCreateV30AdForm(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30AdForm: valid values are %v", v, AllowedBrandAdCreateV30AdFormEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30AdForm) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30AdFormEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_ad_form value
func (v BrandAdCreateV30AdForm) Ptr() *BrandAdCreateV30AdForm {
	return &v
}
