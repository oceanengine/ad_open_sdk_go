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

// BrandAdGetV30AppOrigin
type BrandAdGetV30AppOrigin int64

// List of brand_ad_get_v3.0_app_origin
const (
	Enum_1_BrandAdGetV30AppOrigin BrandAdGetV30AppOrigin = 1
	Enum_2_BrandAdGetV30AppOrigin BrandAdGetV30AppOrigin = 2
	Enum_3_BrandAdGetV30AppOrigin BrandAdGetV30AppOrigin = 3
)

// All allowed values of BrandAdGetV30AppOrigin enum
var AllowedBrandAdGetV30AppOriginEnumValues = []BrandAdGetV30AppOrigin{
	1,
	2,
	3,
}

// NewBrandAdGetV30AppOriginFromValue returns a pointer to a valid BrandAdGetV30AppOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAdGetV30AppOriginFromValue(v int64) (*BrandAdGetV30AppOrigin, error) {
	ev := BrandAdGetV30AppOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAdGetV30AppOrigin: valid values are %v", v, AllowedBrandAdGetV30AppOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAdGetV30AppOrigin) IsValid() bool {
	for _, existing := range AllowedBrandAdGetV30AppOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_ad_get_v3.0_app_origin value
func (v BrandAdGetV30AppOrigin) Ptr() *BrandAdGetV30AppOrigin {
	return &v
}
