/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CluePackageUploadV2Request struct for CluePackageUploadV2Request
type CluePackageUploadV2Request struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AppId string `json:"app_id"`
	//
	AppSecret *string `json:"app_secret,omitempty"`
	//
	CodeResource interface{} `json:"code_resource"`
	//
	InstanceIds []int64 `json:"instance_ids,omitempty"`
	//
	ResourceId *string `json:"resource_id,omitempty"`
	//
	SiteId *string `json:"site_id,omitempty"`
	//
	TemplateId *int64 `json:"template_id,omitempty"`
}
