/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOmGetChallengeDispatchedProviderListV2ResponseData
type StarDemandOmGetChallengeDispatchedProviderListV2ResponseData struct {
	//
	ProviderList []*StarDemandOmGetChallengeDispatchedProviderListV2ResponseDataProviderListInner `json:"provider_list,omitempty"`
	//
	Total *int64 `json:"total,omitempty"`
}
