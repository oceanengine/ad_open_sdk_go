/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionShopInfoUpdateV30Request struct for PromotionShopInfoUpdateV30Request
type PromotionShopInfoUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	PromotionId int64                                     `json:"promotion_id"`
	ShopInfo    PromotionShopInfoUpdateV30RequestShopInfo `json:"shop_info"`
}
