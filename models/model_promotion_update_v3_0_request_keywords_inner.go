/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30RequestKeywordsInner struct for PromotionUpdateV30RequestKeywordsInner
type PromotionUpdateV30RequestKeywordsInner struct {
	//
	Bid       *float32                             `json:"bid,omitempty"`
	MatchType *PromotionUpdateV30KeywordsMatchType `json:"match_type,omitempty"`
	//
	Word string `json:"word"`
}
