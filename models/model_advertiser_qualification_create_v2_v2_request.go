/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdvertiserQualificationCreateV2V2Request struct for AdvertiserQualificationCreateV2V2Request
type AdvertiserQualificationCreateV2V2Request struct {
	//
	AdvertiserId      int64                                              `json:"advertiser_id"`
	QualificationType AdvertiserQualificationCreateV2V2QualificationType `json:"qualification_type"`
	//
	Qualifications []*AdvertiserQualificationCreateV2V2RequestQualificationsInner `json:"qualifications"`
}
