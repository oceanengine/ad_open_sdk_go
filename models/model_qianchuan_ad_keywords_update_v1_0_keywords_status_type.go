/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsUpdateV10KeywordsStatusType
type QianchuanAdKeywordsUpdateV10KeywordsStatusType string

// List of qianchuan_ad_keywords_update_v1.0_keywords_status_type
const (
	DELETE_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "DELETE"
	ENABLE_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "ENABLE"
	PAUSED_QianchuanAdKeywordsUpdateV10KeywordsStatusType QianchuanAdKeywordsUpdateV10KeywordsStatusType = "PAUSED"
)

// Ptr returns reference to qianchuan_ad_keywords_update_v1.0_keywords_status_type value
func (v QianchuanAdKeywordsUpdateV10KeywordsStatusType) Ptr() *QianchuanAdKeywordsUpdateV10KeywordsStatusType {
	return &v
}
