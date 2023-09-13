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

// ReportRubeexGetV2FilteringPlayableType
type ReportRubeexGetV2FilteringPlayableType int64

// List of report_rubeex_get_v2_filtering_playable_type
const (
	Enum_1_ReportRubeexGetV2FilteringPlayableType ReportRubeexGetV2FilteringPlayableType = 1
	Enum_2_ReportRubeexGetV2FilteringPlayableType ReportRubeexGetV2FilteringPlayableType = 2
)

// All allowed values of ReportRubeexGetV2FilteringPlayableType enum
var AllowedReportRubeexGetV2FilteringPlayableTypeEnumValues = []ReportRubeexGetV2FilteringPlayableType{
	1,
	2,
}

// NewReportRubeexGetV2FilteringPlayableTypeFromValue returns a pointer to a valid ReportRubeexGetV2FilteringPlayableType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportRubeexGetV2FilteringPlayableTypeFromValue(v int64) (*ReportRubeexGetV2FilteringPlayableType, error) {
	ev := ReportRubeexGetV2FilteringPlayableType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportRubeexGetV2FilteringPlayableType: valid values are %v", v, AllowedReportRubeexGetV2FilteringPlayableTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportRubeexGetV2FilteringPlayableType) IsValid() bool {
	for _, existing := range AllowedReportRubeexGetV2FilteringPlayableTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_rubeex_get_v2_filtering_playable_type value
func (v ReportRubeexGetV2FilteringPlayableType) Ptr() *ReportRubeexGetV2FilteringPlayableType {
	return &v
}
