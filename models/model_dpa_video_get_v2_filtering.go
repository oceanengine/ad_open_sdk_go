/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// DpaVideoGetV2Filtering
type DpaVideoGetV2Filtering struct {
	//
	EndTime *string `json:"end_time,omitempty"`
	//
	ImageMode *string `json:"image_mode,omitempty"`
	//
	PackageId *string `json:"package_id,omitempty"`
	//
	PackageName *string `json:"package_name,omitempty"`
	//
	ProductIds []int64 `json:"product_ids,omitempty"`
	//
	ProductPlatformIds []int64 `json:"product_platform_ids,omitempty"`
	//
	StartTime *string `json:"start_time,omitempty"`
}
