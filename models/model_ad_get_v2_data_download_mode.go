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

// AdGetV2DataDownloadMode
type AdGetV2DataDownloadMode string

// List of ad_get_v2_data_download_mode
const (
	DEFAULT_AdGetV2DataDownloadMode            AdGetV2DataDownloadMode = "DEFAULT"
	APP_STORE_DELIVERY_AdGetV2DataDownloadMode AdGetV2DataDownloadMode = "APP_STORE_DELIVERY"
)

// All allowed values of AdGetV2DataDownloadMode enum
var AllowedAdGetV2DataDownloadModeEnumValues = []AdGetV2DataDownloadMode{
	"DEFAULT",
	"APP_STORE_DELIVERY",
}

// NewAdGetV2DataDownloadModeFromValue returns a pointer to a valid AdGetV2DataDownloadMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataDownloadModeFromValue(v string) (*AdGetV2DataDownloadMode, error) {
	ev := AdGetV2DataDownloadMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataDownloadMode: valid values are %v", v, AllowedAdGetV2DataDownloadModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataDownloadMode) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataDownloadModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_download_mode value
func (v AdGetV2DataDownloadMode) Ptr() *AdGetV2DataDownloadMode {
	return &v
}
