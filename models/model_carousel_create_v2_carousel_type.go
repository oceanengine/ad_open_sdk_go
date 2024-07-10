/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CarouselCreateV2CarouselType
type CarouselCreateV2CarouselType string

// List of carousel_create_v2_carousel_type
const (
	INFORMATION_FLOW_IMAGE_CarouselCreateV2CarouselType      CarouselCreateV2CarouselType = "INFORMATION_FLOW_IMAGE"
	SEARCH_DISPLAY_WINDOW_IMAGE_CarouselCreateV2CarouselType CarouselCreateV2CarouselType = "SEARCH_DISPLAY_WINDOW_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CarouselCreateV2CarouselType     CarouselCreateV2CarouselType = "TOUTIAO_SEARCH_AD_IMAGE"
	UNSET_CarouselCreateV2CarouselType                       CarouselCreateV2CarouselType = "UNSET"
)

// All allowed values of CarouselCreateV2CarouselType enum
var AllowedCarouselCreateV2CarouselTypeEnumValues = []CarouselCreateV2CarouselType{
	"INFORMATION_FLOW_IMAGE",
	"SEARCH_DISPLAY_WINDOW_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
	"UNSET",
}

// NewCarouselCreateV2CarouselTypeFromValue returns a pointer to a valid CarouselCreateV2CarouselType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarouselCreateV2CarouselTypeFromValue(v string) (*CarouselCreateV2CarouselType, error) {
	ev := CarouselCreateV2CarouselType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarouselCreateV2CarouselType: valid values are %v", v, AllowedCarouselCreateV2CarouselTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarouselCreateV2CarouselType) IsValid() bool {
	for _, existing := range AllowedCarouselCreateV2CarouselTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to carousel_create_v2_carousel_type value
func (v CarouselCreateV2CarouselType) Ptr() *CarouselCreateV2CarouselType {
	return &v
}
