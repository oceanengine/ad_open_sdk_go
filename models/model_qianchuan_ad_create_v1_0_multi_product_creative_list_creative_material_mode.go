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

// QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode
type QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode string

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode       QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeMaterialModeEnumValues = []QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanAdCreateV10MultiProductCreativeListCreativeMaterialModeFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListCreativeMaterialModeFromValue(v string) (*QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_creative_material_mode value
func (v QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode) Ptr() *QianchuanAdCreateV10MultiProductCreativeListCreativeMaterialMode {
	return &v
}
