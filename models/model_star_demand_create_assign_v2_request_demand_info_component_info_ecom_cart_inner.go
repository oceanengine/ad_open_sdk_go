/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateAssignV2RequestDemandInfoComponentInfoEcomCartInner struct for StarDemandCreateAssignV2RequestDemandInfoComponentInfoEcomCartInner
type StarDemandCreateAssignV2RequestDemandInfoComponentInfoEcomCartInner struct {
	// 购物车商品链接
	ProductLink *string `json:"product_link,omitempty"`
	// 购物车商品标题，不超过50字
	Title *string `json:"title,omitempty"`
}
