/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// CgTransferWalletTransferDetailV30DataTransferDirection
type CgTransferWalletTransferDetailV30DataTransferDirection string

// List of cg_transfer_wallet_transfer_detail_v3.0_data_transfer_direction
const (
	TRANSFER_IN_CgTransferWalletTransferDetailV30DataTransferDirection  CgTransferWalletTransferDetailV30DataTransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferWalletTransferDetailV30DataTransferDirection CgTransferWalletTransferDetailV30DataTransferDirection = "TRANSFER_OUT"
)

// All allowed values of CgTransferWalletTransferDetailV30DataTransferDirection enum
var AllowedCgTransferWalletTransferDetailV30DataTransferDirectionEnumValues = []CgTransferWalletTransferDetailV30DataTransferDirection{
	"TRANSFER_IN",
	"TRANSFER_OUT",
}

// NewCgTransferWalletTransferDetailV30DataTransferDirectionFromValue returns a pointer to a valid CgTransferWalletTransferDetailV30DataTransferDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCgTransferWalletTransferDetailV30DataTransferDirectionFromValue(v string) (*CgTransferWalletTransferDetailV30DataTransferDirection, error) {
	ev := CgTransferWalletTransferDetailV30DataTransferDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CgTransferWalletTransferDetailV30DataTransferDirection: valid values are %v", v, AllowedCgTransferWalletTransferDetailV30DataTransferDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CgTransferWalletTransferDetailV30DataTransferDirection) IsValid() bool {
	for _, existing := range AllowedCgTransferWalletTransferDetailV30DataTransferDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cg_transfer_wallet_transfer_detail_v3.0_data_transfer_direction value
func (v CgTransferWalletTransferDetailV30DataTransferDirection) Ptr() *CgTransferWalletTransferDetailV30DataTransferDirection {
	return &v
}
