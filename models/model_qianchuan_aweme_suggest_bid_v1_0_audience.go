/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeSuggestBidV10Audience
type QianchuanAwemeSuggestBidV10Audience struct {
	//
	Age          []*QianchuanAwemeSuggestBidV10AudienceAge       `json:"age,omitempty"`
	AudienceMode QianchuanAwemeSuggestBidV10AudienceAudienceMode `json:"audience_mode"`
	//
	AuthorIds []int64 `json:"author_ids,omitempty"`
	//
	City                 []int64                                                  `json:"city,omitempty"`
	District             *QianchuanAwemeSuggestBidV10AudienceDistrict             `json:"district,omitempty"`
	ExcludeLimitedRegion *QianchuanAwemeSuggestBidV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAwemeSuggestBidV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
}
