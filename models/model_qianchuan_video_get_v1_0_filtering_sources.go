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

// QianchuanVideoGetV10FilteringSources
type QianchuanVideoGetV10FilteringSources string

// List of qianchuan_video_get_v1.0_filtering_sources
const (
	AGENT_QianchuanVideoGetV10FilteringSources           QianchuanVideoGetV10FilteringSources = "AGENT"
	ARTHUR_QianchuanVideoGetV10FilteringSources          QianchuanVideoGetV10FilteringSources = "ARTHUR"
	BP_QianchuanVideoGetV10FilteringSources              QianchuanVideoGetV10FilteringSources = "BP"
	CREATIVE_CENTER_QianchuanVideoGetV10FilteringSources QianchuanVideoGetV10FilteringSources = "CREATIVE_CENTER"
	E_COMMERCE_QianchuanVideoGetV10FilteringSources      QianchuanVideoGetV10FilteringSources = "E_COMMERCE"
	JI_CHUANG_QianchuanVideoGetV10FilteringSources       QianchuanVideoGetV10FilteringSources = "JI_CHUANG"
	LIVE_HIGHLIGHT_QianchuanVideoGetV10FilteringSources  QianchuanVideoGetV10FilteringSources = "LIVE_HIGHLIGHT"
	STAR_QianchuanVideoGetV10FilteringSources            QianchuanVideoGetV10FilteringSources = "STAR"
	TADA_QianchuanVideoGetV10FilteringSources            QianchuanVideoGetV10FilteringSources = "TADA"
	VIDEO_CAPTURE_QianchuanVideoGetV10FilteringSources   QianchuanVideoGetV10FilteringSources = "VIDEO_CAPTURE"
)

// All allowed values of QianchuanVideoGetV10FilteringSources enum
var AllowedQianchuanVideoGetV10FilteringSourcesEnumValues = []QianchuanVideoGetV10FilteringSources{
	"AGENT",
	"ARTHUR",
	"BP",
	"CREATIVE_CENTER",
	"E_COMMERCE",
	"JI_CHUANG",
	"LIVE_HIGHLIGHT",
	"STAR",
	"TADA",
	"VIDEO_CAPTURE",
}

// NewQianchuanVideoGetV10FilteringSourcesFromValue returns a pointer to a valid QianchuanVideoGetV10FilteringSources
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanVideoGetV10FilteringSourcesFromValue(v string) (*QianchuanVideoGetV10FilteringSources, error) {
	ev := QianchuanVideoGetV10FilteringSources(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanVideoGetV10FilteringSources: valid values are %v", v, AllowedQianchuanVideoGetV10FilteringSourcesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanVideoGetV10FilteringSources) IsValid() bool {
	for _, existing := range AllowedQianchuanVideoGetV10FilteringSourcesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_video_get_v1.0_filtering_sources value
func (v QianchuanVideoGetV10FilteringSources) Ptr() *QianchuanVideoGetV10FilteringSources {
	return &v
}
