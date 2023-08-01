/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30IsCommentDisable
type PromotionCreateV30IsCommentDisable string

// List of promotion_create_v3.0_is_comment_disable
const (
	OFF_PromotionCreateV30IsCommentDisable PromotionCreateV30IsCommentDisable = "OFF"
	ON_PromotionCreateV30IsCommentDisable  PromotionCreateV30IsCommentDisable = "ON"
)

// All allowed values of PromotionCreateV30IsCommentDisable enum
var AllowedPromotionCreateV30IsCommentDisableEnumValues = []PromotionCreateV30IsCommentDisable{
	"OFF",
	"ON",
}

// NewPromotionCreateV30IsCommentDisableFromValue returns a pointer to a valid PromotionCreateV30IsCommentDisable
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30IsCommentDisableFromValue(v string) (*PromotionCreateV30IsCommentDisable, error) {
	ev := PromotionCreateV30IsCommentDisable(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30IsCommentDisable: valid values are %v", v, AllowedPromotionCreateV30IsCommentDisableEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30IsCommentDisable) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30IsCommentDisableEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_is_comment_disable value
func (v PromotionCreateV30IsCommentDisable) Ptr() *PromotionCreateV30IsCommentDisable {
	return &v
}
