/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignUpdateStatusV2ResponseData
type CampaignUpdateStatusV2ResponseData struct {
	//
	CampaignIds []int64 `json:"campaign_ids,omitempty"`
	//
	Errors []*CampaignUpdateStatusV2ResponseDataErrorsInner `json:"errors,omitempty"`
}
