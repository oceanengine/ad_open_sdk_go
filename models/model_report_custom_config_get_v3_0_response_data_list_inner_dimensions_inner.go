/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomConfigGetV30ResponseDataListInnerDimensionsInner struct for ReportCustomConfigGetV30ResponseDataListInnerDimensionsInner
type ReportCustomConfigGetV30ResponseDataListInnerDimensionsInner struct {
	//
	Description *string `json:"description,omitempty"`
	//
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
	//
	ExclusionMetrics []string `json:"exclusion_metrics,omitempty"`
	//
	Field *string `json:"field,omitempty"`
	//
	FilterAble   *bool                                                                     `json:"filter_able,omitempty"`
	FilterConfig *ReportCustomConfigGetV30ResponseDataListInnerDimensionsInnerFilterConfig `json:"filter_config,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SortAble *bool `json:"sort_able,omitempty"`
}
