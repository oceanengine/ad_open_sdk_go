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

// CgTransferQueryTransferDetailV30DataTransferStatus
type CgTransferQueryTransferDetailV30DataTransferStatus string

// List of cg_transfer_query_transfer_detail_v3.0_data_transfer_status
const (
	NO_TRANSFER_CgTransferQueryTransferDetailV30DataTransferStatus      CgTransferQueryTransferDetailV30DataTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferQueryTransferDetailV30DataTransferStatus  CgTransferQueryTransferDetailV30DataTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferQueryTransferDetailV30DataTransferStatus     CgTransferQueryTransferDetailV30DataTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferQueryTransferDetailV30DataTransferStatus    CgTransferQueryTransferDetailV30DataTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferQueryTransferDetailV30DataTransferStatus CgTransferQueryTransferDetailV30DataTransferStatus = "TRANSFER_SUCCESS"
)

// All allowed values of CgTransferQueryTransferDetailV30DataTransferStatus enum
var AllowedCgTransferQueryTransferDetailV30DataTransferStatusEnumValues = []CgTransferQueryTransferDetailV30DataTransferStatus{
	"NO_TRANSFER",
	"TRANSFER_FAILED",
	"TRANSFER_ING",
	"TRANSFER_PART",
	"TRANSFER_SUCCESS",
}

// NewCgTransferQueryTransferDetailV30DataTransferStatusFromValue returns a pointer to a valid CgTransferQueryTransferDetailV30DataTransferStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryTransferDetailV30DataTransferStatusFromValue(v string) (*CgTransferQueryTransferDetailV30DataTransferStatus, error) {
	ev := CgTransferQueryTransferDetailV30DataTransferStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryTransferDetailV30DataTransferStatus: valid values are %v", v, AllowedCgTransferQueryTransferDetailV30DataTransferStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryTransferDetailV30DataTransferStatus) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryTransferDetailV30DataTransferStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_transfer_detail_v3.0_data_transfer_status value
func (v CgTransferQueryTransferDetailV30DataTransferStatus) Ptr() *CgTransferQueryTransferDetailV30DataTransferStatus {
	return &v
}
