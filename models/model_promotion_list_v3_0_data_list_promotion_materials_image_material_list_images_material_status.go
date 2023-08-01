/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus
type PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus string

// List of promotion_list_v3.0_data_list_promotion_materials_image_material_list_images_material_status
const (
	MATERIAL_STATUS_ADV_OFFLINE_BUDGET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus        PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_ADV_OFFLINE_BUDGET"
	MATERIAL_STATUS_ADV_PRE_OFFLINE_BUDGET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus    PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_ADV_PRE_OFFLINE_BUDGET"
	MATERIAL_STATUS_AUDIT_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                     PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_AUDIT"
	MATERIAL_STATUS_AWEME_ANCHOR_DISABLED_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus     PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_AWEME_ANCHOR_DISABLED"
	MATERIAL_STATUS_DELETE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                    PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_DELETE"
	MATERIAL_STATUS_DISABLE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                   PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_DISABLE"
	MATERIAL_STATUS_ERROR_DEFAULT_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus             PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_ERROR_DEFAULT"
	MATERIAL_STATUS_FROZEN_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                    PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_FROZEN"
	MATERIAL_STATUS_MATERIAL_DELETE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus           PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_MATERIAL_DELETE"
	MATERIAL_STATUS_NO_SCHEDULE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus               PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_NO_SCHEDULE"
	MATERIAL_STATUS_OFFLINE_AUDIT_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus             PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_OFFLINE_AUDIT"
	MATERIAL_STATUS_OFFLINE_BALANCE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus           PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_OFFLINE_BALANCE"
	MATERIAL_STATUS_OFFLINE_BUDGET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus            PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_OFFLINE_BUDGET"
	MATERIAL_STATUS_OK_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                        PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_OK"
	MATERIAL_STATUS_PREOFFLINE_BUGDET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus         PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PREOFFLINE_BUGDET"
	MATERIAL_STATUS_PREONLINE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                 PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PREONLINE"
	MATERIAL_STATUS_PRODUCT_OFFLINE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus           PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PRODUCT_OFFLINE"
	MATERIAL_STATUS_PROJECT_DISABLE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus           PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PROJECT_DISABLE"
	MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus    PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET"
	MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET"
	MATERIAL_STATUS_PROMOTION_DISABLE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus         PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PROMOTION_DISABLE"
	MATERIAL_STATUS_PROMOTION_QUOTA_LIMIT_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus     PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_PROMOTION_QUOTA_LIMIT"
	MATERIAL_STATUS_REAUDIT_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                   PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_REAUDIT"
	MATERIAL_STATUS_TIME_DONE_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus                 PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_TIME_DONE"
	MATERIAL_STATUS_TIME_NO_REACH_PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus             PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus = "MATERIAL_STATUS_TIME_NO_REACH"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus enum
var AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatusEnumValues = []PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus{
	"MATERIAL_STATUS_ADV_OFFLINE_BUDGET",
	"MATERIAL_STATUS_ADV_PRE_OFFLINE_BUDGET",
	"MATERIAL_STATUS_AUDIT",
	"MATERIAL_STATUS_AWEME_ANCHOR_DISABLED",
	"MATERIAL_STATUS_DELETE",
	"MATERIAL_STATUS_DISABLE",
	"MATERIAL_STATUS_ERROR_DEFAULT",
	"MATERIAL_STATUS_FROZEN",
	"MATERIAL_STATUS_MATERIAL_DELETE",
	"MATERIAL_STATUS_NO_SCHEDULE",
	"MATERIAL_STATUS_OFFLINE_AUDIT",
	"MATERIAL_STATUS_OFFLINE_BALANCE",
	"MATERIAL_STATUS_OFFLINE_BUDGET",
	"MATERIAL_STATUS_OK",
	"MATERIAL_STATUS_PREOFFLINE_BUGDET",
	"MATERIAL_STATUS_PREONLINE",
	"MATERIAL_STATUS_PRODUCT_OFFLINE",
	"MATERIAL_STATUS_PROJECT_DISABLE",
	"MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET",
	"MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET",
	"MATERIAL_STATUS_PROMOTION_DISABLE",
	"MATERIAL_STATUS_PROMOTION_QUOTA_LIMIT",
	"MATERIAL_STATUS_REAUDIT",
	"MATERIAL_STATUS_TIME_DONE",
	"MATERIAL_STATUS_TIME_NO_REACH",
}

// NewPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatusFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatusFromValue(v string) (*PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus, error) {
	ev := PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_image_material_list_images_material_status value
func (v PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus) Ptr() *PromotionListV30DataListPromotionMaterialsImageMaterialListImagesMaterialStatus {
	return &v
}
