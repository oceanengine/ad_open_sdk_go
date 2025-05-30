/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCarouselGetV10ResponseDataCarouselsInner struct for QianchuanCarouselGetV10ResponseDataCarouselsInner
type QianchuanCarouselGetV10ResponseDataCarouselsInner struct {
	Audio *QianchuanCarouselGetV10ResponseDataCarouselsInnerAudio `json:"audio,omitempty"`
	//
	CreateTime *string `json:"create_time,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	Filename  *string                                        `json:"filename,omitempty"`
	ImageMode *QianchuanCarouselGetV10DataCarouselsImageMode `json:"image_mode,omitempty"`
	//
	Images []*QianchuanCarouselGetV10ResponseDataCarouselsInnerImagesInner `json:"images,omitempty"`
	//
	MaterialId *int64                                      `json:"material_id,omitempty"`
	Source     *QianchuanCarouselGetV10DataCarouselsSource `json:"source,omitempty"`
}
