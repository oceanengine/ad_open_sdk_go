/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarMcnProviderGetUnparticipatedTaskV2ResponseDataDataInner struct for StarMcnProviderGetUnparticipatedTaskV2ResponseDataDataInner
type StarMcnProviderGetUnparticipatedTaskV2ResponseDataDataInner struct {
	//
	CommissionRate *int64 `json:"commission_rate,omitempty"`
	//
	CommissionRateIaap *string `json:"commission_rate_iaap,omitempty"`
	//
	ComponentType *int64 `json:"component_type,omitempty"`
	//
	CreateTime string `json:"create_time"`
	//
	DemandIcon *string `json:"demand_icon,omitempty"`
	//
	DemandId int64 `json:"demand_id"`
	//
	DemandName string `json:"demand_name"`
	//
	EvaluateType *int32 `json:"evaluate_type,omitempty"`
	//
	ExpirationTime *string `json:"expiration_time,omitempty"`
	//
	ExpirationTimeEnd *string `json:"expiration_time_end,omitempty"`
	//
	FirstCategoryId *int64 `json:"first_category_id,omitempty"`
	//
	FirstClassCategory *int64 `json:"first_class_category,omitempty"`
	//
	ParticipateCope *int64 `json:"participate_cope,omitempty"`
	//
	PayType *int64 `json:"pay_type,omitempty"`
	// 任务单价 CPM CPA
	Price *int64 `json:"price,omitempty"`
	//
	ProductCategory *int64 `json:"product_category,omitempty"`
	//
	SecondClassCategory *int64                                                                   `json:"second_class_category,omitempty"`
	SmallAppInfo        *StarMcnProviderGetUnparticipatedTaskV2ResponseDataDataInnerSmallAppInfo `json:"small_app_info,omitempty"`
	//
	TaskSupportAd *bool `json:"task_support_ad,omitempty"`
}
