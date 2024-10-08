/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AsyncTaskCreateV2RequestTaskParamsFiltering
type AsyncTaskCreateV2RequestTaskParamsFiltering struct {
	//
	AdId []int64 `json:"ad_id,omitempty"`
	//
	Bidword []string `json:"bidword,omitempty"`
	//
	CampaignId []int64 `json:"campaign_id,omitempty"`
	//
	CampaignType []string `json:"campaign_type,omitempty"`
	//
	ConvertType []string `json:"convert_type,omitempty"`
	//
	CreativeId []int64 `json:"creative_id,omitempty"`
	//
	CreativeMaterialMode []string `json:"creative_material_mode,omitempty"`
	//
	ImageMode []string `json:"image_mode,omitempty"`
	//
	InventoryType []string `json:"inventory_type,omitempty"`
	//
	LandingType []string `json:"landing_type,omitempty"`
	//
	Platform []string `json:"platform,omitempty"`
	//
	Pricing []string `json:"pricing,omitempty"`
	//
	PricingCategory []string `json:"pricing_category,omitempty"`
	//
	Query []string `json:"query,omitempty"`
}
