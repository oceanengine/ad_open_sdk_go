/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandQueryStockV30LandingType
type BrandQueryStockV30LandingType string

// List of brand_query_stock_v3.0_landing_type
const (
	BRAND_CONTENT_BrandQueryStockV30LandingType BrandQueryStockV30LandingType = "BRAND_CONTENT"
)

// All allowed values of BrandQueryStockV30LandingType enum
var AllowedBrandQueryStockV30LandingTypeEnumValues = []BrandQueryStockV30LandingType{
	"BRAND_CONTENT",
}

// NewBrandQueryStockV30LandingTypeFromValue returns a pointer to a valid BrandQueryStockV30LandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30LandingTypeFromValue(v string) (*BrandQueryStockV30LandingType, error) {
	ev := BrandQueryStockV30LandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30LandingType: valid values are %v", v, AllowedBrandQueryStockV30LandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30LandingType) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30LandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_landing_type value
func (v BrandQueryStockV30LandingType) Ptr() *BrandQueryStockV30LandingType {
	return &v
}
