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

// CgTransferQueryCanTransferBalanceV30TransferDirection
type CgTransferQueryCanTransferBalanceV30TransferDirection string

// List of cg_transfer_query_can_transfer_balance_v3.0_transfer_direction
const (
	TRANSFER_IN_CgTransferQueryCanTransferBalanceV30TransferDirection  CgTransferQueryCanTransferBalanceV30TransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferQueryCanTransferBalanceV30TransferDirection CgTransferQueryCanTransferBalanceV30TransferDirection = "TRANSFER_OUT"
)

// All allowed values of CgTransferQueryCanTransferBalanceV30TransferDirection enum
var AllowedCgTransferQueryCanTransferBalanceV30TransferDirectionEnumValues = []CgTransferQueryCanTransferBalanceV30TransferDirection{
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewCgTransferQueryCanTransferBalanceV30TransferDirectionFromValue returns a pointer to a valid CgTransferQueryCanTransferBalanceV30TransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferQueryCanTransferBalanceV30TransferDirectionFromValue(v string) (*CgTransferQueryCanTransferBalanceV30TransferDirection, error) {
	ev := CgTransferQueryCanTransferBalanceV30TransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferQueryCanTransferBalanceV30TransferDirection: valid values are %v", v, AllowedCgTransferQueryCanTransferBalanceV30TransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferQueryCanTransferBalanceV30TransferDirection) IsValid() bool {
	for _, existing := range AllowedCgTransferQueryCanTransferBalanceV30TransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_query_can_transfer_balance_v3.0_transfer_direction value
func (v CgTransferQueryCanTransferBalanceV30TransferDirection) Ptr() *CgTransferQueryCanTransferBalanceV30TransferDirection {
	return &v
}
