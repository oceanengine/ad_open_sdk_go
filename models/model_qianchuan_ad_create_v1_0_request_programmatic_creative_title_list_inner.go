/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInner struct for QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInner
type QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInner struct {
	//
	DynamicWords []*QianchuanAdCreateV10RequestProgrammaticCreativeTitleListInnerDynamicWordsInner `json:"dynamic_words,omitempty"`
	//
	Title     string                                                      `json:"title"`
	TitleType *QianchuanAdCreateV10ProgrammaticCreativeTitleListTitleType `json:"title_type,omitempty"`
}
