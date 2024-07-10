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

// BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType
type BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType int64

// List of brand_creative_get_v3.0_data_creatives_creative_external_info_external_url_type
const (
	Enum_0_BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType = 0
	Enum_1_BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType = 1
	Enum_4_BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType = 4
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType enum
var AllowedBrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlTypeEnumValues = []BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType{
	0,
	1,
	4,
}

// NewBrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlTypeFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlTypeFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_external_info_external_url_type value
func (v BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType) Ptr() *BrandCreativeGetV30DataCreativesCreativeExternalInfoExternalUrlType {
	return &v
}
