/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordUpdateV30RequestKeywordsInner struct for KeywordUpdateV30RequestKeywordsInner
type KeywordUpdateV30RequestKeywordsInner struct {
	//
	Bid     *float64                         `json:"bid,omitempty"`
	BidType *KeywordUpdateV30KeywordsBidType `json:"bid_type,omitempty"`
	//
	IsPause *int64 `json:"is_pause,omitempty"`
	//
	KeywordId *int64                             `json:"keyword_id,omitempty"`
	MatchType *KeywordUpdateV30KeywordsMatchType `json:"match_type,omitempty"`
	//
	WordId *int64 `json:"word_id,omitempty"`
}
