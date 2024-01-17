/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanVideoGetV10ResponseDataListInner struct for QianchuanVideoGetV10ResponseDataListInner
type QianchuanVideoGetV10ResponseDataListInner struct {
	//
	BitRate *int64 `json:"bit_rate,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Duration *float64 `json:"duration,omitempty"`
	//
	Filename *string `json:"filename,omitempty"`
	//
	Format *string `json:"format,omitempty"`
	//
	Height *int64 `json:"height,omitempty"`
	//
	Id        *string                                `json:"id,omitempty"`
	ImageMode *QianchuanVideoGetV10DataListImageMode `json:"image_mode,omitempty"`
	//
	IsAiCreate *bool `json:"is_ai_create,omitempty"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	PosterUrl *string `json:"poster_url,omitempty"`
	//
	Signature *string `json:"signature,omitempty"`
	//
	Size   *int64                              `json:"size,omitempty"`
	Source *QianchuanVideoGetV10DataListSource `json:"source,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
