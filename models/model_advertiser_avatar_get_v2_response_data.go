/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserAvatarGetV2ResponseData
type AdvertiserAvatarGetV2ResponseData struct {
	//
	AdvertiserId *int64                                       `json:"advertiser_id,omitempty"`
	AvatarInfo   *AdvertiserAvatarGetV2ResponseDataAvatarInfo `json:"avatar_info,omitempty"`
	//
	AvatarReason *string `json:"avatar_reason,omitempty"`
	//
	AvatarStatus *string `json:"avatar_status,omitempty"`
	//
	SourceInfo *string `json:"source_info,omitempty"`
	//
	SourceReason *string `json:"source_reason,omitempty"`
	//
	SourceStatus *string `json:"source_status,omitempty"`
}
