/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarChallengeItemsDataV2ResponseData
type StarChallengeItemsDataV2ResponseData struct {
	// 投后数据
	DataList []*StarChallengeItemsDataV2ResponseDataDataListInner `json:"data_list,omitempty"`
	PageInfo *StarChallengeItemsDataV2ResponseDataPageInfo        `json:"page_info,omitempty"`
}
