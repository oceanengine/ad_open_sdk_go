/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferTransferDetailGetV30DataTransferDirection
type CgTransferTransferDetailGetV30DataTransferDirection string

// List of cg_transfer_transfer_detail_get_v3.0_data_transfer_direction
const (
	TRANSFER_IN_CgTransferTransferDetailGetV30DataTransferDirection  CgTransferTransferDetailGetV30DataTransferDirection = "TRANSFER_IN"
	TRANSFER_OUT_CgTransferTransferDetailGetV30DataTransferDirection CgTransferTransferDetailGetV30DataTransferDirection = "TRANSFER_OUT"
)

// Ptr returns reference to cg_transfer_transfer_detail_get_v3.0_data_transfer_direction value
func (v CgTransferTransferDetailGetV30DataTransferDirection) Ptr() *CgTransferTransferDetailGetV30DataTransferDirection {
	return &v
}