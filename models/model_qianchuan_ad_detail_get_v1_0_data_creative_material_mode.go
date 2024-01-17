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

// QianchuanAdDetailGetV10DataCreativeMaterialMode
type QianchuanAdDetailGetV10DataCreativeMaterialMode string

// List of qianchuan_ad_detail_get_v1.0_data_creative_material_mode
const (
	CUSTOM_CREATIVE_QianchuanAdDetailGetV10DataCreativeMaterialMode       QianchuanAdDetailGetV10DataCreativeMaterialMode = "CUSTOM_CREATIVE"
	PROGRAMMATIC_CREATIVE_QianchuanAdDetailGetV10DataCreativeMaterialMode QianchuanAdDetailGetV10DataCreativeMaterialMode = "PROGRAMMATIC_CREATIVE"
)

// All allowed values of QianchuanAdDetailGetV10DataCreativeMaterialMode enum
var AllowedQianchuanAdDetailGetV10DataCreativeMaterialModeEnumValues = []QianchuanAdDetailGetV10DataCreativeMaterialMode{
	"CUSTOM_CREATIVE",
	"PROGRAMMATIC_CREATIVE",
}

// NewQianchuanAdDetailGetV10DataCreativeMaterialModeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataCreativeMaterialMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataCreativeMaterialModeFromValue(v string) (*QianchuanAdDetailGetV10DataCreativeMaterialMode, error) {
	ev := QianchuanAdDetailGetV10DataCreativeMaterialMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataCreativeMaterialMode: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataCreativeMaterialModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataCreativeMaterialMode) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataCreativeMaterialModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_creative_material_mode value
func (v QianchuanAdDetailGetV10DataCreativeMaterialMode) Ptr() *QianchuanAdDetailGetV10DataCreativeMaterialMode {
	return &v
}
