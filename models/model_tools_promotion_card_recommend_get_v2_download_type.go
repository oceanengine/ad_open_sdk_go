/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ToolsPromotionCardRecommendGetV2DownloadType
type ToolsPromotionCardRecommendGetV2DownloadType string

// List of tools_promotion_card_recommend_get_v2_download_type
const (
	DOWNLOAD_URL_ToolsPromotionCardRecommendGetV2DownloadType  ToolsPromotionCardRecommendGetV2DownloadType = "DOWNLOAD_URL"
	EXTERNAL_URL_ToolsPromotionCardRecommendGetV2DownloadType  ToolsPromotionCardRecommendGetV2DownloadType = "EXTERNAL_URL"
	QUICK_APP_URL_ToolsPromotionCardRecommendGetV2DownloadType ToolsPromotionCardRecommendGetV2DownloadType = "QUICK_APP_URL"
)

// All allowed values of ToolsPromotionCardRecommendGetV2DownloadType enum
var AllowedToolsPromotionCardRecommendGetV2DownloadTypeEnumValues = []ToolsPromotionCardRecommendGetV2DownloadType{
	"DOWNLOAD_URL",
	"EXTERNAL_URL",
	"QUICK_APP_URL",
}

// NewToolsPromotionCardRecommendGetV2DownloadTypeFromValue returns a pointer to a valid ToolsPromotionCardRecommendGetV2DownloadType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewToolsPromotionCardRecommendGetV2DownloadTypeFromValue(v string) (*ToolsPromotionCardRecommendGetV2DownloadType, error) {
	ev := ToolsPromotionCardRecommendGetV2DownloadType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ToolsPromotionCardRecommendGetV2DownloadType: valid values are %v", v, AllowedToolsPromotionCardRecommendGetV2DownloadTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ToolsPromotionCardRecommendGetV2DownloadType) IsValid() bool {
	for _, existing := range AllowedToolsPromotionCardRecommendGetV2DownloadTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tools_promotion_card_recommend_get_v2_download_type value
func (v ToolsPromotionCardRecommendGetV2DownloadType) Ptr() *ToolsPromotionCardRecommendGetV2DownloadType {
	return &v
}
