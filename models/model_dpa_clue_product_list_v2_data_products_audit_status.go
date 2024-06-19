/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DpaClueProductListV2DataProductsAuditStatus
type DpaClueProductListV2DataProductsAuditStatus string

// List of dpa_clue_product_list_v2_data_products_audit_status
const (
	AUDIT_STATUS_APPROVE_DpaClueProductListV2DataProductsAuditStatus DpaClueProductListV2DataProductsAuditStatus = "AUDIT_STATUS_APPROVE"
	AUDIT_STATUS_INIT_DpaClueProductListV2DataProductsAuditStatus    DpaClueProductListV2DataProductsAuditStatus = "AUDIT_STATUS_INIT"
	AUDIT_STATUS_REJECT_DpaClueProductListV2DataProductsAuditStatus  DpaClueProductListV2DataProductsAuditStatus = "AUDIT_STATUS_REJECT"
)

// All allowed values of DpaClueProductListV2DataProductsAuditStatus enum
var AllowedDpaClueProductListV2DataProductsAuditStatusEnumValues = []DpaClueProductListV2DataProductsAuditStatus{
	"AUDIT_STATUS_APPROVE",
	"AUDIT_STATUS_INIT",
	"AUDIT_STATUS_REJECT",
}

// NewDpaClueProductListV2DataProductsAuditStatusFromValue returns a pointer to a valid DpaClueProductListV2DataProductsAuditStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2DataProductsAuditStatusFromValue(v string) (*DpaClueProductListV2DataProductsAuditStatus, error) {
	ev := DpaClueProductListV2DataProductsAuditStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2DataProductsAuditStatus: valid values are %v", v, AllowedDpaClueProductListV2DataProductsAuditStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2DataProductsAuditStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2DataProductsAuditStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_data_products_audit_status value
func (v DpaClueProductListV2DataProductsAuditStatus) Ptr() *DpaClueProductListV2DataProductsAuditStatus {
	return &v
}
