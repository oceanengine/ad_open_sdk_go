/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationDeleteV30ResponseDataErrorsInner struct for AdvertiserDeliveryQualificationDeleteV30ResponseDataErrorsInner
type AdvertiserDeliveryQualificationDeleteV30ResponseDataErrorsInner struct {
	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`
	// 投放资质id
	QualificationId *int64 `json:"qualification_id,omitempty"`
}