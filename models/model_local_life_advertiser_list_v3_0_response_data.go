/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// LocalLifeAdvertiserListV30ResponseData
type LocalLifeAdvertiserListV30ResponseData struct {
	//
	AdvList  []*LocalLifeAdvertiserListV30ResponseDataAdvListInner `json:"adv_list,omitempty"`
	PageInfo *LocalLifeAdvertiserListV30ResponseDataPageInfo       `json:"page_info,omitempty"`
}
