/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRangeAuthorWatcherTag 达人观众标签
type StarDemandCreateChallengeV2RequestChallengeInfoParticipateAuthorRangeAuthorWatcherTag struct {
	// 观众年龄范围 「投稿任务创建数据字典」中返回的age_range中的项 形如  \"18-23岁居多\"，\"24-30岁居多\"
	AgeRange []string `json:"age_range,omitempty"`
	// 观众地域 「投稿任务创建数据字典」返回的city_level中的项 形如“一线城市居多”，“二线城市居多”，“全部”
	CityLevel []string `json:"city_level,omitempty"`
	// 观众设备品牌 「投稿任务创建数据字典」返回的device_brands中的项 形如“苹果居多”，“华为居多”
	DeviceBrands []string `json:"device_brands,omitempty"`
	// 观众性别 “男”，“女”
	Gender []string `json:"gender,omitempty"`
}
