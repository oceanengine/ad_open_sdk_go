/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanReportMaterialGetV10DataListCarouselSource
type QianchuanReportMaterialGetV10DataListCarouselSource string

// List of qianchuan_report_material_get_v1.0_data_list_carousel_source
const (
	AWEME_QianchuanReportMaterialGetV10DataListCarouselSource      QianchuanReportMaterialGetV10DataListCarouselSource = "AWEME"
	E_COMMERCE_QianchuanReportMaterialGetV10DataListCarouselSource QianchuanReportMaterialGetV10DataListCarouselSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10DataListCarouselSource  QianchuanReportMaterialGetV10DataListCarouselSource = "JI_CHUANG"
)

// All allowed values of QianchuanReportMaterialGetV10DataListCarouselSource enum
var AllowedQianchuanReportMaterialGetV10DataListCarouselSourceEnumValues = []QianchuanReportMaterialGetV10DataListCarouselSource{
	"AWEME",
	"E_COMMERCE",
	"JI_CHUANG",
}

// NewQianchuanReportMaterialGetV10DataListCarouselSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10DataListCarouselSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10DataListCarouselSourceFromValue(v string) (*QianchuanReportMaterialGetV10DataListCarouselSource, error) {
	ev := QianchuanReportMaterialGetV10DataListCarouselSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10DataListCarouselSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10DataListCarouselSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10DataListCarouselSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10DataListCarouselSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_data_list_carousel_source value
func (v QianchuanReportMaterialGetV10DataListCarouselSource) Ptr() *QianchuanReportMaterialGetV10DataListCarouselSource {
	return &v
}