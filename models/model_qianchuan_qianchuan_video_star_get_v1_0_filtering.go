/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanQianchuanVideoStarGetV10Filtering
type QianchuanQianchuanVideoStarGetV10Filtering struct {
	// 需拉取视频的抖音号，默认查询全部，数量限制xxx个 注意：这里的抖音号为在星图侧adv绑定的抖音号
	AwemeIds []int64 `json:"aweme_ids,omitempty"`
	// 抖音主页视频id，限制0-50 注意：material_ids、aweme_item_id、aweme_item_title_url只能选择一个进行过滤，否则可能会查询不到数据
	AwemeItemIds []int64 `json:"aweme_item_ids,omitempty"`
	// 抖音主页视频id，限制0-50 注意：material_ids、aweme_item_id、aweme_item_title_url只能选择一个进行过滤，否则可能会查询不到数据
	AwemeItemTitleUrl *string `json:"aweme_item_title_url,omitempty"`
	// 素材id，抖音主页视频用来投放才会有，限制0-50 注意：material_ids、aweme_item_id、aweme_item_title_url只能选择一个进行过滤，否则可能会查询不到数据
	MaterialIds []int64 `json:"material_ids,omitempty"`
}
