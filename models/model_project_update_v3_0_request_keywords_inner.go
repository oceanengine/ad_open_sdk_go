/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectUpdateV30RequestKeywordsInner struct for ProjectUpdateV30RequestKeywordsInner
type ProjectUpdateV30RequestKeywordsInner struct {
	//
	Bid       *float64                          `json:"bid,omitempty"`
	BidType   *ProjectUpdateV30KeywordsBidType  `json:"bid_type,omitempty"`
	MatchType ProjectUpdateV30KeywordsMatchType `json:"match_type"`
	//
	Word string `json:"word"`
}
