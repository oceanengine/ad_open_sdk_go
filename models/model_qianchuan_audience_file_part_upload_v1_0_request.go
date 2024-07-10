/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceFilePartUploadV10Request struct for QianchuanAudienceFilePartUploadV10Request
type QianchuanAudienceFilePartUploadV10Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	File *FormFileInfo `json:"file"`
	//
	FileKey *string `json:"file_key,omitempty"`
	//
	IsFinished int32 `json:"is_finished"`
	//
	PartNum int64 `json:"part_num"`
}
