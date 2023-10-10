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

// QianchuanCampaignListGetV10DataListBudgetMode
type QianchuanCampaignListGetV10DataListBudgetMode string

// List of qianchuan_campaign_list_get_v1.0_data_list_budget_mode
const (
	BUDGET_MODE_DAY_QianchuanCampaignListGetV10DataListBudgetMode      QianchuanCampaignListGetV10DataListBudgetMode = "BUDGET_MODE_DAY"
	BUDGET_MODE_INFINITE_QianchuanCampaignListGetV10DataListBudgetMode QianchuanCampaignListGetV10DataListBudgetMode = "BUDGET_MODE_INFINITE"
	BUDGET_MODE_TOTAL_QianchuanCampaignListGetV10DataListBudgetMode    QianchuanCampaignListGetV10DataListBudgetMode = "BUDGET_MODE_TOTAL"
)

// All allowed values of QianchuanCampaignListGetV10DataListBudgetMode enum
var AllowedQianchuanCampaignListGetV10DataListBudgetModeEnumValues = []QianchuanCampaignListGetV10DataListBudgetMode{
	"BUDGET_MODE_DAY",
	"BUDGET_MODE_INFINITE",
	"BUDGET_MODE_TOTAL",
}

// NewQianchuanCampaignListGetV10DataListBudgetModeFromValue returns a pointer to a valid QianchuanCampaignListGetV10DataListBudgetMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10DataListBudgetModeFromValue(v string) (*QianchuanCampaignListGetV10DataListBudgetMode, error) {
	ev := QianchuanCampaignListGetV10DataListBudgetMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10DataListBudgetMode: valid values are %v", v, AllowedQianchuanCampaignListGetV10DataListBudgetModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10DataListBudgetMode) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10DataListBudgetModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_data_list_budget_mode value
func (v QianchuanCampaignListGetV10DataListBudgetMode) Ptr() *QianchuanCampaignListGetV10DataListBudgetMode {
	return &v
}
