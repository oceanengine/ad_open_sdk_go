/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfo
type StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfo struct {
	// 小程序cps 分账周期
	AccountDivideDay *int64 `json:"account_divide_day,omitempty"`
	//
	Brand *string `json:"brand,omitempty"`
	//
	ChallengeStatus *int64 `json:"challenge_status,omitempty"`
	//
	CommissionRate *int64 `json:"commission_rate,omitempty"`
	//
	CommissionRateIaap *string `json:"commission_rate_iaap,omitempty"`
	//
	CreateTime string `json:"create_time"`
	// 任务介绍
	DemandDesc *string `json:"demand_desc,omitempty"`
	//
	DemandIcon *string `json:"demand_icon,omitempty"`
	//
	DemandId int64 `json:"demand_id"`
	//
	DemandName string `json:"demand_name"`
	//
	EvaluateType *int64 `json:"evaluate_type,omitempty"`
	//
	ExpirationTime *string `json:"expiration_time,omitempty"`
	//
	ExpirationTimeEnd *string `json:"expiration_time_end,omitempty"`
	//
	FirstClassCategory *int64 `json:"first_class_category,omitempty"`
	//
	HighProfitChallenge *bool `json:"high_profit_challenge,omitempty"`
	//
	McnProfitRate *int64 `json:"mcn_profit_rate,omitempty"`
	//
	PayType *int64 `json:"pay_type,omitempty"`
	//
	PlayletCategory *int64 `json:"playlet_category,omitempty"`
	//
	PlayletFirstWeek *int64 `json:"playlet_first_week,omitempty"`
	//
	PlayletGender []int64 `json:"playlet_gender,omitempty"`
	//
	PlayletHot *bool `json:"playlet_hot,omitempty"`
	//
	PlayletSecondWeek *int64 `json:"playlet_second_week,omitempty"`
	//
	PlayletTheme []int64 `json:"playlet_theme,omitempty"`
	//
	ProductCategory *int64 `json:"product_category,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProductReleaseTime *string `json:"product_release_time,omitempty"`
	//
	ProfitRateChannel *map[string]int64 `json:"profit_rate_channel,omitempty"`
	//
	SecondClassCategory         *int64                                                                               `json:"second_class_category,omitempty"`
	SmallAppInfo                *StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfoSmallAppInfo                `json:"small_app_info,omitempty"`
	StarMicroGameUniteExtraInfo *StarMcnProviderGetTaskDetailV2ResponseDataTaskDetailInfoStarMicroGameUniteExtraInfo `json:"star_micro_game_unite_extra_info,omitempty"`
	// 任务头图
	TaskHeadImage *string `json:"task_head_image,omitempty"`
	//
	VideoType int64 `json:"video_type"`
}
