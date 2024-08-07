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

// BrandCampaignGetV30CampaignStatus
type BrandCampaignGetV30CampaignStatus int64

// List of brand_campaign_get_v3.0_campaign_status
const (
	Enum_1_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 1
	Enum_10_BrandCampaignGetV30CampaignStatus BrandCampaignGetV30CampaignStatus = 10
	Enum_11_BrandCampaignGetV30CampaignStatus BrandCampaignGetV30CampaignStatus = 11
	Enum_2_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 2
	Enum_3_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 3
	Enum_4_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 4
	Enum_5_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 5
	Enum_6_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 6
	Enum_7_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 7
	Enum_8_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 8
	Enum_9_BrandCampaignGetV30CampaignStatus  BrandCampaignGetV30CampaignStatus = 9
)

// All allowed values of BrandCampaignGetV30CampaignStatus enum
var AllowedBrandCampaignGetV30CampaignStatusEnumValues = []BrandCampaignGetV30CampaignStatus{
	1,
	10,
	11,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
}

// NewBrandCampaignGetV30CampaignStatusFromValue returns a pointer to a valid BrandCampaignGetV30CampaignStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCampaignGetV30CampaignStatusFromValue(v int64) (*BrandCampaignGetV30CampaignStatus, error) {
	ev := BrandCampaignGetV30CampaignStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCampaignGetV30CampaignStatus: valid values are %v", v, AllowedBrandCampaignGetV30CampaignStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCampaignGetV30CampaignStatus) IsValid() bool {
	for _, existing := range AllowedBrandCampaignGetV30CampaignStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_campaign_get_v3.0_campaign_status value
func (v BrandCampaignGetV30CampaignStatus) Ptr() *BrandCampaignGetV30CampaignStatus {
	return &v
}
