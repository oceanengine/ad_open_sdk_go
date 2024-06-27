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

// BrandCampaignGetV30DataCampaignsMarketingPurpose
type BrandCampaignGetV30DataCampaignsMarketingPurpose int64

// List of brand_campaign_get_v3.0_data_campaigns_marketing_purpose
const (
	Enum_0_BrandCampaignGetV30DataCampaignsMarketingPurpose BrandCampaignGetV30DataCampaignsMarketingPurpose = 0
	Enum_1_BrandCampaignGetV30DataCampaignsMarketingPurpose BrandCampaignGetV30DataCampaignsMarketingPurpose = 1
	Enum_2_BrandCampaignGetV30DataCampaignsMarketingPurpose BrandCampaignGetV30DataCampaignsMarketingPurpose = 2
	Enum_3_BrandCampaignGetV30DataCampaignsMarketingPurpose BrandCampaignGetV30DataCampaignsMarketingPurpose = 3
)

// All allowed values of BrandCampaignGetV30DataCampaignsMarketingPurpose enum
var AllowedBrandCampaignGetV30DataCampaignsMarketingPurposeEnumValues = []BrandCampaignGetV30DataCampaignsMarketingPurpose{
	0,
	1,
	2,
	3,
}

// NewBrandCampaignGetV30DataCampaignsMarketingPurposeFromValue returns a pointer to a valid BrandCampaignGetV30DataCampaignsMarketingPurpose
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandCampaignGetV30DataCampaignsMarketingPurposeFromValue(v int64) (*BrandCampaignGetV30DataCampaignsMarketingPurpose, error) {
	ev := BrandCampaignGetV30DataCampaignsMarketingPurpose(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandCampaignGetV30DataCampaignsMarketingPurpose: valid values are %v", v, AllowedBrandCampaignGetV30DataCampaignsMarketingPurposeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandCampaignGetV30DataCampaignsMarketingPurpose) IsValid() bool {
	for _, existing := range AllowedBrandCampaignGetV30DataCampaignsMarketingPurposeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_campaign_get_v3.0_data_campaigns_marketing_purpose value
func (v BrandCampaignGetV30DataCampaignsMarketingPurpose) Ptr() *BrandCampaignGetV30DataCampaignsMarketingPurpose {
	return &v
}
