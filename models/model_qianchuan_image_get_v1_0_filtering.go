/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanImageGetV10Filtering
type QianchuanImageGetV10Filtering struct {
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	ImageIds []string `json:"image_ids,omitempty"`
	//
	ImageMode []*QianchuanImageGetV10FilteringImageMode `json:"image_mode,omitempty"`
	//
	MaterialIds []int64 `json:"material_ids,omitempty"`
	//
	Signatures []string `json:"signatures,omitempty"`
	//
	Sources []*QianchuanImageGetV10FilteringSources `json:"sources,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
}
