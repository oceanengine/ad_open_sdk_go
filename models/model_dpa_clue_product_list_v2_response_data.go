/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaClueProductListV2ResponseData
type DpaClueProductListV2ResponseData struct {
	PageInfo *DpaClueProductListV2ResponseDataPageInfo `json:"page_info,omitempty"`
	//
	Products []*DpaClueProductListV2ResponseDataProductsInner `json:"products,omitempty"`
}
