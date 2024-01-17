/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignCreateV10BudgetMode
type QianchuanCampaignCreateV10BudgetMode string

// List of qianchuan_campaign_create_v1.0_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanCampaignCreateV10BudgetMode      QianchuanCampaignCreateV10BudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanCampaignCreateV10BudgetMode QianchuanCampaignCreateV10BudgetMode = "BUDGET_MODE_INFINITE"
)

// All allowed values of QianchuanCampaignCreateV10BudgetMode enum
var AllowedQianchuanCampaignCreateV10BudgetModeEnumValues = []QianchuanCampaignCreateV10BudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
}

// NewQianchuanCampaignCreateV10BudgetModeFromValue returns a pointer to a valid QianchuanCampaignCreateV10BudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignCreateV10BudgetModeFromValue(v string) (*QianchuanCampaignCreateV10BudgetMode, error) {
	ev := QianchuanCampaignCreateV10BudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignCreateV10BudgetMode: valid values are %v", v, AllowedQianchuanCampaignCreateV10BudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignCreateV10BudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignCreateV10BudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_create_v1.0_budget_mode value
func (v QianchuanCampaignCreateV10BudgetMode) Ptr() *QianchuanCampaignCreateV10BudgetMode {
	return &v
}
