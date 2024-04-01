/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAdPreviewQrcodeGetV30ResponseDataData
type ToolsAdPreviewQrcodeGetV30ResponseDataData struct {
	//
	AdvertiserId int64 `json:"advertiser_id"`
	//
	IdType string `json:"id_type"`
	//
	MaterialId *int64 `json:"material_id,omitempty"`
	//
	PromotionId int64 `json:"promotion_id"`
	//
	QrcodeMsgUrl string `json:"qrcode_msg_url"`
}
