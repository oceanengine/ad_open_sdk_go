/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanVideoGetV10Filtering
type QianchuanVideoGetV10Filtering struct {
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	ImageMode []*QianchuanVideoGetV10FilteringImageMode `json:"image_mode,omitempty"`
	//
	MaterialIds []int64 `json:"material_ids,omitempty"`
	//
	Signatures []string `json:"signatures,omitempty"`
	//
	Sources []*QianchuanVideoGetV10FilteringSources `json:"sources,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	VideoIds []string `json:"video_ids,omitempty"`
}
