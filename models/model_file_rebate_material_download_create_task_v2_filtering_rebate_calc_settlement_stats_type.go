/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType
type FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType string

// List of file_rebate_material_download_create_task_v2_filtering_rebate_calc_settlement_stats_type
const (
	NORMAL_INDUSTRY_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "NORMAL_INDUSTRY"
	E_COMMERCE_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType      FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "E_COMMERCE"
	TASK_INCENTIVES_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "TASK_INCENTIVES"
)

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_rebate_calc_settlement_stats_type value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType {
	return &v
}
