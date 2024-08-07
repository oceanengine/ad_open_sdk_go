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

// QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType
type QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType string

// List of qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_title_list_title_type
const (
	AWEME_CAROUSEL_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "AWEME_CAROUSEL"
	COMMODITY_CARD_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType         QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType = "CUSTOM"
)

// All allowed values of QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType enum
var AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues = []QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType{
	"AWEME_CAROUSEL",
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeFromValue returns a pointer to a valid QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeFromValue(v string) (*QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType, error) {
	ev := QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType: valid values are %v", v, AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_multi_product_creative_list_programmatic_creative_programmatic_creative_title_list_title_type value
func (v QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType) Ptr() *QianchuanAdCreateV10MultiProductCreativeListProgrammaticCreativeProgrammaticCreativeTitleListTitleType {
	return &v
}
