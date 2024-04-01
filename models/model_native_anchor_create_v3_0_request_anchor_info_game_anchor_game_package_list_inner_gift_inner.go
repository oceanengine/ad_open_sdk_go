/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInnerGiftInner struct for NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInnerGiftInner
type NativeAnchorCreateV30RequestAnchorInfoGameAnchorGamePackageListInnerGiftInner struct {
	// 礼品数量，0～6
	GiftAmount *int64 `json:"gift_amount,omitempty"`
	// 礼品图片
	GiftImageUrl *string `json:"gift_image_url,omitempty"`
	// 礼品名称，字符限制0～8
	GiftName *string                                                               `json:"gift_name,omitempty"`
	GiftUnit *NativeAnchorCreateV30AnchorInfoGameAnchorGamePackageListGiftGiftUnit `json:"gift_unit,omitempty"`
}
