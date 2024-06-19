/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative
type QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative int64

// List of qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_dynamic_creative
const (
	Enum_0_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative = 0
	Enum_1_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative = 1
)

// All allowed values of QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative enum
var AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreativeEnumValues = []QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreativeFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreativeFromValue(v int64) (*QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative, error) {
	ev := QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreativeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreativeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_dynamic_creative value
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative) Ptr() *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative {
	return &v
}
