/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkg 推广产品资质信息
type AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkg struct {
	//  来自【推广产品资质规则配置查询接口】，行业的资质规则中的config_id
	ConfigId int64 `json:"config_id"`
	// 一级到三级行业id
	IndustryId []int64 `json:"industry_id"`
	// 一级到三级行业名称
	IndustryName     []string                                                            `json:"industry_name"`
	NecessaryCombine *AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgNecessaryCombine `json:"necessary_combine,omitempty"`
	Permission       AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgPermission        `json:"permission"`
	// 推广产品组id
	PkgId *int64 `json:"pkg_id,omitempty"`
	// 用户提交的推广产品名称
	ProductName string                                           `json:"product_name"`
	Status      AdvertiserDeliveryPkgGetV30DataDeliveryPkgStatus `json:"status"`
	// 选填资质模块
	UnnecessaryCombines []*AdvertiserDeliveryPkgGetV30ResponseDataDeliveryPkgUnnecessaryCombinesInner `json:"unnecessary_combines,omitempty"`
}
