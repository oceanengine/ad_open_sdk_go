/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomGetV30ResponseDataRowsInner struct for ReportCustomGetV30ResponseDataRowsInner
type ReportCustomGetV30ResponseDataRowsInner struct {
	// 维度数据
	Dimensions map[string]string `json:"dimensions"`
	// 指标数据
	Metrics map[string]string `json:"metrics"`
}
