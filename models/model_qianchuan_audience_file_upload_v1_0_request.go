/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceFileUploadV10Request struct for QianchuanAudienceFileUploadV10Request
type QianchuanAudienceFileUploadV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	File *FormFileInfo `json:"file"`
}
