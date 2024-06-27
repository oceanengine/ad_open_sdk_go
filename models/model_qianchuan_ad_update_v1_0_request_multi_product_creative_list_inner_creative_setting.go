/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdUpdateV10RequestMultiProductCreativeListInnerCreativeSetting
type QianchuanAdUpdateV10RequestMultiProductCreativeListInnerCreativeSetting struct {
	//
	AdKeywords []string `json:"ad_keywords,omitempty"`
	// 抖音号id
	AwemeId              *int64                                                                                   `json:"aweme_id,omitempty"`
	CreativeAutoGenerate *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingCreativeAutoGenerate         `json:"creative_auto_generate,omitempty"`
	CreativeCategory     *QianchuanAdUpdateV10RequestMultiProductCreativeListInnerCreativeSettingCreativeCategory `json:"creative_category,omitempty"`
	DynamicCreative      *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingDynamicCreative              `json:"dynamic_creative,omitempty"`
	IsHomepageHide       *QianchuanAdUpdateV10MultiProductCreativeListCreativeSettingIsHomepageHide               `json:"is_homepage_hide,omitempty"`
}
