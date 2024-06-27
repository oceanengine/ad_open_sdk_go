/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode
type BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode int64

// List of brand_creative_get_v3.0_data_creatives_creative_interactive_module_creative_interactive_mode
const (
	Enum_0_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 0
	Enum_1_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 1
	Enum_2_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 2
	Enum_3_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 3
	Enum_4_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 4
	Enum_5_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 5
	Enum_6_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 6
	Enum_7_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 7
	Enum_8_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 8
	Enum_9_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode  BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 9
	Enum_10_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 10
	Enum_11_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 11
	Enum_12_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 12
	Enum_13_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 13
	Enum_14_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 14
	Enum_80_BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode = 80
)

// All allowed values of BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode enum
var AllowedBrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveModeEnumValues = []BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	11,
	12,
	13,
	14,
	80,
}

// NewBrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveModeFromValue returns a pointer to a valid BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveModeFromValue(v int64) (*BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode, error) {
	ev := BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode: valid values are %v", v, AllowedBrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode) IsValid() bool {
	for _, existing := range AllowedBrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_get_v3.0_data_creatives_creative_interactive_module_creative_interactive_mode value
func (v BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode) Ptr() *BrandCreativeGetV30DataCreativesCreativeInteractiveModuleCreativeInteractiveMode {
	return &v
}
