/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserAvatarSubmitV2Request struct for AdvertiserAvatarSubmitV2Request
type AdvertiserAvatarSubmitV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ImageId string `json:"image_id"`
	//
	SourceInfo *string `json:"source_info,omitempty"`
}
