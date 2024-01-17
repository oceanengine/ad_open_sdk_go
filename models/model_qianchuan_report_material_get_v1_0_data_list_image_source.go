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

// QianchuanReportMaterialGetV10DataListImageSource
type QianchuanReportMaterialGetV10DataListImageSource string

// List of qianchuan_report_material_get_v1.0_data_list_image_source
const (
	CREATIVE_CENTER_QianchuanReportMaterialGetV10DataListImageSource QianchuanReportMaterialGetV10DataListImageSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanReportMaterialGetV10DataListImageSource      QianchuanReportMaterialGetV10DataListImageSource = "E_COMMERCE"
	JI_CHUANG_QianchuanReportMaterialGetV10DataListImageSource       QianchuanReportMaterialGetV10DataListImageSource = "JI_CHUANG"
	SQUARE_QianchuanReportMaterialGetV10DataListImageSource          QianchuanReportMaterialGetV10DataListImageSource = "SQUARE"
)

// All allowed values of QianchuanReportMaterialGetV10DataListImageSource enum
var AllowedQianchuanReportMaterialGetV10DataListImageSourceEnumValues = []QianchuanReportMaterialGetV10DataListImageSource{
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
	"SQUARE",
}

// NewQianchuanReportMaterialGetV10DataListImageSourceFromValue returns a pointer to a valid QianchuanReportMaterialGetV10DataListImageSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10DataListImageSourceFromValue(v string) (*QianchuanReportMaterialGetV10DataListImageSource, error) {
	ev := QianchuanReportMaterialGetV10DataListImageSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10DataListImageSource: valid values are %v", v, AllowedQianchuanReportMaterialGetV10DataListImageSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10DataListImageSource) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10DataListImageSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_data_list_image_source value
func (v QianchuanReportMaterialGetV10DataListImageSource) Ptr() *QianchuanReportMaterialGetV10DataListImageSource {
	return &v
}
