/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtAudienceUpdateV30Request struct for LocalCxtAudienceUpdateV30Request
type LocalCxtAudienceUpdateV30Request struct {
	// 年龄
	Age []*LocalCxtAudienceUpdateV30Age `json:"age,omitempty"`
	// 当district=CITY/COUNTRY时，必填
	City     []int64                            `json:"city,omitempty"`
	District *LocalCxtAudienceUpdateV30District `json:"district,omitempty"`
	Gender   *LocalCxtAudienceUpdateV30Gender   `json:"gender,omitempty"`
	//
	LocalAccountId  int64                                     `json:"local_account_id"`
	PoiAroundRadius *LocalCxtAudienceUpdateV30PoiAroundRadius `json:"poi_around_radius,omitempty"`
}