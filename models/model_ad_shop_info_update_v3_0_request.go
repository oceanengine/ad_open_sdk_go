/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdShopInfoUpdateV30Request struct for AdShopInfoUpdateV30Request
type AdShopInfoUpdateV30Request struct {
	//
	AdId int64 `json:"ad_id"`
	//
	AdvertiserId int64                              `json:"advertiser_id"`
	ShopInfo     AdShopInfoUpdateV30RequestShopInfo `json:"shop_info"`
}
