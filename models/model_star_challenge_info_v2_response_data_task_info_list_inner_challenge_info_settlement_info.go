/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoSettlementInfo
type StarChallengeInfoV2ResponseDataTaskInfoListInnerChallengeInfoSettlementInfo struct {
	//
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	//
	AndroidConvertId *int64 `json:"android_convert_id,omitempty"`
	//
	AutoAddBudgetTimes *int64 `json:"auto_add_budget_times,omitempty"`
	//
	AutoAddBudgetTriggerRatio *int64 `json:"auto_add_budget_trigger_ratio,omitempty"`
	//
	EvaluateType int64 `json:"evaluate_type"`
	//
	IosConvertId *int64 `json:"ios_convert_id,omitempty"`
	//
	MaxBudgetAddAmount *int64 `json:"max_budget_add_amount,omitempty"`
	//
	RewardRule *string `json:"reward_rule,omitempty"`
	//
	SupplementInfo *string `json:"supplement_info,omitempty"`
	//
	UnitPrice int64 `json:"unit_price"`
}
