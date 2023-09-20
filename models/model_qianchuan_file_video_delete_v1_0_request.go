/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanFileVideoDeleteV10Request struct for QianchuanFileVideoDeleteV10Request
type QianchuanFileVideoDeleteV10Request struct {
	// 广告主ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 待删除的video_id列表，长度范围：1 ~ 100
	VideoIds []string `json:"video_ids"`
}
