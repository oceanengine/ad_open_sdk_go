/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreative
type QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreative struct {
	ProgrammaticCreativeCard *QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeCard `json:"programmatic_creative_card,omitempty"`
	//
	ProgrammaticCreativeMediaList []*QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeMediaListInner `json:"programmatic_creative_media_list,omitempty"`
	//
	ProgrammaticCreativeTitleList []*QianchuanAdDetailGetV10ResponseDataMultiProductCreativeListInnerProgrammaticCreativeProgrammaticCreativeTitleListInner `json:"programmatic_creative_title_list,omitempty"`
}
