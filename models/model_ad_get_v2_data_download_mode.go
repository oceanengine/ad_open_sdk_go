/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdGetV2DataDownloadMode
type AdGetV2DataDownloadMode string

// List of ad_get_v2_data_download_mode
const (
	APP_STORE_DELIVERY_AdGetV2DataDownloadMode AdGetV2DataDownloadMode = "APP_STORE_DELIVERY"
	DEFAULT_AdGetV2DataDownloadMode            AdGetV2DataDownloadMode = "DEFAULT"
)

// Ptr returns reference to ad_get_v2_data_download_mode value
func (v AdGetV2DataDownloadMode) Ptr() *AdGetV2DataDownloadMode {
	return &v
}
