/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// SpiTaskGetV2ResponseDataCursorInfo
type SpiTaskGetV2ResponseDataCursorInfo struct {
	//
	Count *int64 `json:"count,omitempty"`
	//
	Cursor *int64 `json:"cursor,omitempty"`
	//
	HasMore *bool `json:"has_more,omitempty"`
}
