/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueWechatInstanceUpdateV2Action
type ClueWechatInstanceUpdateV2Action string

// List of clue_wechat_instance_update_v2_action
const (
	ENABLE_ClueWechatInstanceUpdateV2Action  ClueWechatInstanceUpdateV2Action = "ENABLE"
	DISABLE_ClueWechatInstanceUpdateV2Action ClueWechatInstanceUpdateV2Action = "DISABLE"
	DELETE_ClueWechatInstanceUpdateV2Action  ClueWechatInstanceUpdateV2Action = "DELETE"
	ADD_ClueWechatInstanceUpdateV2Action     ClueWechatInstanceUpdateV2Action = "ADD"
)

// All allowed values of ClueWechatInstanceUpdateV2Action enum
var AllowedClueWechatInstanceUpdateV2ActionEnumValues = []ClueWechatInstanceUpdateV2Action{
	"ENABLE",
	"DISABLE",
	"DELETE",
	"ADD",
}

// NewClueWechatInstanceUpdateV2ActionFromValue returns a pointer to a valid ClueWechatInstanceUpdateV2Action
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueWechatInstanceUpdateV2ActionFromValue(v string) (*ClueWechatInstanceUpdateV2Action, error) {
	ev := ClueWechatInstanceUpdateV2Action(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueWechatInstanceUpdateV2Action: valid values are %v", v, AllowedClueWechatInstanceUpdateV2ActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueWechatInstanceUpdateV2Action) IsValid() bool {
	for _, existing := range AllowedClueWechatInstanceUpdateV2ActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_wechat_instance_update_v2_action value
func (v ClueWechatInstanceUpdateV2Action) Ptr() *ClueWechatInstanceUpdateV2Action {
	return &v
}
