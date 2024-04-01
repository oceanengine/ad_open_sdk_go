/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsAdConvertQueryV2PromotionContent
type ToolsAdConvertQueryV2PromotionContent string

// List of tools_ad_convert_query_v2_promotion_content
const (
	EXTERNAL_URL_ToolsAdConvertQueryV2PromotionContent    ToolsAdConvertQueryV2PromotionContent = "EXTERNAL_URL"
	MICRO_APP_ToolsAdConvertQueryV2PromotionContent       ToolsAdConvertQueryV2PromotionContent = "MICRO_APP"
	DOWNLOAD_URL_ToolsAdConvertQueryV2PromotionContent    ToolsAdConvertQueryV2PromotionContent = "DOWNLOAD_URL"
	QUICK_APP_URL_ToolsAdConvertQueryV2PromotionContent   ToolsAdConvertQueryV2PromotionContent = "QUICK_APP_URL"
	GOODS_LINK_ToolsAdConvertQueryV2PromotionContent      ToolsAdConvertQueryV2PromotionContent = "GOODS_LINK"
	LIVE_ROOM_ToolsAdConvertQueryV2PromotionContent       ToolsAdConvertQueryV2PromotionContent = "LIVE_ROOM"
	DOUYIN_ToolsAdConvertQueryV2PromotionContent          ToolsAdConvertQueryV2PromotionContent = "DOUYIN"
	THIRD_PARTY_ToolsAdConvertQueryV2PromotionContent     ToolsAdConvertQueryV2PromotionContent = "THIRD_PARTY"
	NORMAL_ToolsAdConvertQueryV2PromotionContent          ToolsAdConvertQueryV2PromotionContent = "NORMAL"
	AWEME_HOME_PAGE_ToolsAdConvertQueryV2PromotionContent ToolsAdConvertQueryV2PromotionContent = "AWEME_HOME_PAGE"
	SHOP_ToolsAdConvertQueryV2PromotionContent            ToolsAdConvertQueryV2PromotionContent = "SHOP"
)

// All allowed values of ToolsAdConvertQueryV2PromotionContent enum
var AllowedToolsAdConvertQueryV2PromotionContentEnumValues = []ToolsAdConvertQueryV2PromotionContent{
	"EXTERNAL_URL",
	"MICRO_APP",
	"DOWNLOAD_URL",
	"QUICK_APP_URL",
	"GOODS_LINK",
	"LIVE_ROOM",
	"DOUYIN",
	"THIRD_PARTY",
	"NORMAL",
	"AWEME_HOME_PAGE",
	"SHOP",
}

// NewToolsAdConvertQueryV2PromotionContentFromValue returns a pointer to a valid ToolsAdConvertQueryV2PromotionContent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsAdConvertQueryV2PromotionContentFromValue(v string) (*ToolsAdConvertQueryV2PromotionContent, error) {
	ev := ToolsAdConvertQueryV2PromotionContent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsAdConvertQueryV2PromotionContent: valid values are %v", v, AllowedToolsAdConvertQueryV2PromotionContentEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsAdConvertQueryV2PromotionContent) IsValid() bool {
	for _, existing := range AllowedToolsAdConvertQueryV2PromotionContentEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_ad_convert_query_v2_promotion_content value
func (v ToolsAdConvertQueryV2PromotionContent) Ptr() *ToolsAdConvertQueryV2PromotionContent {
	return &v
}
