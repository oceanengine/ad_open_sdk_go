/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV30Request struct for KeywordUpdateV30Request
type KeywordUpdateV30Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Keywords []*KeywordUpdateV30RequestKeywordsInner `json:"keywords"`
	//
	PromotionId int64 `json:"promotion_id"`
}