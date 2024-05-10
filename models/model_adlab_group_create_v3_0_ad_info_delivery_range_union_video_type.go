/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType
type AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType string

// List of adlab_group_create_v3.0_ad_info_delivery_range_union_video_type
const (
	ORIGIN_VIDEO_AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType = "ORIGIN_VIDEO"
	REWARD_VIDEO_AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType = "REWARD_VIDEO"
)

// All allowed values of AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType enum
var AllowedAdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoTypeEnumValues = []AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType{
	"ORIGIN_VIDEO",
	"REWARD_VIDEO",
}

// NewAdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoTypeFromValue returns a pointer to a valid AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoTypeFromValue(v string) (*AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType, error) {
	ev := AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType: valid values are %v", v, AllowedAdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType) IsValid() bool {
	for _, existing := range AllowedAdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_create_v3.0_ad_info_delivery_range_union_video_type value
func (v AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType) Ptr() *AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType {
	return &v
}
