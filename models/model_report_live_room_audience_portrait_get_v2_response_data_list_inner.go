/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportLiveRoomAudiencePortraitGetV2ResponseDataListInner struct for ReportLiveRoomAudiencePortraitGetV2ResponseDataListInner
type ReportLiveRoomAudiencePortraitGetV2ResponseDataListInner struct {
	//
	Age *string `json:"age,omitempty"`
	//
	CityName *string `json:"city_name,omitempty"`
	//
	Fields map[string]interface{} `json:"fields,omitempty"`
	//
	Gender *string `json:"gender,omitempty"`
	//
	Platform *string `json:"platform,omitempty"`
	//
	ProvinceName *string `json:"province_name,omitempty"`
}
