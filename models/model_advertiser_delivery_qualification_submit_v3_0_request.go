/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserDeliveryQualificationSubmitV30Request struct for AdvertiserDeliveryQualificationSubmitV30Request
type AdvertiserDeliveryQualificationSubmitV30Request struct {
	// 广告主账户ID
	AdvertiserId int64 `json:"advertiser_id"`
	// 投放资质信息
	Qualifications []*AdvertiserDeliveryQualificationSubmitV30RequestQualificationsInner `json:"qualifications"`
}
