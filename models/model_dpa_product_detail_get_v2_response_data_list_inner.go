/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaProductDetailGetV2ResponseDataListInner struct for DpaProductDetailGetV2ResponseDataListInner
type DpaProductDetailGetV2ResponseDataListInner struct {
	//
	BrandName *string `json:"brand_name,omitempty"`
	//
	City []string `json:"city,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	ExternalUrl *string `json:"external_url,omitempty"`
	//
	FirstCategory *string `json:"first_category,omitempty"`
	//
	HasVideo *bool `json:"has_video,omitempty"`
	//
	ImageUrl *string `json:"image_url,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OfflineTime **string `json:"offline_time,omitempty"`
	//
	OnlineTime **string `json:"online_time,omitempty"`
	//
	PlatformId *int64 `json:"platform_id,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	SubCategory *string `json:"sub_category,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	ThirdCategory *string `json:"third_category,omitempty"`
	//
	Title *string `json:"title,omitempty"`
}
