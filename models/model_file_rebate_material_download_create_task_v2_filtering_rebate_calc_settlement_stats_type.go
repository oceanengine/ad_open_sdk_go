/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType
type FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType string

// List of file_rebate_material_download_create_task_v2_filtering_rebate_calc_settlement_stats_type
const (
	NORMAL_INDUSTRY_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "NORMAL_INDUSTRY"
	E_COMMERCE_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType      FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "E_COMMERCE"
	TASK_INCENTIVES_FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType = "TASK_INCENTIVES"
)

// All allowed values of FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType enum
var AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsTypeEnumValues = []FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType{
	"NORMAL_INDUSTRY",
	"E_COMMERCE",
	"TASK_INCENTIVES",
}

// NewFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsTypeFromValue returns a pointer to a valid FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsTypeFromValue(v string) (*FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType, error) {
	ev := FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType: valid values are %v", v, AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType) IsValid() bool {
	for _, existing := range AllowedFileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to file_rebate_material_download_create_task_v2_filtering_rebate_calc_settlement_stats_type value
func (v FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType) Ptr() *FileRebateMaterialDownloadCreateTaskV2FilteringRebateCalcSettlementStatsType {
	return &v
}
