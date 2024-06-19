/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsUpdateV10Request struct for QianchuanAdKeywordsUpdateV10Request
type QianchuanAdKeywordsUpdateV10Request struct {
	// 计划ID
	AdId int64 `json:"ad_id"`
	// 千川广告账户id
	AdvertiserId int64 `json:"advertiser_id"`
	//
	Keywords []*QianchuanAdKeywordsUpdateV10RequestKeywordsInner `json:"keywords"`
}
