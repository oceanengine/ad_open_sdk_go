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

// QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide
type QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide int64

// List of qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_is_homepage_hide
const (
	Enum_0_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide = 0
	Enum_1_QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide = 1
)

// All allowed values of QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide enum
var AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHideEnumValues = []QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide{
	0,
	1,
}

// NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHideFromValue returns a pointer to a valid QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHideFromValue(v int64) (*QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide, error) {
	ev := QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide: valid values are %v", v, AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHideEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide) IsValid() bool {
	for _, existing := range AllowedQianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHideEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_detail_get_v1.0_data_multi_product_creative_list_creative_setting_is_homepage_hide value
func (v QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide) Ptr() *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide {
	return &v
}
