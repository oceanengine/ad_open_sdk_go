/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DmpDataSourceFileUploadV2Request struct for DmpDataSourceFileUploadV2Request
type DmpDataSourceFileUploadV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	File *FormFileInfo `json:"file"`
	//
	FileSignature *string `json:"file_signature,omitempty"`
}