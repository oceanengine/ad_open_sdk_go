/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate
type QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate int64

// List of qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_creative_auto_generate
const (
	Enum_0_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate = 0
	Enum_1_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate = 1
)

// All allowed values of QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate enum
var AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues = []QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerateFromValue(v int64) (*QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate, error) {
	ev := QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_creative_auto_generate value
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate) Ptr() *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate {
	return &v
}