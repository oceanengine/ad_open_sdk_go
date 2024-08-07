/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanCampaignListGetV10DataListStatus
type QianchuanCampaignListGetV10DataListStatus string

// List of qianchuan_campaign_list_get_v1.0_data_list_status
const (
	DELETE_QianchuanCampaignListGetV10DataListStatus  QianchuanCampaignListGetV10DataListStatus = "DELETE"
	DISABLE_QianchuanCampaignListGetV10DataListStatus QianchuanCampaignListGetV10DataListStatus = "DISABLE"
	ENABLE_QianchuanCampaignListGetV10DataListStatus  QianchuanCampaignListGetV10DataListStatus = "ENABLE"
)

// All allowed values of QianchuanCampaignListGetV10DataListStatus enum
var AllowedQianchuanCampaignListGetV10DataListStatusEnumValues = []QianchuanCampaignListGetV10DataListStatus{
	"DELETE",
	"DISABLE",
	"ENABLE",
}

// NewQianchuanCampaignListGetV10DataListStatusFromValue returns a pointer to a valid QianchuanCampaignListGetV10DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanCampaignListGetV10DataListStatusFromValue(v string) (*QianchuanCampaignListGetV10DataListStatus, error) {
	ev := QianchuanCampaignListGetV10DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanCampaignListGetV10DataListStatus: valid values are %v", v, AllowedQianchuanCampaignListGetV10DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanCampaignListGetV10DataListStatus) IsValid() bool {
	for _, existing := range AllowedQianchuanCampaignListGetV10DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_campaign_list_get_v1.0_data_list_status value
func (v QianchuanCampaignListGetV10DataListStatus) Ptr() *QianchuanCampaignListGetV10DataListStatus {
	return &v
}
