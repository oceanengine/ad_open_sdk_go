/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalPromotionStatusUpdateV30DataOptStatus
type LocalPromotionStatusUpdateV30DataOptStatus string

// List of local_promotion_status_update_v3.0_data_opt_status
const (
	ENABLE_LocalPromotionStatusUpdateV30DataOptStatus LocalPromotionStatusUpdateV30DataOptStatus = "ENABLE"
	PAUSED_LocalPromotionStatusUpdateV30DataOptStatus LocalPromotionStatusUpdateV30DataOptStatus = "PAUSED"
)

// Ptr returns reference to local_promotion_status_update_v3.0_data_opt_status value
func (v LocalPromotionStatusUpdateV30DataOptStatus) Ptr() *LocalPromotionStatusUpdateV30DataOptStatus {
	return &v
}
