/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileQualitySubmitV30Request struct for FileQualitySubmitV30Request
type FileQualitySubmitV30Request struct {
	//
	AdvertiserId    int64                               `json:"advertiser_id"`
	MaterialChannel FileQualitySubmitV30MaterialChannel `json:"material_channel"`
	MaterialType    FileQualitySubmitV30MaterialType    `json:"material_type"`
	//
	VideoUrl string `json:"video_url"`
}
