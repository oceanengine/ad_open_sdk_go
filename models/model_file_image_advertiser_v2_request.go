/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileImageAdvertiserV2Request struct for FileImageAdvertiserV2Request
type FileImageAdvertiserV2Request struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	ImageFile **FormFileInfo `json:"image_file,omitempty"`
	//
	ImageSignature *string `json:"image_signature,omitempty"`
	//
	ImageUrl   *string                         `json:"image_url,omitempty"`
	UploadTo   FileImageAdvertiserV2UploadTo   `json:"upload_to"`
	UploadType FileImageAdvertiserV2UploadType `json:"upload_type"`
}
