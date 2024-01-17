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

// CampaignGetV2DataListDedicateType
type CampaignGetV2DataListDedicateType string

// List of campaign_get_v2_data_list_dedicate_type
const (
	UNSET_CampaignGetV2DataListDedicateType     CampaignGetV2DataListDedicateType = "UNSET"
	DEDICATED_CampaignGetV2DataListDedicateType CampaignGetV2DataListDedicateType = "DEDICATED"
)

// All allowed values of CampaignGetV2DataListDedicateType enum
var AllowedCampaignGetV2DataListDedicateTypeEnumValues = []CampaignGetV2DataListDedicateType{
	"UNSET",
	"DEDICATED",
}

// NewCampaignGetV2DataListDedicateTypeFromValue returns a pointer to a valid CampaignGetV2DataListDedicateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2DataListDedicateTypeFromValue(v string) (*CampaignGetV2DataListDedicateType, error) {
	ev := CampaignGetV2DataListDedicateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2DataListDedicateType: valid values are %v", v, AllowedCampaignGetV2DataListDedicateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2DataListDedicateType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2DataListDedicateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_data_list_dedicate_type value
func (v CampaignGetV2DataListDedicateType) Ptr() *CampaignGetV2DataListDedicateType {
	return &v
}
