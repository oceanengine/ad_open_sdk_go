/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CampaignGetV2FilteringLandingType
type CampaignGetV2FilteringLandingType string

// List of campaign_get_v2_filtering_landing_type
const (
	APP_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "APP"
	LINK_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LINK"
	DPA_CampaignGetV2FilteringLandingType            CampaignGetV2FilteringLandingType = "DPA"
	AWEME_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "AWEME"
	QUICK_APP_CampaignGetV2FilteringLandingType      CampaignGetV2FilteringLandingType = "QUICK_APP"
	SHOP_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "SHOP"
	ARTICLE_CampaignGetV2FilteringLandingType        CampaignGetV2FilteringLandingType = "ARTICLE"
	STORE_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "STORE"
	BRAND_EXTERNAL_CampaignGetV2FilteringLandingType CampaignGetV2FilteringLandingType = "BRAND_EXTERNAL"
	LIVE_CampaignGetV2FilteringLandingType           CampaignGetV2FilteringLandingType = "LIVE"
	GOODS_CampaignGetV2FilteringLandingType          CampaignGetV2FilteringLandingType = "GOODS"
)

// All allowed values of CampaignGetV2FilteringLandingType enum
var AllowedCampaignGetV2FilteringLandingTypeEnumValues = []CampaignGetV2FilteringLandingType{
	"APP",
	"LINK",
	"DPA",
	"AWEME",
	"QUICK_APP",
	"SHOP",
	"ARTICLE",
	"STORE",
	"BRAND_EXTERNAL",
	"LIVE",
	"GOODS",
}

// NewCampaignGetV2FilteringLandingTypeFromValue returns a pointer to a valid CampaignGetV2FilteringLandingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignGetV2FilteringLandingTypeFromValue(v string) (*CampaignGetV2FilteringLandingType, error) {
	ev := CampaignGetV2FilteringLandingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignGetV2FilteringLandingType: valid values are %v", v, AllowedCampaignGetV2FilteringLandingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignGetV2FilteringLandingType) IsValid() bool {
	for _, existing := range AllowedCampaignGetV2FilteringLandingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign_get_v2_filtering_landing_type value
func (v CampaignGetV2FilteringLandingType) Ptr() *CampaignGetV2FilteringLandingType {
	return &v
}
