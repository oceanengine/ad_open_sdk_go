/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferWalletTransferCreateV30TransferDirection
type CgTransferWalletTransferCreateV30TransferDirection string

// List of cg_transfer_wallet_transfer_create_v3.0_transfer_direction
const (
	TRANSFER_IN_CgTransferWalletTransferCreateV30TransferDirection  CgTransferWalletTransferCreateV30TransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferWalletTransferCreateV30TransferDirection CgTransferWalletTransferCreateV30TransferDirection = "TRANSFER_OUT"
)

// All allowed values of CgTransferWalletTransferCreateV30TransferDirection enum
var AllowedCgTransferWalletTransferCreateV30TransferDirectionEnumValues = []CgTransferWalletTransferCreateV30TransferDirection{
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewCgTransferWalletTransferCreateV30TransferDirectionFromValue returns a pointer to a valid CgTransferWalletTransferCreateV30TransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferCreateV30TransferDirectionFromValue(v string) (*CgTransferWalletTransferCreateV30TransferDirection, error) {
	ev := CgTransferWalletTransferCreateV30TransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferCreateV30TransferDirection: valid values are %v", v, AllowedCgTransferWalletTransferCreateV30TransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferCreateV30TransferDirection) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferCreateV30TransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_create_v3.0_transfer_direction value
func (v CgTransferWalletTransferCreateV30TransferDirection) Ptr() *CgTransferWalletTransferCreateV30TransferDirection {
	return &v
}
