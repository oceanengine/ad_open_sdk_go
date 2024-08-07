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

// NativeAnchorUpdateV30AnchorInfoPrivateChatButton
type NativeAnchorUpdateV30AnchorInfoPrivateChatButton string

// List of native_anchor_update_v3.0_anchor_info_private_chat_button
const (
	PRIVATE_MESSAGE_NativeAnchorUpdateV30AnchorInfoPrivateChatButton  NativeAnchorUpdateV30AnchorInfoPrivateChatButton = "PRIVATE_MESSAGE"
	CONSULT_NOW_NativeAnchorUpdateV30AnchorInfoPrivateChatButton      NativeAnchorUpdateV30AnchorInfoPrivateChatButton = "CONSULT_NOW"
	CONSULT_ADVISOR_NativeAnchorUpdateV30AnchorInfoPrivateChatButton  NativeAnchorUpdateV30AnchorInfoPrivateChatButton = "CONSULT_ADVISOR"
	CONSULT_DESIGNER_NativeAnchorUpdateV30AnchorInfoPrivateChatButton NativeAnchorUpdateV30AnchorInfoPrivateChatButton = "CONSULT_DESIGNER"
	ASK_TEACHER_NativeAnchorUpdateV30AnchorInfoPrivateChatButton      NativeAnchorUpdateV30AnchorInfoPrivateChatButton = "ASK_TEACHER"
)

// All allowed values of NativeAnchorUpdateV30AnchorInfoPrivateChatButton enum
var AllowedNativeAnchorUpdateV30AnchorInfoPrivateChatButtonEnumValues = []NativeAnchorUpdateV30AnchorInfoPrivateChatButton{
	"PRIVATE_MESSAGE",
	"CONSULT_NOW",
	"CONSULT_ADVISOR",
	"CONSULT_DESIGNER",
	"ASK_TEACHER",
}

// NewNativeAnchorUpdateV30AnchorInfoPrivateChatButtonFromValue returns a pointer to a valid NativeAnchorUpdateV30AnchorInfoPrivateChatButton
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNativeAnchorUpdateV30AnchorInfoPrivateChatButtonFromValue(v string) (*NativeAnchorUpdateV30AnchorInfoPrivateChatButton, error) {
	ev := NativeAnchorUpdateV30AnchorInfoPrivateChatButton(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NativeAnchorUpdateV30AnchorInfoPrivateChatButton: valid values are %v", v, AllowedNativeAnchorUpdateV30AnchorInfoPrivateChatButtonEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NativeAnchorUpdateV30AnchorInfoPrivateChatButton) IsValid() bool {
	for _, existing := range AllowedNativeAnchorUpdateV30AnchorInfoPrivateChatButtonEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to native_anchor_update_v3.0_anchor_info_private_chat_button value
func (v NativeAnchorUpdateV30AnchorInfoPrivateChatButton) Ptr() *NativeAnchorUpdateV30AnchorInfoPrivateChatButton {
	return &v
}
