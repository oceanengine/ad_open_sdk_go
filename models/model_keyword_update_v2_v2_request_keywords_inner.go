/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV2V2RequestKeywordsInner struct for KeywordUpdateV2V2RequestKeywordsInner
type KeywordUpdateV2V2RequestKeywordsInner struct {
	//
	Bid     *float64                          `json:"bid,omitempty"`
	BidType *KeywordUpdateV2V2KeywordsBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                              `json:"keyword_id,omitempty"`
	MatchType *KeywordUpdateV2V2KeywordsMatchType `json:"match_type,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
