/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeProductAvailableGetV10ResponseDataProductListInner struct for QianchuanAwemeProductAvailableGetV10ResponseDataProductListInner
type QianchuanAwemeProductAvailableGetV10ResponseDataProductListInner struct {
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	ChannelId   *int64                                                          `json:"channel_id,omitempty"`
	ChannelType *QianchuanAwemeProductAvailableGetV10DataProductListChannelType `json:"channel_type,omitempty"`
	//
	DiscountHigherPrice *float64 `json:"discount_higher_price,omitempty"`
	//
	DiscountLowerPrice *float64 `json:"discount_lower_price,omitempty"`
	//
	Id int64 `json:"id"`
	//
	Img string `json:"img"`
	//
	Inventory *int64 `json:"inventory,omitempty"`
	//
	MarketingPrice *float64 `json:"marketing_price,omitempty"`
	//
	Name string `json:"name"`
	//
	ProductRate *int64 `json:"product_rate,omitempty"`
	//
	SupportProductNewOpen *bool `json:"support_product_new_open,omitempty"`
}
