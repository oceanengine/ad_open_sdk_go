/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgDeleteV30ResponseData
type AdvertiserDeliveryPkgDeleteV30ResponseData struct {
	// 删除失败的行业产品列表
	Errors []*AdvertiserDeliveryPkgDeleteV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 删除成功的推广产品id
	PkgIds []int64 `json:"pkg_ids,omitempty"`
}
