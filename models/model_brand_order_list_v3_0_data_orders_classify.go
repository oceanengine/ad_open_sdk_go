/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandOrderListV30DataOrdersClassify
type BrandOrderListV30DataOrdersClassify string

// List of brand_order_list_v3.0_data_orders_classify
const (
	EXCHANGE_BrandOrderListV30DataOrdersClassify   BrandOrderListV30DataOrdersClassify = "EXCHANGE"
	INTERNAL_BrandOrderListV30DataOrdersClassify   BrandOrderListV30DataOrdersClassify = "INTERNAL"
	SALE_BrandOrderListV30DataOrdersClassify       BrandOrderListV30DataOrdersClassify = "SALE"
	SUBSIDIARY_BrandOrderListV30DataOrdersClassify BrandOrderListV30DataOrdersClassify = "SUBSIDIARY"
	SUPPLEMENT_BrandOrderListV30DataOrdersClassify BrandOrderListV30DataOrdersClassify = "SUPPLEMENT"
)

// Ptr returns reference to brand_order_list_v3.0_data_orders_classify value
func (v BrandOrderListV30DataOrdersClassify) Ptr() *BrandOrderListV30DataOrdersClassify {
	return &v
}
