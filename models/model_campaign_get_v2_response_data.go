/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CampaignGetV2ResponseData
type CampaignGetV2ResponseData struct {
	//
	List     []*CampaignGetV2ResponseDataListInner `json:"list,omitempty"`
	PageInfo *CampaignGetV2ResponseDataPageInfo    `json:"page_info,omitempty"`
}
