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

// CampaignUpdateStatusV2OptStatus
type CampaignUpdateStatusV2OptStatus string

// List of campaign_update_status_v2_opt_status
const (
	ENABLE_CampaignUpdateStatusV2OptStatus  CampaignUpdateStatusV2OptStatus = "enable"
	DISABLE_CampaignUpdateStatusV2OptStatus CampaignUpdateStatusV2OptStatus = "disable"
	DELETE_CampaignUpdateStatusV2OptStatus  CampaignUpdateStatusV2OptStatus = "delete"
)

// All allowed values of CampaignUpdateStatusV2OptStatus enum
var AllowedCampaignUpdateStatusV2OptStatusEnumValues = []CampaignUpdateStatusV2OptStatus{
	"enable",
	"disable",
	"delete",
}

// NewCampaignUpdateStatusV2OptStatusFromValue returns a pointer to a valid CampaignUpdateStatusV2OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignUpdateStatusV2OptStatusFromValue(v string) (*CampaignUpdateStatusV2OptStatus, error) {
	ev := CampaignUpdateStatusV2OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignUpdateStatusV2OptStatus: valid values are %v", v, AllowedCampaignUpdateStatusV2OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignUpdateStatusV2OptStatus) IsValid() bool {
	for _, existing := range AllowedCampaignUpdateStatusV2OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_update_status_v2_opt_status value
func (v CampaignUpdateStatusV2OptStatus) Ptr() *CampaignUpdateStatusV2OptStatus {
	return &v
}
