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

// QianchuanReportMaterialGetV10DataListMaterialType
type QianchuanReportMaterialGetV10DataListMaterialType string

// List of qianchuan_report_material_get_v1.0_data_list_material_type
const (
	VIDEO_QianchuanReportMaterialGetV10DataListMaterialType    QianchuanReportMaterialGetV10DataListMaterialType = "video"
	IMAGE_QianchuanReportMaterialGetV10DataListMaterialType    QianchuanReportMaterialGetV10DataListMaterialType = "image"
	CAROUSEL_QianchuanReportMaterialGetV10DataListMaterialType QianchuanReportMaterialGetV10DataListMaterialType = "carousel"
)

// All allowed values of QianchuanReportMaterialGetV10DataListMaterialType enum
var AllowedQianchuanReportMaterialGetV10DataListMaterialTypeEnumValues = []QianchuanReportMaterialGetV10DataListMaterialType{
	"video",
	"image",
	"carousel",
}

// NewQianchuanReportMaterialGetV10DataListMaterialTypeFromValue returns a pointer to a valid QianchuanReportMaterialGetV10DataListMaterialType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportMaterialGetV10DataListMaterialTypeFromValue(v string) (*QianchuanReportMaterialGetV10DataListMaterialType, error) {
	ev := QianchuanReportMaterialGetV10DataListMaterialType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportMaterialGetV10DataListMaterialType: valid values are %v", v, AllowedQianchuanReportMaterialGetV10DataListMaterialTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportMaterialGetV10DataListMaterialType) IsValid() bool {
	for _, existing := range AllowedQianchuanReportMaterialGetV10DataListMaterialTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_material_get_v1.0_data_list_material_type value
func (v QianchuanReportMaterialGetV10DataListMaterialType) Ptr() *QianchuanReportMaterialGetV10DataListMaterialType {
	return &v
}
