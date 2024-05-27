/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerCreativeSetting
type QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerCreativeSetting struct {
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// 抖音号id
	AwemeId              *int64                                                                                           `json:"aweme_id,omitempty"`
	CreativeAutoGenerate *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingCreativeAutoGenerate          `json:"creative_auto_generate,omitempty"`
	CreativeCategory     *QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerCreativeSettingCreativeCategory `json:"creative_category,omitempty"`
	DynamicCreative      *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingDynamicCreative               `json:"dynamic_creative,omitempty"`
	IsHomepageHide       *QianchuanAdDetailGetV10DataMultiProductCreativeListCreativeSettingIsHomepageHide                `json:"is_homepage_hide,omitempty"`
}
