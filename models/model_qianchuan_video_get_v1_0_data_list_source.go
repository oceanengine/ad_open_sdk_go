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

// QianchuanVideoGetV10DataListSource
type QianchuanVideoGetV10DataListSource string

// List of qianchuan_video_get_v1.0_data_list_source
const (
	ARTHUR_QianchuanVideoGetV10DataListSource          QianchuanVideoGetV10DataListSource = "ARTHUR"
	AUTO_GENERATE_QianchuanVideoGetV10DataListSource   QianchuanVideoGetV10DataListSource = "AUTO_GENERATE"
	BP_QianchuanVideoGetV10DataListSource              QianchuanVideoGetV10DataListSource = "BP"
	CREATIVE_CENTER_QianchuanVideoGetV10DataListSource QianchuanVideoGetV10DataListSource = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanVideoGetV10DataListSource      QianchuanVideoGetV10DataListSource = "E_COMMERCE"
	JI_CHUANG_QianchuanVideoGetV10DataListSource       QianchuanVideoGetV10DataListSource = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanVideoGetV10DataListSource  QianchuanVideoGetV10DataListSource = "LIVE_HIGHLIGHT"
	STAR_QianchuanVideoGetV10DataListSource            QianchuanVideoGetV10DataListSource = "STAR"
	TADA_QianchuanVideoGetV10DataListSource            QianchuanVideoGetV10DataListSource = "TADA"
	VIDEO_CAPTURE_QianchuanVideoGetV10DataListSource   QianchuanVideoGetV10DataListSource = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanVideoGetV10DataListSource enum
var AllowedQianchuanVideoGetV10DataListSourceEnumValues = []QianchuanVideoGetV10DataListSource{
	"ARTHUR",
	"AUTO_GENERATE",
	"BP",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
	"LIVE_HIGHLIGHT",
	"STAR",
	"TADA",
	"VIDEO_CAPTURE",
}

// NewQianchuanVideoGetV10DataListSourceFromValue returns a pointer to a valid QianchuanVideoGetV10DataListSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanVideoGetV10DataListSourceFromValue(v string) (*QianchuanVideoGetV10DataListSource, error) {
	ev := QianchuanVideoGetV10DataListSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanVideoGetV10DataListSource: valid values are %v", v, AllowedQianchuanVideoGetV10DataListSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanVideoGetV10DataListSource) IsValid() bool {
	for _, existing := range AllowedQianchuanVideoGetV10DataListSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_video_get_v1.0_data_list_source value
func (v QianchuanVideoGetV10DataListSource) Ptr() *QianchuanVideoGetV10DataListSource {
	return &v
}
