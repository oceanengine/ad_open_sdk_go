/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeEstimateProfitV10Audience
type QianchuanAwemeEstimateProfitV10Audience struct {
	//
	Age          []*QianchuanAwemeEstimateProfitV10AudienceAge       `json:"age,omitempty"`
	AudienceMode QianchuanAwemeEstimateProfitV10AudienceAudienceMode `json:"audience_mode"`
	//
	AuthorIds []int64 `json:"author_ids,omitempty"`
	//
	Behaviors []*QianchuanAwemeEstimateProfitV10AudienceBehaviors `json:"behaviors,omitempty"`
	//
	City                 []int64                                                      `json:"city,omitempty"`
	District             *QianchuanAwemeEstimateProfitV10AudienceDistrict             `json:"district,omitempty"`
	ExcludeLimitedRegion *QianchuanAwemeEstimateProfitV10AudienceExcludeLimitedRegion `json:"exclude_limited_region,omitempty"`
	Gender               *QianchuanAwemeEstimateProfitV10AudienceGender               `json:"gender,omitempty"`
	//
	InterestCategories []int64 `json:"interest_categories,omitempty"`
}
