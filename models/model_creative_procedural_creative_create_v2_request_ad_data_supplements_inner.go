/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// CreativeProceduralCreativeCreateV2RequestAdDataSupplementsInner struct for CreativeProceduralCreativeCreateV2RequestAdDataSupplementsInner
type CreativeProceduralCreativeCreateV2RequestAdDataSupplementsInner struct {
	//
	Games          []*CreativeProceduralCreativeCreateV2RequestAdDataSupplementsInnerGamesInner `json:"games,omitempty"`
	SupplementType *CreativeProceduralCreativeCreateV2AdDataSupplementsSupplementType           `json:"supplement_type,omitempty"`
}
