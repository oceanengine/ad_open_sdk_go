/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CarouselListV2FilteringCarouselType
type CarouselListV2FilteringCarouselType string

// List of carousel_list_v2_filtering_carousel_type
const (
	INFORMATION_FLOW_IMAGE_CarouselListV2FilteringCarouselType      CarouselListV2FilteringCarouselType = "INFORMATION_FLOW_IMAGE"
	SEARCH_DISPLAY_WINDOW_IMAGE_CarouselListV2FilteringCarouselType CarouselListV2FilteringCarouselType = "SEARCH_DISPLAY_WINDOW_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CarouselListV2FilteringCarouselType     CarouselListV2FilteringCarouselType = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CarouselListV2FilteringCarouselType enum
var AllowedCarouselListV2FilteringCarouselTypeEnumValues = []CarouselListV2FilteringCarouselType{
	"INFORMATION_FLOW_IMAGE",
	"SEARCH_DISPLAY_WINDOW_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewCarouselListV2FilteringCarouselTypeFromValue returns a pointer to a valid CarouselListV2FilteringCarouselType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarouselListV2FilteringCarouselTypeFromValue(v string) (*CarouselListV2FilteringCarouselType, error) {
	ev := CarouselListV2FilteringCarouselType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarouselListV2FilteringCarouselType: valid values are %v", v, AllowedCarouselListV2FilteringCarouselTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarouselListV2FilteringCarouselType) IsValid() bool {
	for _, existing := range AllowedCarouselListV2FilteringCarouselTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to carousel_list_v2_filtering_carousel_type value
func (v CarouselListV2FilteringCarouselType) Ptr() *CarouselListV2FilteringCarouselType {
	return &v
}
