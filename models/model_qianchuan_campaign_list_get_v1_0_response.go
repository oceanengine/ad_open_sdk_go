/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanCampaignListGetV10Response struct for QianchuanCampaignListGetV10Response
type QianchuanCampaignListGetV10Response struct {
	//
	Code *int64                                   `json:"code,omitempty"`
	Data *QianchuanCampaignListGetV10ResponseData `json:"data,omitempty"`
	//
	Message *string `json:"message,omitempty"`
	//
	RequestId *string `json:"request_id,omitempty"`
}
