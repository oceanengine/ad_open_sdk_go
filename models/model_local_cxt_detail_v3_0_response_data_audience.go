/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtDetailV30ResponseDataAudience
type LocalCxtDetailV30ResponseDataAudience struct {
	//
	Age []*LocalCxtDetailV30DataAudienceAge `json:"age,omitempty"`
	//
	City            []int64                                       `json:"city,omitempty"`
	District        *LocalCxtDetailV30DataAudienceDistrict        `json:"district,omitempty"`
	Gender          *LocalCxtDetailV30DataAudienceGender          `json:"gender,omitempty"`
	PoiAroundRadius *LocalCxtDetailV30DataAudiencePoiAroundRadius `json:"poi_around_radius,omitempty"`
}
