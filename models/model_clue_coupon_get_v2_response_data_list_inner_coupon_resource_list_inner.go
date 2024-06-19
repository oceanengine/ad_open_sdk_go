/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponGetV2ResponseDataListInnerCouponResourceListInner struct for ClueCouponGetV2ResponseDataListInnerCouponResourceListInner
type ClueCouponGetV2ResponseDataListInnerCouponResourceListInner struct {
	//
	AndroidLink *string `json:"android_link,omitempty"`
	//
	AuditMsg *string `json:"audit_msg,omitempty"`
	//
	Code *string `json:"code,omitempty"`
	//
	CodeReferUrl *string                                            `json:"code_refer_url,omitempty"`
	CodeType     *ClueCouponGetV2DataListCouponResourceListCodeType `json:"code_type,omitempty"`
	//
	Condition *string `json:"condition,omitempty"`
	//
	Discount *int64 `json:"discount,omitempty"`
	//
	GiftList []*ClueCouponGetV2ResponseDataListInnerCouponResourceListInnerGiftListInner `json:"gift_list,omitempty"`
	//
	HeadImageUrl *string                                                `json:"head_image_url,omitempty"`
	IndustryType *ClueCouponGetV2DataListCouponResourceListIndustryType `json:"industry_type,omitempty"`
	//
	IosLink *string `json:"ios_link,omitempty"`
	//
	Link *string `json:"link,omitempty"`
	//
	LogoImageUrl *string `json:"logo_image_url,omitempty"`
	//
	MaxAmount *int64 `json:"max_amount,omitempty"`
	//
	MerchantName *string `json:"merchant_name,omitempty"`
	//
	MinAmount *int64 `json:"min_amount,omitempty"`
	//
	Notification *string `json:"notification,omitempty"`
	//
	ReliefAmount *int64 `json:"relief_amount,omitempty"`
	//
	ResourceId   *int64                                                 `json:"resource_id,omitempty"`
	ResourceType *ClueCouponGetV2DataListCouponResourceListResourceType `json:"resource_type,omitempty"`
	//
	ServiceNum *string `json:"service_num,omitempty"`
	//
	Stock *int64 `json:"stock,omitempty"`
	//
	StoreIds []int64 `json:"store_ids,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	UseType *string `json:"use_type,omitempty"`
	//
	ValidDays *int64 `json:"valid_days,omitempty"`
	//
	ValidEnd *string `json:"valid_end,omitempty"`
	//
	ValidStart *string `json:"valid_start,omitempty"`
}
