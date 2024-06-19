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

// QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension
type QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension string

// List of qianchuan_report_uni_promotion_dimension_data_room_get_v1.0_dimension
const (
	TIME_GRANULARITY_DAILY_QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension  QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension = "TIME_GRANULARITY_DAILY"
	TIME_GRANULARITY_HOURLY_QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension = "TIME_GRANULARITY_HOURLY"
)

// All allowed values of QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension enum
var AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10DimensionEnumValues = []QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension{
	"TIME_GRANULARITY_DAILY",
	"TIME_GRANULARITY_HOURLY",
}

// NewQianchuanReportUniPromotionDimensionDataRoomGetV10DimensionFromValue returns a pointer to a valid QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanReportUniPromotionDimensionDataRoomGetV10DimensionFromValue(v string) (*QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension, error) {
	ev := QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension: valid values are %v", v, AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10DimensionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension) IsValid() bool {
	for _, existing := range AllowedQianchuanReportUniPromotionDimensionDataRoomGetV10DimensionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_report_uni_promotion_dimension_data_room_get_v1.0_dimension value
func (v QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension) Ptr() *QianchuanReportUniPromotionDimensionDataRoomGetV10Dimension {
	return &v
}
