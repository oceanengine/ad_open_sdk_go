/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner struct for AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner
type AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner struct {
	//
	AttachmentIds     []int64                                                                 `json:"attachment_ids"`
	QualificationType AdvertiserDeliveryQualificationSubmitV30QualificationsQualificationType `json:"qualification_type"`
}
