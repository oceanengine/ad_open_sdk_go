/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserAvatarUploadV2Request struct for AdvertiserAvatarUploadV2Request
type AdvertiserAvatarUploadV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ImageFile *FormFileInfo `json:"image_file"`
}
