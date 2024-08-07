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

// PromotionListV30FilteringRejectReasonType
type PromotionListV30FilteringRejectReasonType string

// List of promotion_list_v3.0_filtering_reject_reason_type
const (
	DISCOMFORT_PromotionListV30FilteringRejectReasonType    PromotionListV30FilteringRejectReasonType = "DISCOMFORT"
	ELSE_PromotionListV30FilteringRejectReasonType          PromotionListV30FilteringRejectReasonType = "ELSE"
	EXAGGERATION_PromotionListV30FilteringRejectReasonType  PromotionListV30FilteringRejectReasonType = "EXAGGERATION"
	LOW_MATERAIL_PromotionListV30FilteringRejectReasonType  PromotionListV30FilteringRejectReasonType = "LOW_MATERAIL"
	NONE_PromotionListV30FilteringRejectReasonType          PromotionListV30FilteringRejectReasonType = "NONE"
	QUALITY_POOR_PromotionListV30FilteringRejectReasonType  PromotionListV30FilteringRejectReasonType = "QUALITY_POOR"
	REVIEW_REJECT_PromotionListV30FilteringRejectReasonType PromotionListV30FilteringRejectReasonType = "REVIEW_REJECT"
)

// All allowed values of PromotionListV30FilteringRejectReasonType enum
var AllowedPromotionListV30FilteringRejectReasonTypeEnumValues = []PromotionListV30FilteringRejectReasonType{
	"DISCOMFORT",
	"ELSE",
	"EXAGGERATION",
	"LOW_MATERAIL",
	"NONE",
	"QUALITY_POOR",
	"REVIEW_REJECT",
}

// NewPromotionListV30FilteringRejectReasonTypeFromValue returns a pointer to a valid PromotionListV30FilteringRejectReasonType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30FilteringRejectReasonTypeFromValue(v string) (*PromotionListV30FilteringRejectReasonType, error) {
	ev := PromotionListV30FilteringRejectReasonType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30FilteringRejectReasonType: valid values are %v", v, AllowedPromotionListV30FilteringRejectReasonTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30FilteringRejectReasonType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30FilteringRejectReasonTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_filtering_reject_reason_type value
func (v PromotionListV30FilteringRejectReasonType) Ptr() *PromotionListV30FilteringRejectReasonType {
	return &v
}
