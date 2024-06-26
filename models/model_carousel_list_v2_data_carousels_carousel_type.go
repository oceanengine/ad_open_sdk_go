/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CarouselListV2DataCarouselsCarouselType
type CarouselListV2DataCarouselsCarouselType string

// List of carousel_list_v2_data_carousels_carousel_type
const (
	INFORMATION_FLOW_IMAGE_CarouselListV2DataCarouselsCarouselType      CarouselListV2DataCarouselsCarouselType = "INFORMATION_FLOW_IMAGE"
	SEARCH_DISPLAY_WINDOW_IMAGE_CarouselListV2DataCarouselsCarouselType CarouselListV2DataCarouselsCarouselType = "SEARCH_DISPLAY_WINDOW_IMAGE"
	TOUTIAO_SEARCH_AD_IMAGE_CarouselListV2DataCarouselsCarouselType     CarouselListV2DataCarouselsCarouselType = "TOUTIAO_SEARCH_AD_IMAGE"
)

// All allowed values of CarouselListV2DataCarouselsCarouselType enum
var AllowedCarouselListV2DataCarouselsCarouselTypeEnumValues = []CarouselListV2DataCarouselsCarouselType{
	"INFORMATION_FLOW_IMAGE",
	"SEARCH_DISPLAY_WINDOW_IMAGE",
	"TOUTIAO_SEARCH_AD_IMAGE",
}

// NewCarouselListV2DataCarouselsCarouselTypeFromValue returns a pointer to a valid CarouselListV2DataCarouselsCarouselType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCarouselListV2DataCarouselsCarouselTypeFromValue(v string) (*CarouselListV2DataCarouselsCarouselType, error) {
	ev := CarouselListV2DataCarouselsCarouselType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CarouselListV2DataCarouselsCarouselType: valid values are %v", v, AllowedCarouselListV2DataCarouselsCarouselTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CarouselListV2DataCarouselsCarouselType) IsValid() bool {
	for _, existing := range AllowedCarouselListV2DataCarouselsCarouselTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to carousel_list_v2_data_carousels_carousel_type value
func (v CarouselListV2DataCarouselsCarouselType) Ptr() *CarouselListV2DataCarouselsCarouselType {
	return &v
}
