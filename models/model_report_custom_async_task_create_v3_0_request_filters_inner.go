/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportCustomAsyncTaskCreateV30RequestFiltersInner struct for ReportCustomAsyncTaskCreateV30RequestFiltersInner
type ReportCustomAsyncTaskCreateV30RequestFiltersInner struct {
	//
	Field string `json:"field"`
	//
	Operator int64 `json:"operator"`
	//
	Type int64 `json:"type"`
	//
	Values []string `json:"values"`
}
