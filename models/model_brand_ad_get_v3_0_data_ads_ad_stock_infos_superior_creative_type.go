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

// BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType
type BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType int64

// List of brand_ad_get_v3.0_data_ads_ad_stock_infos_superior_creative_type
const (
	Enum_1_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 1
	Enum_2_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 2
	Enum_3_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 3
	Enum_4_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 4
	Enum_5_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 5
	Enum_6_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 6
	Enum_7_BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType = 7
)

// All allowed values of BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType enum
var AllowedBrandAdGetV30DataAdsAdStockInfosSuperiorCreativeTypeEnumValues = []BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType{
	1,
	2,
	3,
	4,
	5,
	6,
	7,
}

// NewBrandAdGetV30DataAdsAdStockInfosSuperiorCreativeTypeFromValue returns a pointer to a valid BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAdStockInfosSuperiorCreativeTypeFromValue(v int64) (*BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType, error) {
	ev := BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType: valid values are %v", v, AllowedBrandAdGetV30DataAdsAdStockInfosSuperiorCreativeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAdStockInfosSuperiorCreativeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_ad_stock_infos_superior_creative_type value
func (v BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType) Ptr() *BrandAdGetV30DataAdsAdStockInfosSuperiorCreativeType {
	return &v
}
