/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30ResponseData
type AdvertiserQualificationGetV30ResponseData struct {
	// 广告主ID
	AdvertiserId *int64 `json:"advertiser_id,omitempty"`
	// 行业资质
	Industries        []*AdvertiserQualificationGetV30ResponseDataIndustriesInner `json:"industries,omitempty"`
	IndustryQuaStatus *AdvertiserQualificationGetV30DataIndustryQuaStatus         `json:"industry_qua_status,omitempty"`
	// 拒绝理由
	RejectReason *string                                           `json:"reject_reason,omitempty"`
	Status       *AdvertiserQualificationGetV30DataStatus          `json:"status,omitempty"`
	Subject      *AdvertiserQualificationGetV30ResponseDataSubject `json:"subject,omitempty"`
}
