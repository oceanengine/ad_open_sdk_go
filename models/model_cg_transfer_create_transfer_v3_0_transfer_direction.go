/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferCreateTransferV30TransferDirection
type CgTransferCreateTransferV30TransferDirection string

// List of cg_transfer_create_transfer_v3.0_transfer_direction
const (
	TRANSFER_IN_CgTransferCreateTransferV30TransferDirection  CgTransferCreateTransferV30TransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferCreateTransferV30TransferDirection CgTransferCreateTransferV30TransferDirection = "TRANSFER_OUT"
)

// Ptr returns reference to cg_transfer_create_transfer_v3.0_transfer_direction value
func (v CgTransferCreateTransferV30TransferDirection) Ptr() *CgTransferCreateTransferV30TransferDirection {
	return &v
}
