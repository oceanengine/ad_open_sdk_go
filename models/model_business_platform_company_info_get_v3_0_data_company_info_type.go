/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BusinessPlatformCompanyInfoGetV30DataCompanyInfoType
type BusinessPlatformCompanyInfoGetV30DataCompanyInfoType string

// List of business_platform_company_info_get_v3.0_data_company_info_type
const (
	BP_OTHER_BusinessPlatformCompanyInfoGetV30DataCompanyInfoType BusinessPlatformCompanyInfoGetV30DataCompanyInfoType = "BP_OTHER"
	BP_OWN_BusinessPlatformCompanyInfoGetV30DataCompanyInfoType   BusinessPlatformCompanyInfoGetV30DataCompanyInfoType = "BP_OWN"
)

// All allowed values of BusinessPlatformCompanyInfoGetV30DataCompanyInfoType enum
var AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoTypeEnumValues = []BusinessPlatformCompanyInfoGetV30DataCompanyInfoType{
	"BP_OTHER",
	"BP_OWN",
}

// NewBusinessPlatformCompanyInfoGetV30DataCompanyInfoTypeFromValue returns a pointer to a valid BusinessPlatformCompanyInfoGetV30DataCompanyInfoType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBusinessPlatformCompanyInfoGetV30DataCompanyInfoTypeFromValue(v string) (*BusinessPlatformCompanyInfoGetV30DataCompanyInfoType, error) {
	ev := BusinessPlatformCompanyInfoGetV30DataCompanyInfoType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BusinessPlatformCompanyInfoGetV30DataCompanyInfoType: valid values are %v", v, AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BusinessPlatformCompanyInfoGetV30DataCompanyInfoType) IsValid() bool {
	for _, existing := range AllowedBusinessPlatformCompanyInfoGetV30DataCompanyInfoTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to business_platform_company_info_get_v3.0_data_company_info_type value
func (v BusinessPlatformCompanyInfoGetV30DataCompanyInfoType) Ptr() *BusinessPlatformCompanyInfoGetV30DataCompanyInfoType {
	return &v
}
