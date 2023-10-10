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

// AdvertiserBudgetGetV2DataListBudgetMode
type AdvertiserBudgetGetV2DataListBudgetMode string

// List of advertiser_budget_get_v2_data_list_budget_mode
const (
	BUDGET_MODE_DAY_AdvertiserBudgetGetV2DataListBudgetMode      AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_AdvertiserBudgetGetV2DataListBudgetMode AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_AdvertiserBudgetGetV2DataListBudgetMode    AdvertiserBudgetGetV2DataListBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of AdvertiserBudgetGetV2DataListBudgetMode enum
var AllowedAdvertiserBudgetGetV2DataListBudgetModeEnumValues = []AdvertiserBudgetGetV2DataListBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewAdvertiserBudgetGetV2DataListBudgetModeFromValue returns a pointer to a valid AdvertiserBudgetGetV2DataListBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdvertiserBudgetGetV2DataListBudgetModeFromValue(v string) (*AdvertiserBudgetGetV2DataListBudgetMode, error) {
	ev := AdvertiserBudgetGetV2DataListBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdvertiserBudgetGetV2DataListBudgetMode: valid values are %v", v, AllowedAdvertiserBudgetGetV2DataListBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdvertiserBudgetGetV2DataListBudgetMode) IsValid() bool {
	for _, existing := range AllowedAdvertiserBudgetGetV2DataListBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to advertiser_budget_get_v2_data_list_budget_mode value
func (v AdvertiserBudgetGetV2DataListBudgetMode) Ptr() *AdvertiserBudgetGetV2DataListBudgetMode {
	return &v
}
