/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationListV30ResponseData
type AdvertiserDeliveryQualificationListV30ResponseData struct {
	//
	List     []*AdvertiserDeliveryQualificationListV30ResponseDataListInner `json:"list,omitempty"`
	PageInfo *AdvertiserDeliveryQualificationListV30ResponseDataPageInfo    `json:"page_info,omitempty"`
}
