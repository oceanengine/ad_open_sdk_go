/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignUpdateV10BudgetMode
type QianchuanCampaignUpdateV10BudgetMode string

// List of qianchuan_campaign_update_v1.0_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanCampaignUpdateV10BudgetMode      QianchuanCampaignUpdateV10BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanCampaignUpdateV10BudgetMode QianchuanCampaignUpdateV10BudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of QianchuanCampaignUpdateV10BudgetMode enum
var AllowedQianchuanCampaignUpdateV10BudgetModeEnumValues = []QianchuanCampaignUpdateV10BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
}

// NewQianchuanCampaignUpdateV10BudgetModeFromValue returns a pointer to a valid QianchuanCampaignUpdateV10BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignUpdateV10BudgetModeFromValue(v string) (*QianchuanCampaignUpdateV10BudgetMode, error) {
	ev := QianchuanCampaignUpdateV10BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignUpdateV10BudgetMode: valid values are %v", v, AllowedQianchuanCampaignUpdateV10BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignUpdateV10BudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignUpdateV10BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_update_v1.0_budget_mode value
func (v QianchuanCampaignUpdateV10BudgetMode) Ptr() *QianchuanCampaignUpdateV10BudgetMode {
	return &v
}
