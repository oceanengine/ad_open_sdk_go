/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaVideoGetV2ResponseDataListInner struct for DpaVideoGetV2ResponseDataListInner
type DpaVideoGetV2ResponseDataListInner struct {
	//
	CreateTime *string                                      `json:"create_time,omitempty"`
	ImageInfo  *DpaVideoGetV2ResponseDataListInnerImageInfo `json:"image_info,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	PackageId *string `json:"package_id,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	ProductPlatformId *int64                                       `json:"product_platform_id,omitempty"`
	VideoInfo         *DpaVideoGetV2ResponseDataListInnerVideoInfo `json:"video_info,omitempty"`
}