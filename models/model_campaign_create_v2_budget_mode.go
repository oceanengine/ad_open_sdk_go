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

// CampaignCreateV2BudgetMode
type CampaignCreateV2BudgetMode string

// List of campaign_create_v2_budget_mode
const (
	BUDGET_MODE_DAY_CampaignCreateV2BudgetMode      CampaignCreateV2BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_CampaignCreateV2BudgetMode CampaignCreateV2BudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_CampaignCreateV2BudgetMode    CampaignCreateV2BudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of CampaignCreateV2BudgetMode enum
var AllowedCampaignCreateV2BudgetModeEnumValues = []CampaignCreateV2BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewCampaignCreateV2BudgetModeFromValue returns a pointer to a valid CampaignCreateV2BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignCreateV2BudgetModeFromValue(v string) (*CampaignCreateV2BudgetMode, error) {
	ev := CampaignCreateV2BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignCreateV2BudgetMode: valid values are %v", v, AllowedCampaignCreateV2BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignCreateV2BudgetMode) IsValid() bool {
	for _, existing := range AllowedCampaignCreateV2BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_create_v2_budget_mode value
func (v CampaignCreateV2BudgetMode) Ptr() *CampaignCreateV2BudgetMode {
	return &v
}
