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

// QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType
type QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType string

// List of qianchuan_ad_update_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_title_list_title_type
const (
	AWEME_CAROUSEL_QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "AWEME_CAROUSEL"
	COMMODITY_CARD_QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType         QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "CUSTOM"
)

// All allowed values of QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType enum
var AllowedQianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues = []QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType{
	"AWEME_CAROUSEL",
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeFromValue returns a pointer to a valid QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeFromValue(v string) (*QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType, error) {
	ev := QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType: valid values are %v", v, AllowedQianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_update_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_title_list_title_type value
func (v QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType) Ptr() *QianchuanAdUpdateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType {
	return &v
}
