/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanBatchCampaignStatusUpdateV10OptStatus
type QianchuanBatchCampaignStatusUpdateV10OptStatus string

// List of qianchuan_batch_campaign_status_update_v1.0_opt_status
const (
	DELETE_QianchuanBatchCampaignStatusUpdateV10OptStatus  QianchuanBatchCampaignStatusUpdateV10OptStatus = "DELETE"
	DISABLE_QianchuanBatchCampaignStatusUpdateV10OptStatus QianchuanBatchCampaignStatusUpdateV10OptStatus = "DISABLE"
	ENABLE_QianchuanBatchCampaignStatusUpdateV10OptStatus  QianchuanBatchCampaignStatusUpdateV10OptStatus = "ENABLE"
)

// All allowed values of QianchuanBatchCampaignStatusUpdateV10OptStatus enum
var AllowedQianchuanBatchCampaignStatusUpdateV10OptStatusEnumValues = []QianchuanBatchCampaignStatusUpdateV10OptStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
}

// NewQianchuanBatchCampaignStatusUpdateV10OptStatusFromValue returns a pointer to a valid QianchuanBatchCampaignStatusUpdateV10OptStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanBatchCampaignStatusUpdateV10OptStatusFromValue(v string) (*QianchuanBatchCampaignStatusUpdateV10OptStatus, error) {
	ev := QianchuanBatchCampaignStatusUpdateV10OptStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanBatchCampaignStatusUpdateV10OptStatus: valid values are %v", v, AllowedQianchuanBatchCampaignStatusUpdateV10OptStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanBatchCampaignStatusUpdateV10OptStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanBatchCampaignStatusUpdateV10OptStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_batch_campaign_status_update_v1.0_opt_status value
func (v QianchuanBatchCampaignStatusUpdateV10OptStatus) Ptr() *QianchuanBatchCampaignStatusUpdateV10OptStatus {
	return &v
}
