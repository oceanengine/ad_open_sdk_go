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

// QianchuanAdCreateV10CreativeMaterialMode
type QianchuanAdCreateV10CreativeMaterialMode string

// List of qianchuan_ad_create_v1.0_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanAdCreateV10CreativeMaterialMode       QianchuanAdCreateV10CreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanAdCreateV10CreativeMaterialMode QianchuanAdCreateV10CreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanAdCreateV10CreativeMaterialMode enum
var AllowedQianchuanAdCreateV10CreativeMaterialModeEnumValues = []QianchuanAdCreateV10CreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanAdCreateV10CreativeMaterialModeFromValue returns a pointer to a valid QianchuanAdCreateV10CreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10CreativeMaterialModeFromValue(v string) (*QianchuanAdCreateV10CreativeMaterialMode, error) {
	ev := QianchuanAdCreateV10CreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10CreativeMaterialMode: valid values are %v", v, AllowedQianchuanAdCreateV10CreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10CreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10CreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_creative_material_mode value
func (v QianchuanAdCreateV10CreativeMaterialMode) Ptr() *QianchuanAdCreateV10CreativeMaterialMode {
	return &v
}
