/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthListGetV10FilteringAuthType
type QianchuanAwemeAuthListGetV10FilteringAuthType string

// List of qianchuan_aweme_auth_list_get_v1.0_filtering_auth_type
const (
	ALL_QianchuanAwemeAuthListGetV10FilteringAuthType              QianchuanAwemeAuthListGetV10FilteringAuthType = "ALL"
	AWEME_COOPERATOR_QianchuanAwemeAuthListGetV10FilteringAuthType QianchuanAwemeAuthListGetV10FilteringAuthType = "AWEME_COOPERATOR"
	OFFICIAL_QianchuanAwemeAuthListGetV10FilteringAuthType         QianchuanAwemeAuthListGetV10FilteringAuthType = "OFFICIAL"
	ROOM_QianchuanAwemeAuthListGetV10FilteringAuthType             QianchuanAwemeAuthListGetV10FilteringAuthType = "ROOM"
	SELF_QianchuanAwemeAuthListGetV10FilteringAuthType             QianchuanAwemeAuthListGetV10FilteringAuthType = "SELF"
	VIDEO_QianchuanAwemeAuthListGetV10FilteringAuthType            QianchuanAwemeAuthListGetV10FilteringAuthType = "VIDEO"
)

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_filtering_auth_type value
func (v QianchuanAwemeAuthListGetV10FilteringAuthType) Ptr() *QianchuanAwemeAuthListGetV10FilteringAuthType {
	return &v
}
