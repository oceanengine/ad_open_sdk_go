/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// DouplusOrderListV30DataOrderListAdListAudienceDistrict
type DouplusOrderListV30DataOrderListAdListAudienceDistrict string

// List of douplus_order_list_v3.0_data_order_list_ad_list_audience_district
const (
	BUSINESS_DouplusOrderListV30DataOrderListAdListAudienceDistrict DouplusOrderListV30DataOrderListAdListAudienceDistrict = "BUSINESS"
	COUNTRY_DouplusOrderListV30DataOrderListAdListAudienceDistrict  DouplusOrderListV30DataOrderListAdListAudienceDistrict = "COUNTRY"
	PROVINCE_DouplusOrderListV30DataOrderListAdListAudienceDistrict DouplusOrderListV30DataOrderListAdListAudienceDistrict = "PROVINCE"
)

// All allowed values of DouplusOrderListV30DataOrderListAdListAudienceDistrict enum
var AllowedDouplusOrderListV30DataOrderListAdListAudienceDistrictEnumValues = []DouplusOrderListV30DataOrderListAdListAudienceDistrict{
	"BUSINESS",
	"COUNTRY",
	"PROVINCE",
}

// NewDouplusOrderListV30DataOrderListAdListAudienceDistrictFromValue returns a pointer to a valid DouplusOrderListV30DataOrderListAdListAudienceDistrict
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDouplusOrderListV30DataOrderListAdListAudienceDistrictFromValue(v string) (*DouplusOrderListV30DataOrderListAdListAudienceDistrict, error) {
	ev := DouplusOrderListV30DataOrderListAdListAudienceDistrict(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DouplusOrderListV30DataOrderListAdListAudienceDistrict: valid values are %v", v, AllowedDouplusOrderListV30DataOrderListAdListAudienceDistrictEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DouplusOrderListV30DataOrderListAdListAudienceDistrict) IsValid() bool {
	for _, existing := range AllowedDouplusOrderListV30DataOrderListAdListAudienceDistrictEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to douplus_order_list_v3.0_data_order_list_ad_list_audience_district value
func (v DouplusOrderListV30DataOrderListAdListAudienceDistrict) Ptr() *DouplusOrderListV30DataOrderListAdListAudienceDistrict {
	return &v
}
