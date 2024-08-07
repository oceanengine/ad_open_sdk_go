/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable
type CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable int64

// List of creative_procedural_creative_update_v2_ad_data_is_comment_disable
const (
	Enum_0_CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable = 0
	Enum_1_CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable = 1
)

// All allowed values of CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable enum
var AllowedCreativeProceduralCreativeUpdateV2AdDataIsCommentDisableEnumValues = []CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable{
	0,
	1,
}

// NewCreativeProceduralCreativeUpdateV2AdDataIsCommentDisableFromValue returns a pointer to a valid CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCreativeProceduralCreativeUpdateV2AdDataIsCommentDisableFromValue(v int64) (*CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable, error) {
	ev := CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable: valid values are %v", v, AllowedCreativeProceduralCreativeUpdateV2AdDataIsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable) IsValid() bool {
	for _, existing := range AllowedCreativeProceduralCreativeUpdateV2AdDataIsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to creative_procedural_creative_update_v2_ad_data_is_comment_disable value
func (v CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable) Ptr() *CreativeProceduralCreativeUpdateV2AdDataIsCommentDisable {
	return &v
}
