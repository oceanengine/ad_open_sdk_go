/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeDeleteV2Request struct for CreativeCustomCreativeDeleteV2Request
type CreativeCustomCreativeDeleteV2Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	CreativeIds []int64 `json:"creative_ids,omitempty"`
}
