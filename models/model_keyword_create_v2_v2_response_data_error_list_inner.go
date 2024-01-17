/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// KeywordCreateV2V2ResponseDataErrorListInner struct for KeywordCreateV2V2ResponseDataErrorListInner
type KeywordCreateV2V2ResponseDataErrorListInner struct {
	//
	Bid     *float64                               `json:"bid,omitempty"`
	BidType *KeywordCreateV2V2DataErrorListBidType `json:"bid_type,omitempty"`
	//
	ErrorReason *string                                  `json:"error_reason,omitempty"`
	MatchType   *KeywordCreateV2V2DataErrorListMatchType `json:"match_type,omitempty"`
	//
	Word string `json:"word"`
}
