/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserAttachmentUploadV30Request struct for AdvertiserAttachmentUploadV30Request
type AdvertiserAttachmentUploadV30Request struct {
	//
	AdvertiserId   int64                                       `json:"advertiser_id"`
	AttachmentType AdvertiserAttachmentUploadV30AttachmentType `json:"attachment_type"`
	// 文件名 注意：不要包含文件路径，不要含有'/'等非法字符
	Filename string `json:"filename"`
	// 图片数据
	ImageData *FormFileInfo `json:"image_data"`
}
