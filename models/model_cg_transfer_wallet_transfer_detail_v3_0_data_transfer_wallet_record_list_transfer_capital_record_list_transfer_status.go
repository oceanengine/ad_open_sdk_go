/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus
type CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus string

// List of cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_capital_record_list_transfer_status
const (
	NO_TRANSFER_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus      CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus = "NO_TRANSFER"
	TRANSFER_FAILED_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus  CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_FAILED"
	TRANSFER_ING_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus     CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_ING"
	TRANSFER_PART_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus    CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_PART"
	TRANSFER_SUCCESS_CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus = "TRANSFER_SUCCESS"
)

// Ptr returns reference to cg_transfer_wallet_transfer_detail_v3.0_data_transfer_wallet_record_list_transfer_capital_record_list_transfer_status value
func (v CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus) Ptr() *CgTransferWalletTransferDetailV30DataTransferWalletRecordListTransferCapitalRecordListTransferStatus {
	return &v
}
