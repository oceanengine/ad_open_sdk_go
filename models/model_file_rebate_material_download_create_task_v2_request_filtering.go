/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialDownloadCreateTaskV2RequestFiltering 过滤条件
type FileRebateMaterialDownloadCreateTaskV2RequestFiltering struct {
	// 广告主id
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 客户名称
	CustomerName *string `json:"customer_name,omitempty"`
	// 素材首投所在月 格式：yyyy-MM
	FirstEffectivePeriod *string                                                          `json:"first_effective_period,omitempty"`
	IsAccumulation       *FileRebateMaterialDownloadCreateTaskV2FilteringIsAccumulation   `json:"is_accumulation,omitempty"`
	IsLiveRebateType     *FileRebateMaterialDownloadCreateTaskV2FilteringIsLiveRebateType `json:"is_live_rebate_type,omitempty"`
	IsValidMaterial      *FileRebateMaterialDownloadCreateTaskV2FilteringIsValidMaterial  `json:"is_valid_material,omitempty"`
	// 素材首投日期范围结束日期 格式：yyyy-MM-dd
	MaterialFirstEffectiveEndDate *string `json:"material_first_effective_end_date,omitempty"`
	// 素材首投日期范围开始日期 格式：yyyy-MM-dd
	MaterialFirstEffectiveStartDate *string                                                             `json:"material_first_effective_start_date,omitempty"`
	MaterialIsEffective             *FileRebateMaterialDownloadCreateTaskV2FilteringMaterialIsEffective `json:"material_is_effective,omitempty"`
	// 素材标签筛选项（如传入多个标签，取交集）
	MaterialTag []*FileRebateMaterialDownloadCreateTaskV2FilteringMaterialTag `json:"material_tag,omitempty"`
	OperatorTag *FileRebateMaterialDownloadCreateTaskV2FilteringOperatorTag   `json:"operator_tag,omitempty"`
	// 【政策粒度】累量消耗范围区间上限 - 仅支持录入整数,单位：元
	PolicyCostMax *int64 `json:"policy_cost_max,omitempty"`
	// 【政策粒度】累量消耗范围区间下限 - 仅支持录入整数,单位：元
	PolicyCostMin *int64 `json:"policy_cost_min,omitempty"`
	// 一级结算行业
	RebateCalcFirstIndustryName *string                                                              `json:"rebate_calc_first_industry_name,omitempty"`
	RebateCalcPolicyType        *FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcPolicyType `json:"rebate_calc_policy_type,omitempty"`
	// 二级结算行业
	RebateCalcSecondIndustryName  *string                                                                       `json:"rebate_calc_second_industry_name,omitempty"`
	RebateCalcSettlementStatsType *FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType `json:"rebate_calc_settlement_stats_type,omitempty"`
	// 本期已累量天数
	ThisPeriodCumDayNum *int64 `json:"this_period_cum_day_num,omitempty"`
}
