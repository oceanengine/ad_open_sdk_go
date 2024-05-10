/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignListGetV10FilterStatus
type QianchuanCampaignListGetV10FilterStatus string

// List of qianchuan_campaign_list_get_v1.0_filter_status
const (
	ALL_QianchuanCampaignListGetV10FilterStatus     QianchuanCampaignListGetV10FilterStatus = "ALL"
	DELETE_QianchuanCampaignListGetV10FilterStatus  QianchuanCampaignListGetV10FilterStatus = "DELETE"
	DISABLE_QianchuanCampaignListGetV10FilterStatus QianchuanCampaignListGetV10FilterStatus = "DISABLE"
	ENABLE_QianchuanCampaignListGetV10FilterStatus  QianchuanCampaignListGetV10FilterStatus = "ENABLE"
)

// All allowed values of QianchuanCampaignListGetV10FilterStatus enum
var AllowedQianchuanCampaignListGetV10FilterStatusEnumValues = []QianchuanCampaignListGetV10FilterStatus{
	"ALL",
	"DELETE",
	"DISABLE",
	"ENABLE",
}

// NewQianchuanCampaignListGetV10FilterStatusFromValue returns a pointer to a valid QianchuanCampaignListGetV10FilterStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10FilterStatusFromValue(v string) (*QianchuanCampaignListGetV10FilterStatus, error) {
	ev := QianchuanCampaignListGetV10FilterStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10FilterStatus: valid values are %v", v, AllowedQianchuanCampaignListGetV10FilterStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10FilterStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10FilterStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_filter_status value
func (v QianchuanCampaignListGetV10FilterStatus) Ptr() *QianchuanCampaignListGetV10FilterStatus {
	return &v
}
