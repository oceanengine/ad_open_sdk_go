/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CdpBrandGetV30ResponseData
type CdpBrandGetV30ResponseData struct {
	// 品牌所属类别
	BrandCategoryInfo []*CdpBrandGetV30ResponseDataBrandCategoryInfoInner `json:"brand_category_info,omitempty"`
	//
	BrandInfoData []*CdpBrandGetV30ResponseDataBrandInfoDataInner `json:"brand_info_data,omitempty"`
}
