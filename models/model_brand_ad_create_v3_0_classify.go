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

// BrandAdCreateV30Classify
type BrandAdCreateV30Classify string

// List of brand_ad_create_v3.0_classify
const (
	SALE_BrandAdCreateV30Classify BrandAdCreateV30Classify = "SALE"
)

// All allowed values of BrandAdCreateV30Classify enum
var AllowedBrandAdCreateV30ClassifyEnumValues = []BrandAdCreateV30Classify{
	"SALE",
}

// NewBrandAdCreateV30ClassifyFromValue returns a pointer to a valid BrandAdCreateV30Classify
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdCreateV30ClassifyFromValue(v string) (*BrandAdCreateV30Classify, error) {
	ev := BrandAdCreateV30Classify(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdCreateV30Classify: valid values are %v", v, AllowedBrandAdCreateV30ClassifyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdCreateV30Classify) IsValid() bool {
	for _, existing := range AllowedBrandAdCreateV30ClassifyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_create_v3.0_classify value
func (v BrandAdCreateV30Classify) Ptr() *BrandAdCreateV30Classify {
	return &v
}
