/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionCreateV30AdDownloadStatus
type PromotionCreateV30AdDownloadStatus string

// List of promotion_create_v3.0_ad_download_status
const (
	OFF_PromotionCreateV30AdDownloadStatus PromotionCreateV30AdDownloadStatus = "OFF"
	ON_PromotionCreateV30AdDownloadStatus  PromotionCreateV30AdDownloadStatus = "ON"
)

// All allowed values of PromotionCreateV30AdDownloadStatus enum
var AllowedPromotionCreateV30AdDownloadStatusEnumValues = []PromotionCreateV30AdDownloadStatus{
	"OFF",
	"ON",
}

// NewPromotionCreateV30AdDownloadStatusFromValue returns a pointer to a valid PromotionCreateV30AdDownloadStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionCreateV30AdDownloadStatusFromValue(v string) (*PromotionCreateV30AdDownloadStatus, error) {
	ev := PromotionCreateV30AdDownloadStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionCreateV30AdDownloadStatus: valid values are %v", v, AllowedPromotionCreateV30AdDownloadStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionCreateV30AdDownloadStatus) IsValid() bool {
	for _, existing := range AllowedPromotionCreateV30AdDownloadStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_create_v3.0_ad_download_status value
func (v PromotionCreateV30AdDownloadStatus) Ptr() *PromotionCreateV30AdDownloadStatus {
	return &v
}
