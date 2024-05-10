/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AnalyticsAttributionV30RequestContextAd
type AnalyticsAttributionV30RequestContextAd struct {
	// 广告计划id；若您可以确定当前订单来自对应广告计划id，则可以上报此字段获得更加精准的归因结果
	AdId *int64 `json:"ad_id,omitempty"`
	// 广告账户id；同上，若您可以确定当前订单来自对应广告账户id，则可以上报此字段获得更加精准的归因结果
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
	// false: 客户未归因，上报为需要巨量引擎归因的事件
	Attributed *bool `json:"attributed,omitempty"`
	// 广告组id；同上，若您可以确定当前订单来自对应广告组id，则可以上报此字段获得更加精准的归因结果
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 点击时间，指用户点击页面跳转时间，上报秒级时间戳
	ClickTime *int64 `json:"click_time,omitempty"`
	// 客户id；若您可以确定当前订单来自对应客户id，则可以上报此字段获得更加精准的归因结果
	CustomerIds []int64 `json:"customer_ids,omitempty"`
}
