/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType
type QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType string

// List of qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_type
const (
	ALL_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType              QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "ALL"
	AWEME_COOPERATOR_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "AWEME_COOPERATOR"
	OFFICIAL_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType         QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "OFFICIAL"
	ROOM_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType             QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "ROOM"
	SELF_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType             QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "SELF"
	VIDEO_QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType            QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType = "VIDEO"
)

// Ptr returns reference to qianchuan_aweme_auth_list_get_v1.0_data_authorization_infos_auth_type value
func (v QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType) Ptr() *QianchuanAwemeAuthListGetV10DataAuthorizationInfosAuthType {
	return &v
}
