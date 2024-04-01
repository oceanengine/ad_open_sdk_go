/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdvertiserGetV2DataListGender
type ReportAdvertiserGetV2DataListGender string

// List of report_advertiser_get_v2_data_list_gender
const (
	NONE_ReportAdvertiserGetV2DataListGender             ReportAdvertiserGetV2DataListGender = "NONE"
	GENDER_UNLIMITED_ReportAdvertiserGetV2DataListGender ReportAdvertiserGetV2DataListGender = "GENDER_UNLIMITED"
	GENDER_MALE_ReportAdvertiserGetV2DataListGender      ReportAdvertiserGetV2DataListGender = "GENDER_MALE"
	GENDER_FEMALE_ReportAdvertiserGetV2DataListGender    ReportAdvertiserGetV2DataListGender = "GENDER_FEMALE"
)

// All allowed values of ReportAdvertiserGetV2DataListGender enum
var AllowedReportAdvertiserGetV2DataListGenderEnumValues = []ReportAdvertiserGetV2DataListGender{
	"NONE",
	"GENDER_UNLIMITED",
	"GENDER_MALE",
	"GENDER_FEMALE",
}

// NewReportAdvertiserGetV2DataListGenderFromValue returns a pointer to a valid ReportAdvertiserGetV2DataListGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdvertiserGetV2DataListGenderFromValue(v string) (*ReportAdvertiserGetV2DataListGender, error) {
	ev := ReportAdvertiserGetV2DataListGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdvertiserGetV2DataListGender: valid values are %v", v, AllowedReportAdvertiserGetV2DataListGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdvertiserGetV2DataListGender) IsValid() bool {
	for _, existing := range AllowedReportAdvertiserGetV2DataListGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_advertiser_get_v2_data_list_gender value
func (v ReportAdvertiserGetV2DataListGender) Ptr() *ReportAdvertiserGetV2DataListGender {
	return &v
}
