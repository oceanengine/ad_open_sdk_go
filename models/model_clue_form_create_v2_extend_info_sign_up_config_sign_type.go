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

// ClueFormCreateV2ExtendInfoSignUpConfigSignType
type ClueFormCreateV2ExtendInfoSignUpConfigSignType string

// List of clue_form_create_v2_extend_info_sign_up_config_sign_type
const (
	SIGN_TYPE_SCROLL_WALL_ClueFormCreateV2ExtendInfoSignUpConfigSignType ClueFormCreateV2ExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_WALL"
	SIGN_TYPE_SCROLL_BAR_ClueFormCreateV2ExtendInfoSignUpConfigSignType  ClueFormCreateV2ExtendInfoSignUpConfigSignType = "SIGN_TYPE_SCROLL_BAR"
)

// All allowed values of ClueFormCreateV2ExtendInfoSignUpConfigSignType enum
var AllowedClueFormCreateV2ExtendInfoSignUpConfigSignTypeEnumValues = []ClueFormCreateV2ExtendInfoSignUpConfigSignType{
	"SIGN_TYPE_SCROLL_WALL",
	"SIGN_TYPE_SCROLL_BAR",
}

// NewClueFormCreateV2ExtendInfoSignUpConfigSignTypeFromValue returns a pointer to a valid ClueFormCreateV2ExtendInfoSignUpConfigSignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ExtendInfoSignUpConfigSignTypeFromValue(v string) (*ClueFormCreateV2ExtendInfoSignUpConfigSignType, error) {
	ev := ClueFormCreateV2ExtendInfoSignUpConfigSignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ExtendInfoSignUpConfigSignType: valid values are %v", v, AllowedClueFormCreateV2ExtendInfoSignUpConfigSignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ExtendInfoSignUpConfigSignType) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ExtendInfoSignUpConfigSignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_extend_info_sign_up_config_sign_type value
func (v ClueFormCreateV2ExtendInfoSignUpConfigSignType) Ptr() *ClueFormCreateV2ExtendInfoSignUpConfigSignType {
	return &v
}
