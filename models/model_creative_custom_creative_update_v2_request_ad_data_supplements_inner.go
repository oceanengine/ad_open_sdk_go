/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeCustomCreativeUpdateV2RequestAdDataSupplementsInner struct for CreativeCustomCreativeUpdateV2RequestAdDataSupplementsInner
type CreativeCustomCreativeUpdateV2RequestAdDataSupplementsInner struct {
	//
	Games          []*CreativeCustomCreativeUpdateV2RequestAdDataSupplementsInnerGamesInner `json:"games,omitempty"`
	SupplementType *CreativeCustomCreativeUpdateV2AdDataSupplementsSupplementType           `json:"supplement_type,omitempty"`
}