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

// BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType
type BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType string

// List of brand_creative_create_v3.0_track_url_info_track_url_send_type
const (
	CLIENT_BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType = "CLIENT"
	SERVER_BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType = "SERVER"
)

// All allowed values of BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType enum
var AllowedBrandCreativeCreateV30TrackUrlInfoTrackUrlSendTypeEnumValues = []BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType{
	"CLIENT",
	"SERVER",
}

// NewBrandCreativeCreateV30TrackUrlInfoTrackUrlSendTypeFromValue returns a pointer to a valid BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCreativeCreateV30TrackUrlInfoTrackUrlSendTypeFromValue(v string) (*BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType, error) {
	ev := BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType: valid values are %v", v, AllowedBrandCreativeCreateV30TrackUrlInfoTrackUrlSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType) IsValid() bool {
	for _, existing := range AllowedBrandCreativeCreateV30TrackUrlInfoTrackUrlSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_creative_create_v3.0_track_url_info_track_url_send_type value
func (v BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType) Ptr() *BrandCreativeCreateV30TrackUrlInfoTrackUrlSendType {
	return &v
}
