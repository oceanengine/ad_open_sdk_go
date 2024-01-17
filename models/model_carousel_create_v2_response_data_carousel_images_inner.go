/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CarouselCreateV2ResponseDataCarouselImagesInner struct for CarouselCreateV2ResponseDataCarouselImagesInner
type CarouselCreateV2ResponseDataCarouselImagesInner struct {
	//
	Height *int64 `json:"height,omitempty"`
	//
	ImageId      *string                                                      `json:"image_id,omitempty"`
	ImageSubject *CarouselCreateV2ResponseDataCarouselImagesInnerImageSubject `json:"image_subject,omitempty"`
	//
	Ratio *float64 `json:"ratio,omitempty"`
	//
	Size *int64 `json:"size,omitempty"`
	//
	Url *string `json:"url,omitempty"`
	//
	Width *int64 `json:"width,omitempty"`
}
