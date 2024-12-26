/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalCxtDetailV30DataStatus
type LocalCxtDetailV30DataStatus string

// List of local_cxt_detail_v3.0_data_status
const (
	ADVERTISER_OFFLINE_BUDGET_LocalCxtDetailV30DataStatus     LocalCxtDetailV30DataStatus = "ADVERTISER_OFFLINE_BUDGET"
	ADVERTISER_PRE_OFFLINE_BUDGET_LocalCxtDetailV30DataStatus LocalCxtDetailV30DataStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	CREATE_LocalCxtDetailV30DataStatus                        LocalCxtDetailV30DataStatus = "CREATE"
	ENABLE_LocalCxtDetailV30DataStatus                        LocalCxtDetailV30DataStatus = "ENABLE"
	EXTERNAL_URL_DISABLE_LocalCxtDetailV30DataStatus          LocalCxtDetailV30DataStatus = "EXTERNAL_URL_DISABLE"
	FROZEN_LocalCxtDetailV30DataStatus                        LocalCxtDetailV30DataStatus = "FROZEN"
	OFFLINE_AUDIT_LocalCxtDetailV30DataStatus                 LocalCxtDetailV30DataStatus = "OFFLINE_AUDIT"
	OFFLINE_BALANCE_LocalCxtDetailV30DataStatus               LocalCxtDetailV30DataStatus = "OFFLINE_BALANCE"
	PAUSED_LocalCxtDetailV30DataStatus                        LocalCxtDetailV30DataStatus = "PAUSED"
	SYSTEM_DISABLE_LocalCxtDetailV30DataStatus                LocalCxtDetailV30DataStatus = "SYSTEM_DISABLE"
)

// Ptr returns reference to local_cxt_detail_v3.0_data_status value
func (v LocalCxtDetailV30DataStatus) Ptr() *LocalCxtDetailV30DataStatus {
	return &v
}