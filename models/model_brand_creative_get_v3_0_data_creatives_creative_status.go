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

// BrandCreativeGetV30DataCreativesCreativeStatus
type BrandCreativeGetV30DataCreativesCreativeStatus int64

// List of brand_creative_get_v3.0_data_creatives_creative_status
const (
	Enum_0_BrandCreativeGetV30DataCreativesCreativeStatus  BrandCreativeGetV30DataCreativesCreativeStatus = 0
	Enum_2_BrandCreativeGetV30DataCreativesCreativeStatus  BrandCreativeGetV30DataCreativesCreativeStatus = 2
	Enum_3_BrandCreativeGetV30DataCreativesCreativeStatus  BrandCreativeGetV30DataCreativesCreativeStatus = 3
	Enum_4_BrandCreativeGetV30DataCreativesCreativeStatus  BrandCreativeGetV30DataCreativesCreativeStatus = 4
	Enum_41_BrandCreativeGetV30DataCreativesCreativeStatus BrandCreativeGetV30DataCreativesCreativeStatus = 41
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeStatus enum
var AllowedBrandCreativeGetV30DataCreativesCreativeStatusEnumValues = []BrandCreativeGetV30DataCreativesCreativeStatus{
	0,
	2,
	3,
	4,
	41,
}

// NewBrandCreativeGetV30DataCreativesCreativeStatusFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeStatusFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeStatus, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeStatus: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeStatus) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_status value
func (v BrandCreativeGetV30DataCreativesCreativeStatus) Ptr() *BrandCreativeGetV30DataCreativesCreativeStatus {
	return &v
}
