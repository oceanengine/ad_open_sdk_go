/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdKeywordsGetV10DataListMatchType
type QianchuanAdKeywordsGetV10DataListMatchType string

// List of qianchuan_ad_keywords_get_v1.0_data_list_match_type
const (
	EXTENSIVE_QianchuanAdKeywordsGetV10DataListMatchType QianchuanAdKeywordsGetV10DataListMatchType = "EXTENSIVE"
	PHRASE_QianchuanAdKeywordsGetV10DataListMatchType    QianchuanAdKeywordsGetV10DataListMatchType = "PHRASE"
	PRECISION_QianchuanAdKeywordsGetV10DataListMatchType QianchuanAdKeywordsGetV10DataListMatchType = "PRECISION"
)

// Ptr returns reference to qianchuan_ad_keywords_get_v1.0_data_list_match_type value
func (v QianchuanAdKeywordsGetV10DataListMatchType) Ptr() *QianchuanAdKeywordsGetV10DataListMatchType {
	return &v
}
