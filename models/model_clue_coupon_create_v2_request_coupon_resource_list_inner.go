/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponCreateV2RequestCouponResourceListInner struct for ClueCouponCreateV2RequestCouponResourceListInner
type ClueCouponCreateV2RequestCouponResourceListInner struct {
	//
	Amount *int64 `json:"amount,omitempty"`
	//
	AndroidLink *string `json:"android_link,omitempty"`
	//
	Code     *string                                      `json:"code,omitempty"`
	CodeType ClueCouponCreateV2CouponResourceListCodeType `json:"code_type"`
	//
	Condition string `json:"condition"`
	//
	Discount *int64 `json:"discount,omitempty"`
	//
	GiftList []*ClueCouponCreateV2RequestCouponResourceListInnerGiftListInner `json:"gift_list,omitempty"`
	//
	HeadImageUrl string                                           `json:"head_image_url"`
	IndustryType ClueCouponCreateV2CouponResourceListIndustryType `json:"industry_type"`
	//
	IosLink *string `json:"ios_link,omitempty"`
	//
	Link *string `json:"link,omitempty"`
	//
	LogoImageUrl string `json:"logo_image_url"`
	//
	MaxAmount *int64 `json:"max_amount,omitempty"`
	//
	MerchantName string `json:"merchant_name"`
	//
	MinAmount *int64 `json:"min_amount,omitempty"`
	//
	Notification string `json:"notification"`
	//
	ReliefAmount *int64                                           `json:"relief_amount,omitempty"`
	ResourceType ClueCouponCreateV2CouponResourceListResourceType `json:"resource_type"`
	//
	ServiceNum string `json:"service_num"`
	//
	Stock *int64 `json:"stock,omitempty"`
	//
	StoreIds []int64 `json:"store_ids,omitempty"`
	//
	Title   string                                      `json:"title"`
	UseType ClueCouponCreateV2CouponResourceListUseType `json:"use_type"`
	//
	ValidDays *int64 `json:"valid_days,omitempty"`
	//
	ValidEnd **string `json:"valid_end,omitempty"`
	//
	ValidStart **string `json:"valid_start,omitempty"`
}
