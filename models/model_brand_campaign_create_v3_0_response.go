/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCampaignCreateV30Response struct for BrandCampaignCreateV30Response
type BrandCampaignCreateV30Response struct {
	//
	Code *int64                              `json:"code,omitempty"`
	Data *BrandCampaignCreateV30ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
