/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AudiencePackageGetV2ResponseData
type AudiencePackageGetV2ResponseData struct {
	//
	AudiencePackages []*AudiencePackageGetV2ResponseDataAudiencePackagesInner `json:"audience_packages,omitempty"`
	PageInfo         *AudiencePackageGetV2ResponseDataPageInfo                `json:"page_info,omitempty"`
}
