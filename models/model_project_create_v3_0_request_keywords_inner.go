/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestKeywordsInner struct for ProjectCreateV30RequestKeywordsInner
type ProjectCreateV30RequestKeywordsInner struct {
	//
	Bid       *float64                          `json:"bid,omitempty"`
	BidType   *ProjectCreateV30KeywordsBidType  `json:"bid_type,omitempty"`
	MatchType ProjectCreateV30KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
