/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselAdGetV2ResponseDataCarouselsInnerImagesInner struct for CarouselAdGetV2ResponseDataCarouselsInnerImagesInner
type CarouselAdGetV2ResponseDataCarouselsInnerImagesInner struct {
	//
	Height *int64 `json:"height,omitempty"`
	//
	ImageId      *string                                                           `json:"image_id,omitempty"`
	ImageSubject *CarouselAdGetV2ResponseDataCarouselsInnerImagesInnerImageSubject `json:"image_subject,omitempty"`
	//
	Ratio *int64 `json:"ratio,omitempty"`
	//
	Size *int64 `json:"size,omitempty"`
	//
	Url map[string]interface{} `json:"url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
