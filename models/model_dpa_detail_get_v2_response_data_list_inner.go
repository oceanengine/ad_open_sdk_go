/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaDetailGetV2ResponseDataListInner struct for DpaDetailGetV2ResponseDataListInner
type DpaDetailGetV2ResponseDataListInner struct {
	//
	Age []int64 `json:"age,omitempty"`
	//
	Bought    *int64                                        `json:"bought,omitempty"`
	BrandInfo *DpaDetailGetV2ResponseDataListInnerBrandInfo `json:"brand_info,omitempty"`
	//
	City []string `json:"city,omitempty"`
	//
	Comments *int64 `json:"comments,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	Feature *string `json:"feature,omitempty"`
	//
	FirstCategory *string `json:"first_category,omitempty"`
	//
	FirstCategoryId *string `json:"first_category_id,omitempty"`
	//
	HasVideo *int32 `json:"has_video,omitempty"`
	//
	ImageUrl *string `json:"image_url,omitempty"`
	//
	ImageUrls []*DpaDetailGetV2ResponseDataListInnerImageUrlsInner `json:"image_urls,omitempty"`
	//
	Label       []string                                        `json:"label,omitempty"`
	LandingInfo *DpaDetailGetV2ResponseDataListInnerLandingInfo `json:"landing_info,omitempty"`
	//
	Mark *float64 `json:"mark,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	OfflineTime *int64 `json:"offline_time,omitempty"`
	//
	OnlineTime *int64 `json:"online_time,omitempty"`
	//
	OuterId *string `json:"outer_id,omitempty"`
	//
	PlatformId *int64                                        `json:"platform_id,omitempty"`
	PriceInfo  *DpaDetailGetV2ResponseDataListInnerPriceInfo `json:"price_info,omitempty"`
	//
	ProductId *int64 `json:"product_id,omitempty"`
	//
	Profession *map[string]string `json:"profession,omitempty"`
	//
	Province       []string                                           `json:"province,omitempty"`
	ShopKeeperInfo *DpaDetailGetV2ResponseDataListInnerShopKeeperInfo `json:"shop_keeper_info,omitempty"`
	//
	SpuId *string `json:"spu_id,omitempty"`
	//
	Status *int32 `json:"status,omitempty"`
	//
	Stock *int32 `json:"stock,omitempty"`
	//
	SubCategory *string `json:"sub_category,omitempty"`
	//
	SubCategoryId *string `json:"sub_category_id,omitempty"`
	//
	Tags []string `json:"tags,omitempty"`
	//
	ThirdCategory *string `json:"third_category,omitempty"`
	//
	ThirdCategoryId *string `json:"third_category_id,omitempty"`
	//
	Title *string `json:"title,omitempty"`
	//
	Video *string `json:"video,omitempty"`
	//
	Videos []*DpaDetailGetV2ResponseDataListInnerVideosInner `json:"videos,omitempty"`
}
