/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// BusinessPlatformCompanyInfoGetV30ResponseData
type BusinessPlatformCompanyInfoGetV30ResponseData struct {
	//
	CompanyInfo []*BusinessPlatformCompanyInfoGetV30ResponseDataCompanyInfoInner `json:"company_info,omitempty"`
	PageInfo    *BusinessPlatformCompanyInfoGetV30ResponseDataPageInfo           `json:"page_info,omitempty"`
}
