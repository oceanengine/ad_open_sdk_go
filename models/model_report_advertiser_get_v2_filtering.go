/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ReportAdvertiserGetV2Filtering
type ReportAdvertiserGetV2Filtering struct {
	//
	DeliveryMode    []*ReportAdvertiserGetV2FilteringDeliveryMode  `json:"delivery_mode,omitempty"`
	PlatformVersion *ReportAdvertiserGetV2FilteringPlatformVersion `json:"platform_version,omitempty"`
}
