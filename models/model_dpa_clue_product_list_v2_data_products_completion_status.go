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

// DpaClueProductListV2DataProductsCompletionStatus
type DpaClueProductListV2DataProductsCompletionStatus string

// List of dpa_clue_product_list_v2_data_products_completion_status
const (
	AD_COMPLETED_DpaClueProductListV2DataProductsCompletionStatus    DpaClueProductListV2DataProductsCompletionStatus = "AD_COMPLETED"
	ALL_COMPLETED_DpaClueProductListV2DataProductsCompletionStatus   DpaClueProductListV2DataProductsCompletionStatus = "ALL_COMPLETED"
	LEADS_COMPLETED_DpaClueProductListV2DataProductsCompletionStatus DpaClueProductListV2DataProductsCompletionStatus = "LEADS_COMPLETED"
	TO_BE_COMPLETED_DpaClueProductListV2DataProductsCompletionStatus DpaClueProductListV2DataProductsCompletionStatus = "TO_BE_COMPLETED"
)

// All allowed values of DpaClueProductListV2DataProductsCompletionStatus enum
var AllowedDpaClueProductListV2DataProductsCompletionStatusEnumValues = []DpaClueProductListV2DataProductsCompletionStatus{
	"AD_COMPLETED",
	"ALL_COMPLETED",
	"LEADS_COMPLETED",
	"TO_BE_COMPLETED",
}

// NewDpaClueProductListV2DataProductsCompletionStatusFromValue returns a pointer to a valid DpaClueProductListV2DataProductsCompletionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDpaClueProductListV2DataProductsCompletionStatusFromValue(v string) (*DpaClueProductListV2DataProductsCompletionStatus, error) {
	ev := DpaClueProductListV2DataProductsCompletionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DpaClueProductListV2DataProductsCompletionStatus: valid values are %v", v, AllowedDpaClueProductListV2DataProductsCompletionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DpaClueProductListV2DataProductsCompletionStatus) IsValid() bool {
	for _, existing := range AllowedDpaClueProductListV2DataProductsCompletionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dpa_clue_product_list_v2_data_products_completion_status value
func (v DpaClueProductListV2DataProductsCompletionStatus) Ptr() *DpaClueProductListV2DataProductsCompletionStatus {
	return &v
}
