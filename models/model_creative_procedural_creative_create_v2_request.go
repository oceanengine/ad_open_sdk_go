/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2Request struct for CreativeProceduralCreativeCreateV2Request
type CreativeProceduralCreativeCreateV2Request struct {
	AdData CreativeProceduralCreativeCreateV2RequestAdData `json:"ad_data"`
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64                                             `json:"advertiser_id"`
	Creative     CreativeProceduralCreativeCreateV2RequestCreative `json:"creative"`
}
