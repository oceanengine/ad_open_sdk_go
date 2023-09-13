/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType
type QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType string

// List of qianchuan_ad_create_v1.0_programmatic_creative_title_list_title_type
const (
	COMMODITY_CARD_QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType = "COMMODITY_CARD"
	CUSTOM_QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType         QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType = "CUSTOM"
)

// All allowed values of QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType enum
var AllowedQianchuanAdCreateV10ProgrammaticCreativeTitleListTitleTypeEnumValues = []QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType{
	"COMMODITY_CARD",
	"CUSTOM",
}

// NewQianchuanAdCreateV10ProgrammaticCreativeTitleListTitleTypeFromValue returns a pointer to a valid QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAdCreateV10ProgrammaticCreativeTitleListTitleTypeFromValue(v string) (*QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType, error) {
	ev := QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType: valid values are %v", v, AllowedQianchuanAdCreateV10ProgrammaticCreativeTitleListTitleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType) IsValid() bool {
	for _, existing := range AllowedQianchuanAdCreateV10ProgrammaticCreativeTitleListTitleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_ad_create_v1.0_programmatic_creative_title_list_title_type value
func (v QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType) Ptr() *QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType {
	return &v
}
