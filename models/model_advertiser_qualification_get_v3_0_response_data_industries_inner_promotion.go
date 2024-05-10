/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationGetV30ResponseDataIndustriesInnerPromotion 推广资质，相关字段见下
type AdvertiserQualificationGetV30ResponseDataIndustriesInnerPromotion struct {
	// 推广内容
	Content *string `json:"content,omitempty"`
	// 资质id
	QualificationId *int64 `json:"qualification_id,omitempty"`
	// 拒绝理由
	RejectReason *string                                                     `json:"reject_reason,omitempty"`
	Status       *AdvertiserQualificationGetV30DataIndustriesPromotionStatus `json:"status,omitempty"`
}
