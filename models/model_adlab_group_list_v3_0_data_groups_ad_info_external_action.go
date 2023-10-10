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

// AdlabGroupListV30DataGroupsAdInfoExternalAction
type AdlabGroupListV30DataGroupsAdInfoExternalAction string

// List of adlab_group_list_v3.0_data_groups_ad_info_external_action
const (
	AUTO_MIN_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoExternalAction        AdlabGroupListV30DataGroupsAdInfoExternalAction = "AUTO_MIN_SECOND_STAGE"
	BID_PER_ACTION_AdlabGroupListV30DataGroupsAdInfoExternalAction               AdlabGroupListV30DataGroupsAdInfoExternalAction = "BID_PER_ACTION"
	DEEP_BID_DEFAULT_AdlabGroupListV30DataGroupsAdInfoExternalAction             AdlabGroupListV30DataGroupsAdInfoExternalAction = "DEEP_BID_DEFAULT"
	DEEP_BID_MIN_AdlabGroupListV30DataGroupsAdInfoExternalAction                 AdlabGroupListV30DataGroupsAdInfoExternalAction = "DEEP_BID_MIN"
	DEEP_BID_PACING_AdlabGroupListV30DataGroupsAdInfoExternalAction              AdlabGroupListV30DataGroupsAdInfoExternalAction = "DEEP_BID_PACING"
	DEEP_BID_TYPE_RETENTION_DAYS_AdlabGroupListV30DataGroupsAdInfoExternalAction AdlabGroupListV30DataGroupsAdInfoExternalAction = "DEEP_BID_TYPE_RETENTION_DAYS"
	MIN_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoExternalAction             AdlabGroupListV30DataGroupsAdInfoExternalAction = "MIN_SECOND_STAGE"
	PACING_SECOND_STAGE_AdlabGroupListV30DataGroupsAdInfoExternalAction          AdlabGroupListV30DataGroupsAdInfoExternalAction = "PACING_SECOND_STAGE"
	ROI_COEFFICIENT_AdlabGroupListV30DataGroupsAdInfoExternalAction              AdlabGroupListV30DataGroupsAdInfoExternalAction = "ROI_COEFFICIENT"
	ROI_PACING_AdlabGroupListV30DataGroupsAdInfoExternalAction                   AdlabGroupListV30DataGroupsAdInfoExternalAction = "ROI_PACING"
	SMARTBID_AdlabGroupListV30DataGroupsAdInfoExternalAction                     AdlabGroupListV30DataGroupsAdInfoExternalAction = "SMARTBID"
	SOCIAL_ROI_AdlabGroupListV30DataGroupsAdInfoExternalAction                   AdlabGroupListV30DataGroupsAdInfoExternalAction = "SOCIAL_ROI"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoExternalAction enum
var AllowedAdlabGroupListV30DataGroupsAdInfoExternalActionEnumValues = []AdlabGroupListV30DataGroupsAdInfoExternalAction{
	"AUTO_MIN_SECOND_STAGE",
	"BID_PER_ACTION",
	"DEEP_BID_DEFAULT",
	"DEEP_BID_MIN",
	"DEEP_BID_PACING",
	"DEEP_BID_TYPE_RETENTION_DAYS",
	"MIN_SECOND_STAGE",
	"PACING_SECOND_STAGE",
	"ROI_COEFFICIENT",
	"ROI_PACING",
	"SMARTBID",
	"SOCIAL_ROI",
}

// NewAdlabGroupListV30DataGroupsAdInfoExternalActionFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoExternalAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoExternalActionFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoExternalAction, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoExternalAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoExternalAction: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoExternalActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoExternalAction) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoExternalActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_external_action value
func (v AdlabGroupListV30DataGroupsAdInfoExternalAction) Ptr() *AdlabGroupListV30DataGroupsAdInfoExternalAction {
	return &v
}
