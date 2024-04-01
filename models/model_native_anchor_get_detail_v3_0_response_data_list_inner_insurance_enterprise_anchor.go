/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchor 外跳锚点
type NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchor struct {
	// 详情介绍，1-18个字
	BannerDescription *string                                                                            `json:"banner_description,omitempty"`
	BannerImage       *NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchorBannerImage `json:"banner_image,omitempty"`
	ConversionBtn     *NativeAnchorGetDetailV30DataListInsuranceEnterpriseAnchorConversionBtn            `json:"conversion_btn,omitempty"`
	// 点击按钮时的跳转链接，此处填写的跳转链接会应用到“转化按钮”与“详情介绍”，以http开头
	DetailUrl *string `json:"detail_url,omitempty"`
	// 产品描述
	ProductDescriptions []string                                                                            `json:"product_descriptions,omitempty"`
	ProductImage        *NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchorProductImage `json:"product_image,omitempty"`
	// 产品详情下的产品名称
	ProductName *string `json:"product_name,omitempty"`
	// 服务描述，1-6字
	ProductServiceDescription *string `json:"product_service_description,omitempty"`
	// 产品特点，最多返回3个
	ProductTags []string `json:"product_tags,omitempty"`
	// 产品名称，1-15字
	ProductTitle *string `json:"product_title,omitempty"`
	// 单项服务名称，5-10组
	SingleProductInfo []*NativeAnchorGetDetailV30ResponseDataListInnerInsuranceEnterpriseAnchorSingleProductInfoInner `json:"single_product_info,omitempty"`
}
