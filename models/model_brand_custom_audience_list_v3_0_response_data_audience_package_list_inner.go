/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCustomAudienceListV30ResponseDataAudiencePackageListInner struct for BrandCustomAudienceListV30ResponseDataAudiencePackageListInner
type BrandCustomAudienceListV30ResponseDataAudiencePackageListInner struct {
	Category *BrandCustomAudienceListV30DataAudiencePackageListCategory `json:"category,omitempty"`
	// 过期时间
	ExpireTime int64 `json:"expire_time"`
	// 人群包ID
	Id int64 `json:"id"`
	// 人群包名称
	Name string `json:"name"`
}
