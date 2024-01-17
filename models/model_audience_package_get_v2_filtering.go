/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2Filtering
type AudiencePackageGetV2Filtering struct {
	AdType        *AudiencePackageGetV2FilteringAdType        `json:"ad_type,omitempty"`
	DeliveryRange *AudiencePackageGetV2FilteringDeliveryRange `json:"delivery_range,omitempty"`
	LandingType   *AudiencePackageGetV2FilteringLandingType   `json:"landing_type,omitempty"`
	MarketingGoal *AudiencePackageGetV2FilteringMarketingGoal `json:"marketing_goal,omitempty"`
}
