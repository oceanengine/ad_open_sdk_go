/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderCreateV30ObjectInfosMarketingGoal
type DouplusOrderCreateV30ObjectInfosMarketingGoal string

// List of douplus_order_create_v3.0_object_infos_marketing_goal
const (
	VIDEO_PROM_GOODS_DouplusOrderCreateV30ObjectInfosMarketingGoal DouplusOrderCreateV30ObjectInfosMarketingGoal = "VIDEO_PROM_GOODS"
	LIVE_PROM_GOODS_DouplusOrderCreateV30ObjectInfosMarketingGoal  DouplusOrderCreateV30ObjectInfosMarketingGoal = "LIVE_PROM_GOODS"
)

// All allowed values of DouplusOrderCreateV30ObjectInfosMarketingGoal enum
var AllowedDouplusOrderCreateV30ObjectInfosMarketingGoalEnumValues = []DouplusOrderCreateV30ObjectInfosMarketingGoal{
	"VIDEO_PROM_GOODS",
	"LIVE_PROM_GOODS",
}

// NewDouplusOrderCreateV30ObjectInfosMarketingGoalFromValue returns a pointer to a valid DouplusOrderCreateV30ObjectInfosMarketingGoal
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderCreateV30ObjectInfosMarketingGoalFromValue(v string) (*DouplusOrderCreateV30ObjectInfosMarketingGoal, error) {
	ev := DouplusOrderCreateV30ObjectInfosMarketingGoal(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderCreateV30ObjectInfosMarketingGoal: valid values are %v", v, AllowedDouplusOrderCreateV30ObjectInfosMarketingGoalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderCreateV30ObjectInfosMarketingGoal) IsValid() bool {
	for _, existing := range AllowedDouplusOrderCreateV30ObjectInfosMarketingGoalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_create_v3.0_object_infos_marketing_goal value
func (v DouplusOrderCreateV30ObjectInfosMarketingGoal) Ptr() *DouplusOrderCreateV30ObjectInfosMarketingGoal {
	return &v
}
