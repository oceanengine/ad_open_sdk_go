/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAudienceListGetV10FilteringAudienceType
type QianchuanAudienceListGetV10FilteringAudienceType string

// List of qianchuan_audience_list_get_v1.0_filtering_audience_type
const (
	BASIC_QianchuanAudienceListGetV10FilteringAudienceType  QianchuanAudienceListGetV10FilteringAudienceType = "BASIC"
	EXTEND_QianchuanAudienceListGetV10FilteringAudienceType QianchuanAudienceListGetV10FilteringAudienceType = "EXTEND"
	UPLOAD_QianchuanAudienceListGetV10FilteringAudienceType QianchuanAudienceListGetV10FilteringAudienceType = "UPLOAD"
)

// Ptr returns reference to qianchuan_audience_list_get_v1.0_filtering_audience_type value
func (v QianchuanAudienceListGetV10FilteringAudienceType) Ptr() *QianchuanAudienceListGetV10FilteringAudienceType {
	return &v
}
