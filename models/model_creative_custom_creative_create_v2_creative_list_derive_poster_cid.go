/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeCustomCreativeCreateV2CreativeListDerivePosterCid
type CreativeCustomCreativeCreateV2CreativeListDerivePosterCid int64

// List of creative_custom_creative_create_v2_creative_list_derive_poster_cid
const (
	Enum_0_CreativeCustomCreativeCreateV2CreativeListDerivePosterCid CreativeCustomCreativeCreateV2CreativeListDerivePosterCid = 0
	Enum_1_CreativeCustomCreativeCreateV2CreativeListDerivePosterCid CreativeCustomCreativeCreateV2CreativeListDerivePosterCid = 1
)

// All allowed values of CreativeCustomCreativeCreateV2CreativeListDerivePosterCid enum
var AllowedCreativeCustomCreativeCreateV2CreativeListDerivePosterCidEnumValues = []CreativeCustomCreativeCreateV2CreativeListDerivePosterCid{
	0,
	1,
}

// NewCreativeCustomCreativeCreateV2CreativeListDerivePosterCidFromValue returns a pointer to a valid CreativeCustomCreativeCreateV2CreativeListDerivePosterCid
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeCustomCreativeCreateV2CreativeListDerivePosterCidFromValue(v int64) (*CreativeCustomCreativeCreateV2CreativeListDerivePosterCid, error) {
	ev := CreativeCustomCreativeCreateV2CreativeListDerivePosterCid(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeCustomCreativeCreateV2CreativeListDerivePosterCid: valid values are %v", v, AllowedCreativeCustomCreativeCreateV2CreativeListDerivePosterCidEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeCustomCreativeCreateV2CreativeListDerivePosterCid) IsValid() bool {
	for _, existing := range AllowedCreativeCustomCreativeCreateV2CreativeListDerivePosterCidEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_custom_creative_create_v2_creative_list_derive_poster_cid value
func (v CreativeCustomCreativeCreateV2CreativeListDerivePosterCid) Ptr() *CreativeCustomCreativeCreateV2CreativeListDerivePosterCid {
	return &v
}
