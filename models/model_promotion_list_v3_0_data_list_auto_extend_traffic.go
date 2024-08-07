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

// PromotionListV30DataListAutoExtendTraffic
type PromotionListV30DataListAutoExtendTraffic string

// List of promotion_list_v3.0_data_list_auto_extend_traffic
const (
	OFF_PromotionListV30DataListAutoExtendTraffic PromotionListV30DataListAutoExtendTraffic = "OFF"
	ON_PromotionListV30DataListAutoExtendTraffic  PromotionListV30DataListAutoExtendTraffic = "ON"
)

// All allowed values of PromotionListV30DataListAutoExtendTraffic enum
var AllowedPromotionListV30DataListAutoExtendTrafficEnumValues = []PromotionListV30DataListAutoExtendTraffic{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListAutoExtendTrafficFromValue returns a pointer to a valid PromotionListV30DataListAutoExtendTraffic
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListAutoExtendTrafficFromValue(v string) (*PromotionListV30DataListAutoExtendTraffic, error) {
	ev := PromotionListV30DataListAutoExtendTraffic(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListAutoExtendTraffic: valid values are %v", v, AllowedPromotionListV30DataListAutoExtendTrafficEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListAutoExtendTraffic) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListAutoExtendTrafficEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_auto_extend_traffic value
func (v PromotionListV30DataListAutoExtendTraffic) Ptr() *PromotionListV30DataListAutoExtendTraffic {
	return &v
}
