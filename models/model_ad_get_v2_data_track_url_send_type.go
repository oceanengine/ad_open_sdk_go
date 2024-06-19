/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataTrackUrlSendType
type AdGetV2DataTrackUrlSendType string

// List of ad_get_v2_data_track_url_send_type
const (
	CLIENT_SEND_AdGetV2DataTrackUrlSendType AdGetV2DataTrackUrlSendType = "CLIENT_SEND"
	SERVER_SEND_AdGetV2DataTrackUrlSendType AdGetV2DataTrackUrlSendType = "SERVER_SEND"
)

// All allowed values of AdGetV2DataTrackUrlSendType enum
var AllowedAdGetV2DataTrackUrlSendTypeEnumValues = []AdGetV2DataTrackUrlSendType{
	"CLIENT_SEND",
	"SERVER_SEND",
}

// NewAdGetV2DataTrackUrlSendTypeFromValue returns a pointer to a valid AdGetV2DataTrackUrlSendType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataTrackUrlSendTypeFromValue(v string) (*AdGetV2DataTrackUrlSendType, error) {
	ev := AdGetV2DataTrackUrlSendType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataTrackUrlSendType: valid values are %v", v, AllowedAdGetV2DataTrackUrlSendTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataTrackUrlSendType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataTrackUrlSendTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_track_url_send_type value
func (v AdGetV2DataTrackUrlSendType) Ptr() *AdGetV2DataTrackUrlSendType {
	return &v
}
