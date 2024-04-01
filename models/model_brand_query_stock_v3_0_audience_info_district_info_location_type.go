/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandQueryStockV30AudienceInfoDistrictInfoLocationType
type BrandQueryStockV30AudienceInfoDistrictInfoLocationType string

// List of brand_query_stock_v3.0_audience_info_district_info_location_type
const (
	ALL_LOCATION_BrandQueryStockV30AudienceInfoDistrictInfoLocationType BrandQueryStockV30AudienceInfoDistrictInfoLocationType = "ALL_LOCATION"
	TRAVEL_BrandQueryStockV30AudienceInfoDistrictInfoLocationType       BrandQueryStockV30AudienceInfoDistrictInfoLocationType = "TRAVEL"
)

// All allowed values of BrandQueryStockV30AudienceInfoDistrictInfoLocationType enum
var AllowedBrandQueryStockV30AudienceInfoDistrictInfoLocationTypeEnumValues = []BrandQueryStockV30AudienceInfoDistrictInfoLocationType{
	"ALL_LOCATION",
	"TRAVEL",
}

// NewBrandQueryStockV30AudienceInfoDistrictInfoLocationTypeFromValue returns a pointer to a valid BrandQueryStockV30AudienceInfoDistrictInfoLocationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandQueryStockV30AudienceInfoDistrictInfoLocationTypeFromValue(v string) (*BrandQueryStockV30AudienceInfoDistrictInfoLocationType, error) {
	ev := BrandQueryStockV30AudienceInfoDistrictInfoLocationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandQueryStockV30AudienceInfoDistrictInfoLocationType: valid values are %v", v, AllowedBrandQueryStockV30AudienceInfoDistrictInfoLocationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandQueryStockV30AudienceInfoDistrictInfoLocationType) IsValid() bool {
	for _, existing := range AllowedBrandQueryStockV30AudienceInfoDistrictInfoLocationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_query_stock_v3.0_audience_info_district_info_location_type value
func (v BrandQueryStockV30AudienceInfoDistrictInfoLocationType) Ptr() *BrandQueryStockV30AudienceInfoDistrictInfoLocationType {
	return &v
}
