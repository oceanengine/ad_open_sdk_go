/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// NativeAnchorGetV30FilteringStatus
type NativeAnchorGetV30FilteringStatus string

// List of native_anchor_get_v3.0_filtering_status
const (
	AUDIT_FAILD_NativeAnchorGetV30FilteringStatus   NativeAnchorGetV30FilteringStatus = "AUDIT_FAILD"
	AUDIT_SUCCESS_NativeAnchorGetV30FilteringStatus NativeAnchorGetV30FilteringStatus = "AUDIT_SUCCESS"
	CREATE_NativeAnchorGetV30FilteringStatus        NativeAnchorGetV30FilteringStatus = "CREATE"
	DELETE_NativeAnchorGetV30FilteringStatus        NativeAnchorGetV30FilteringStatus = "DELETE"
)

// Ptr returns reference to native_anchor_get_v3.0_filtering_status value
func (v NativeAnchorGetV30FilteringStatus) Ptr() *NativeAnchorGetV30FilteringStatus {
	return &v
}
