/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdUdUpdateV2Request struct for AdUdUpdateV2Request
type AdUdUpdateV2Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64                     `json:"advertiser_id"`
	UdShop       AdUdUpdateV2RequestUdShop `json:"ud_shop"`
}
