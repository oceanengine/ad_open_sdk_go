/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationDeleteV30ResponseData
type AdvertiserDeliveryQualificationDeleteV30ResponseData struct {
	// 删除失败的投放资质列表
	Errors []*AdvertiserDeliveryQualificationDeleteV30ResponseDataErrorsInner `json:"errors,omitempty"`
	// 删除成功的投放资质id
	QualificationIds []int64 `json:"qualification_ids,omitempty"`
}
