/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfo 附加创意
type BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfo struct {
	AdvancedCreativeType *BrandCreativeGetV30DataCreativesCreativeAdvancedCreativeInfoAdvancedCreativeType  `json:"advanced_creative_type,omitempty"`
	Card                 *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoCard     `json:"card,omitempty"`
	LiveCard             *BrandCreativeGetV30ResponseDataCreativesInnerCreativeAdvancedCreativeInfoLiveCard `json:"live_card,omitempty"`
}
