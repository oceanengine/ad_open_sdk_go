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

// ClueCouponDetailV2DataFormEnableLayer
type ClueCouponDetailV2DataFormEnableLayer string

// List of clue_coupon_detail_v2_data_form_enable_layer
const (
	Enum_0_ClueCouponDetailV2DataFormEnableLayer ClueCouponDetailV2DataFormEnableLayer = "0"
	Enum_1_ClueCouponDetailV2DataFormEnableLayer ClueCouponDetailV2DataFormEnableLayer = "1"
)

// All allowed values of ClueCouponDetailV2DataFormEnableLayer enum
var AllowedClueCouponDetailV2DataFormEnableLayerEnumValues = []ClueCouponDetailV2DataFormEnableLayer{
	"0",
	"1",
}

// NewClueCouponDetailV2DataFormEnableLayerFromValue returns a pointer to a valid ClueCouponDetailV2DataFormEnableLayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormEnableLayerFromValue(v string) (*ClueCouponDetailV2DataFormEnableLayer, error) {
	ev := ClueCouponDetailV2DataFormEnableLayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormEnableLayer: valid values are %v", v, AllowedClueCouponDetailV2DataFormEnableLayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormEnableLayer) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormEnableLayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_enable_layer value
func (v ClueCouponDetailV2DataFormEnableLayer) Ptr() *ClueCouponDetailV2DataFormEnableLayer {
	return &v
}
