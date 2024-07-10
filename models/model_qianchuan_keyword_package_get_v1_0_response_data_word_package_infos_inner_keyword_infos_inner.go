/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInnerKeywordInfosInner struct for QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInnerKeywordInfosInner
type QianchuanKeywordPackageGetV10ResponseDataWordPackageInfosInnerKeywordInfosInner struct {
	//
	Degree *int64 `json:"degree,omitempty"`
	//
	KeywordId        *int64                                                                         `json:"keyword_id,omitempty"`
	KeywordMatchType *QianchuanKeywordPackageGetV10DataWordPackageInfosKeywordInfosKeywordMatchType `json:"keyword_match_type,omitempty"`
	//
	KeywordName string `json:"keyword_name"`
	//
	SearchSum *int64 `json:"search_sum,omitempty"`
	//
	WordPackageName *string `json:"word_package_name,omitempty"`
}
