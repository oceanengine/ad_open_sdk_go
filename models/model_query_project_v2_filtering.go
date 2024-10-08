/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QueryProjectV2Filtering
type QueryProjectV2Filtering struct {
	// 投放账号ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 客户ID
	CustomerId *int64 `json:"customer_id,omitempty"`
	// 应回款日期
	Deadline *string `json:"deadline,omitempty"`
	// 平台列表
	PlatformList []*QueryProjectV2FilteringPlatformList `json:"platform_list,omitempty"`
	// 项目结束时间-开始 例如 2023-07-06
	ProjectEndDateBegin *string `json:"project_end_date_begin,omitempty"`
	// 项目结束时间-结束 例如 2023-07-06
	ProjectEndDateEnd *string `json:"project_end_date_end,omitempty"`
	// 项目开始时间-开始 例如 2023-07-06
	ProjectStartDateBegin *string `json:"project_start_date_begin,omitempty"`
	// 项目开始时间-结束 例如 2023-07-06
	ProjectStartDateEnd *string `json:"project_start_date_end,omitempty"`
	// 项目状态
	ProjectStatusList []*QueryProjectV2FilteringProjectStatusList `json:"project_status_list,omitempty"`
	// 项目回款状态
	ReceiptStatusList []*QueryProjectV2FilteringReceiptStatusList `json:"receipt_status_list,omitempty"`
	// 投放类型
	ServingTypeList []*QueryProjectV2FilteringServingTypeList `json:"serving_type_list,omitempty"`
}
