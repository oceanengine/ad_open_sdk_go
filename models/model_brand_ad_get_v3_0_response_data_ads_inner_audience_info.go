/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandAdGetV30ResponseDataAdsInnerAudienceInfo 定向信息
type BrandAdGetV30ResponseDataAdsInnerAudienceInfo struct {
	// 网络
	Ac   []int64                               `json:"ac,omitempty"`
	AcV2 *BrandAdGetV30DataAdsAudienceInfoAcV2 `json:"ac_v2,omitempty"`
	// 新版兴趣(行为)定向
	ActionCategories []int64 `json:"action_categories,omitempty"`
	// 新版兴趣(行为)定向支持的层级  ['1']、['2']
	ActionCategoriesLevel []string `json:"action_categories_level,omitempty"`
	// 兴趣分类
	AdTag []int64 `json:"ad_tag,omitempty"`
	// 年龄 [[0, 17], [18, 23]]
	Age         []*[]int64                                   `json:"age,omitempty"`
	BrandSafety *BrandAdGetV30DataAdsAudienceInfoBrandSafety `json:"brand_safety,omitempty"`
	// 商圈
	BusinessIds []int64 `json:"business_ids,omitempty"`
	// 地域定向 - 城市
	City []int64 `json:"city,omitempty"`
	// 内容关键词
	ContentKeyWords []string `json:"content_key_words,omitempty"`
	// 人群包
	DisplayRetargetingTags []int64 `json:"display_retargeting_tags,omitempty"`
	// 地域类型,all:不限;select:城市;local:商圈
	District      *string                                        `json:"district,omitempty"`
	DistrictSplit *BrandAdGetV30DataAdsAudienceInfoDistrictSplit `json:"district_split,omitempty"`
	DistrictType  *BrandAdGetV30DataAdsAudienceInfoDistrictType  `json:"district_type,omitempty"`
	// 频道标识
	Entry           []int64                                          `json:"entry,omitempty"`
	ExcludeDistrict *BrandAdGetV30DataAdsAudienceInfoExcludeDistrict `json:"exclude_district,omitempty"`
	// 性别
	Gender *int64 `json:"gender,omitempty"`
	// 商圈地理位置信息
	Geolocation []*BrandAdGetV30ResponseDataAdsInnerAudienceInfoGeolocationInner `json:"geolocation,omitempty"`
	// 商圈拆分
	IsBusinessSplit *int64 `json:"is_business_split,omitempty"`
	// 平台
	Platform   []int64                                     `json:"platform,omitempty"`
	PlatformV2 *BrandAdGetV30DataAdsAudienceInfoPlatformV2 `json:"platform_v2,omitempty"`
	// 地域定向 - 省份
	Province []int64 `json:"province,omitempty"`
	// 人群包
	RetargetingTagType *map[string]BrandAdGetV30ResponseDataAdsInnerAudienceInfoRetargetingTagTypeValue `json:"retargeting_tag_type,omitempty"`
	// 人群包
	RetargetingTags []int64                                          `json:"retargeting_tags,omitempty"`
	RetargetingType *BrandAdGetV30DataAdsAudienceInfoRetargetingType `json:"retargeting_type,omitempty"`
	// 搜索关键词
	SearchKeyword []string `json:"search_keyword,omitempty"`
	// 词包Id
	WordBagIds []int64 `json:"word_bag_ids,omitempty"`
	// 询价单关联的词条列表
	WordList []string `json:"word_list,omitempty"`
}
