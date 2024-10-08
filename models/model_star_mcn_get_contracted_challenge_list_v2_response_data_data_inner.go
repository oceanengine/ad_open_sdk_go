/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnGetContractedChallengeListV2ResponseDataDataInner struct for StarMcnGetContractedChallengeListV2ResponseDataDataInner
type StarMcnGetContractedChallengeListV2ResponseDataDataInner struct {
	//
	AccountDivideDay *int64 `json:"account_divide_day,omitempty"`
	//
	Brand *string `json:"brand,omitempty"`
	//
	ChallengeStatus *int32 `json:"challenge_status,omitempty"`
	//
	CommissionRate *int64 `json:"commission_rate,omitempty"`
	//
	CommissionRateIaap *string `json:"commission_rate_iaap,omitempty"`
	//
	CreateTime string `json:"create_time"`
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
	IaaIncome *int64 `json:"iaa_income,omitempty"`
	//
	ItemComponentType *int64 `json:"item_component_type,omitempty"`
	//
	McnProfitRate *int64 `json:"mcn_profit_rate,omitempty"`
	//
	PayType *int64 `json:"pay_type,omitempty"`
	//
	Price *int64 `json:"price,omitempty"`
	//
	ProductCategory *int32 `json:"product_category,omitempty"`
	//
	ProductName *string `json:"product_name,omitempty"`
	//
	ProfitRateChannel *map[string]int64 `json:"profit_rate_channel,omitempty"`
	//
	RealCost *int64 `json:"real_cost,omitempty"`
	//
	SecondClassCategory *int64                                                                `json:"second_class_category,omitempty"`
	SmallAppInfo        *StarMcnGetContractedChallengeListV2ResponseDataDataInnerSmallAppInfo `json:"small_app_info,omitempty"`
	//
	TotalBudget *int64 `json:"total_budget,omitempty"`
	//
	VideoType int64 `json:"video_type"`
}
