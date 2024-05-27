/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanEstimateEffectV10DeepBidType
type QianchuanEstimateEffectV10DeepBidType string

// List of qianchuan_estimate_effect_v1.0_deep_bid_type
const (
	MIN_QianchuanEstimateEffectV10DeepBidType QianchuanEstimateEffectV10DeepBidType = "MIN"
)

// All allowed values of QianchuanEstimateEffectV10DeepBidType enum
var AllowedQianchuanEstimateEffectV10DeepBidTypeEnumValues = []QianchuanEstimateEffectV10DeepBidType{
	"MIN",
}

// NewQianchuanEstimateEffectV10DeepBidTypeFromValue returns a pointer to a valid QianchuanEstimateEffectV10DeepBidType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanEstimateEffectV10DeepBidTypeFromValue(v string) (*QianchuanEstimateEffectV10DeepBidType, error) {
	ev := QianchuanEstimateEffectV10DeepBidType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanEstimateEffectV10DeepBidType: valid values are %v", v, AllowedQianchuanEstimateEffectV10DeepBidTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanEstimateEffectV10DeepBidType) IsValid() bool {
	for _, existing := range AllowedQianchuanEstimateEffectV10DeepBidTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_estimate_effect_v1.0_deep_bid_type value
func (v QianchuanEstimateEffectV10DeepBidType) Ptr() *QianchuanEstimateEffectV10DeepBidType {
	return &v
}
