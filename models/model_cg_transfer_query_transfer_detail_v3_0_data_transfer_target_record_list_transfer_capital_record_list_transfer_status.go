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

// CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus
type CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus string

// List of cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_capital_record_list_transfer_status
const (
	NO_TRANSFER_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus      CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus  CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus     CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus    CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus enum
var AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatusEnumValues = []CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatusFromValue returns a pointer to a valid CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatusFromValue(v string) (*CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus, error) {
	ev := CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus: valid values are %v", v, AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_detail_v3.0_data_transfer_target_record_list_transfer_capital_record_list_transfer_status value
func (v CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus) Ptr() *CgTransferQueryTransferDetailV30DataTransferTargetRecordListTransferCapitalRecordListTransferStatus {
	return &v
}
