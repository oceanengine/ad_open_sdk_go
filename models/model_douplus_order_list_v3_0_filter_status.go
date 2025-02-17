/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DouplusOrderListV30FilterStatus
type DouplusOrderListV30FilterStatus string

// List of douplus_order_list_v3.0_filter_status
const (
	AUDITING_DouplusOrderListV30FilterStatus      DouplusOrderListV30FilterStatus = "AUDITING"
	AUDIT_PAUSE_DouplusOrderListV30FilterStatus   DouplusOrderListV30FilterStatus = "AUDIT_PAUSE"
	DELIVERIED_DouplusOrderListV30FilterStatus    DouplusOrderListV30FilterStatus = "DELIVERIED"
	DELIVERING_DouplusOrderListV30FilterStatus    DouplusOrderListV30FilterStatus = "DELIVERING"
	OFFLINE_AUDIT_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "OFFLINE_AUDIT"
	TIME_NO_REACH_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "TIME_NO_REACH"
	UNDELIVERIED_DouplusOrderListV30FilterStatus  DouplusOrderListV30FilterStatus = "UNDELIVERIED"
	UNPAID_DouplusOrderListV30FilterStatus        DouplusOrderListV30FilterStatus = "UNPAID"
	WAIT_TO_START_DouplusOrderListV30FilterStatus DouplusOrderListV30FilterStatus = "WAIT_TO_START"
)

// Ptr returns reference to douplus_order_list_v3.0_filter_status value
func (v DouplusOrderListV30FilterStatus) Ptr() *DouplusOrderListV30FilterStatus {
	return &v
}
