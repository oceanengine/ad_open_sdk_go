/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanUniPromotionProductGetV10ResponseDataProductListInner struct for QianchuanUniPromotionProductGetV10ResponseDataProductListInner
type QianchuanUniPromotionProductGetV10ResponseDataProductListInner struct {
	//
	AuditTime *string `json:"audit_time,omitempty"`
	//
	CategoryName *string `json:"category_name,omitempty"`
	//
	ChannelId   *int64                                                        `json:"channel_id,omitempty"`
	ChannelType *QianchuanUniPromotionProductGetV10DataProductListChannelType `json:"channel_type,omitempty"`
	//
	GrayReason []string `json:"gray_reason,omitempty"`
	//
	Id *int64 `json:"id,omitempty"`
	//
	Img *string `json:"img,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	SellNum *int64 `json:"sell_num,omitempty"`
	//
	SquareImageList []*QianchuanUniPromotionProductGetV10ResponseDataProductListInnerSquareImageListInner `json:"square_image_list,omitempty"`
	//
	StockNum *int64 `json:"stock_num,omitempty"`
	//
	Tag []string `json:"tag,omitempty"`
}
