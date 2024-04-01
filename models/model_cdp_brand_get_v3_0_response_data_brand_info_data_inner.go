/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CdpBrandGetV30ResponseDataBrandInfoDataInner struct for CdpBrandGetV30ResponseDataBrandInfoDataInner
type CdpBrandGetV30ResponseDataBrandInfoDataInner struct {
	//
	BrandNameId *int64 `json:"brand_name_id,omitempty"`
	//
	CdpBrandId *int64 `json:"cdp_brand_id,omitempty"`
	//
	CdpBrandName *string `json:"cdp_brand_name,omitempty"`
	//
	EcomBrandId *int64 `json:"ecom_brand_id,omitempty"`
	//
	SubBrandMap map[string]interface{} `json:"sub_brand_map,omitempty"`
}
