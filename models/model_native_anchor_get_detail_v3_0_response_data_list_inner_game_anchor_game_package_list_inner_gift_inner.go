/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInnerGiftInner struct for NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInnerGiftInner
type NativeAnchorGetDetailV30ResponseDataListInnerGameAnchorGamePackageListInnerGiftInner struct {
	//
	GiftAmount *int64 `json:"gift_amount,omitempty"`
	//
	GiftImageUrl *string `json:"gift_image_url,omitempty"`
	//
	GiftName *string                                                                `json:"gift_name,omitempty"`
	GiftUnit *NativeAnchorGetDetailV30DataListGameAnchorGamePackageListGiftGiftUnit `json:"gift_unit,omitempty"`
}
