/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInnerFilterConfig
type QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInnerFilterConfig struct {
	//
	Operator *int64 `json:"operator,omitempty"`
	//
	RangeValues []*QianchuanReportCustomConfigGetV10ResponseDataCustomConfigDatasInnerDimensionsInnerFilterConfigRangeValuesInner `json:"range_values,omitempty"`
	//
	Type *int64 `json:"type,omitempty"`
	//
	ValueLimit *int64 `json:"value_limit,omitempty"`
}