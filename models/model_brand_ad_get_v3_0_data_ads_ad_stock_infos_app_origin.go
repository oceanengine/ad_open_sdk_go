/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAdGetV30DataAdsAdStockInfosAppOrigin
type BrandAdGetV30DataAdsAdStockInfosAppOrigin int64

// List of brand_ad_get_v3.0_data_ads_ad_stock_infos_app_origin
const (
	Enum_1_BrandAdGetV30DataAdsAdStockInfosAppOrigin BrandAdGetV30DataAdsAdStockInfosAppOrigin = 1
	Enum_2_BrandAdGetV30DataAdsAdStockInfosAppOrigin BrandAdGetV30DataAdsAdStockInfosAppOrigin = 2
	Enum_3_BrandAdGetV30DataAdsAdStockInfosAppOrigin BrandAdGetV30DataAdsAdStockInfosAppOrigin = 3
)

// All allowed values of BrandAdGetV30DataAdsAdStockInfosAppOrigin enum
var AllowedBrandAdGetV30DataAdsAdStockInfosAppOriginEnumValues = []BrandAdGetV30DataAdsAdStockInfosAppOrigin{
	1,
	2,
	3,
}

// NewBrandAdGetV30DataAdsAdStockInfosAppOriginFromValue returns a pointer to a valid BrandAdGetV30DataAdsAdStockInfosAppOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30DataAdsAdStockInfosAppOriginFromValue(v int64) (*BrandAdGetV30DataAdsAdStockInfosAppOrigin, error) {
	ev := BrandAdGetV30DataAdsAdStockInfosAppOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30DataAdsAdStockInfosAppOrigin: valid values are %v", v, AllowedBrandAdGetV30DataAdsAdStockInfosAppOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30DataAdsAdStockInfosAppOrigin) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30DataAdsAdStockInfosAppOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_data_ads_ad_stock_infos_app_origin value
func (v BrandAdGetV30DataAdsAdStockInfosAppOrigin) Ptr() *BrandAdGetV30DataAdsAdStockInfosAppOrigin {
	return &v
}
