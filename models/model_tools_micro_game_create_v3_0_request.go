/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsMicroGameCreateV30Request struct for ToolsMicroGameCreateV30Request
type ToolsMicroGameCreateV30Request struct {
	//
	AdvertiserId int64                           `json:"advertiser_id"`
	AgeLimit     ToolsMicroGameCreateV30AgeLimit `json:"age_limit"`
	//
	AppId string `json:"app_id"`
	// 画风分类名称
	ArtStyle string `json:"art_style"`
	// 折扣比例
	DiscountRate *int64 `json:"discount_rate,omitempty"`
	// 特殊标签名称
	FeatureTags []string                               `json:"feature_tags"`
	GameLink    ToolsMicroGameCreateV30RequestGameLink `json:"game_link"`
	// 是否有折扣
	HasDiscount *bool `json:"has_discount,omitempty"`
	// 是否包含网赚内容
	HasOnlineEarning *bool `json:"has_online_earning,omitempty"`
	//
	Introduction        string                                      `json:"introduction"`
	MaxPaymentTierRange *ToolsMicroGameCreateV30MaxPaymentTierRange `json:"max_payment_tier_range,omitempty"`
	MidPaymentTierRange *ToolsMicroGameCreateV30MidPaymentTierRange `json:"mid_payment_tier_range,omitempty"`
	MinPaymentTierRange *ToolsMicroGameCreateV30MinPaymentTierRange `json:"min_payment_tier_range,omitempty"`
	// 网络环境名称
	NetworkEnvironment []string `json:"network_environment"`
	//
	Remark       string                              `json:"remark"`
	RevenueModel ToolsMicroGameCreateV30RevenueModel `json:"revenue_model"`
	// 周卡/月卡/季卡种类
	ScheduleCards []*ToolsMicroGameCreateV30ScheduleCards `json:"schedule_cards,omitempty"`
	// 细分类型
	TagInfo string `json:"tag_info"`
	// 题材标签名称
	ThemeTag string `json:"theme_tag"`
}
