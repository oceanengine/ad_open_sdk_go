/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeUpdateV2Request struct for CreativeProceduralCreativeUpdateV2Request
type CreativeProceduralCreativeUpdateV2Request struct {
	AdData CreativeProceduralCreativeUpdateV2RequestAdData `json:"ad_data"`
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64                                             `json:"advertiser_id"`
	Creative     CreativeProceduralCreativeUpdateV2RequestCreative `json:"creative"`
}
