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

// ReportAdGetV2DataListGender
type ReportAdGetV2DataListGender string

// List of report_ad_get_v2_data_list_gender
const (
	NONE_ReportAdGetV2DataListGender             ReportAdGetV2DataListGender = "NONE"
	GENDER_MALE_ReportAdGetV2DataListGender      ReportAdGetV2DataListGender = "GENDER_MALE"
	GENDER_UNLIMITED_ReportAdGetV2DataListGender ReportAdGetV2DataListGender = "GENDER_UNLIMITED"
	GENDER_FEMALE_ReportAdGetV2DataListGender    ReportAdGetV2DataListGender = "GENDER_FEMALE"
)

// All allowed values of ReportAdGetV2DataListGender enum
var AllowedReportAdGetV2DataListGenderEnumValues = []ReportAdGetV2DataListGender{
	"NONE",
	"GENDER_MALE",
	"GENDER_UNLIMITED",
	"GENDER_FEMALE",
}

// NewReportAdGetV2DataListGenderFromValue returns a pointer to a valid ReportAdGetV2DataListGender
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListGenderFromValue(v string) (*ReportAdGetV2DataListGender, error) {
	ev := ReportAdGetV2DataListGender(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListGender: valid values are %v", v, AllowedReportAdGetV2DataListGenderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListGender) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListGenderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_gender value
func (v ReportAdGetV2DataListGender) Ptr() *ReportAdGetV2DataListGender {
	return &v
}
