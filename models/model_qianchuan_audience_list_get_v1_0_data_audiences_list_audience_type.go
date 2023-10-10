/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// QianchuanAudienceListGetV10DataAudiencesListAudienceType
type QianchuanAudienceListGetV10DataAudiencesListAudienceType string

// List of qianchuan_audience_list_get_v1.0_data_audiences_list_audience_type
const (
	BASIC_QianchuanAudienceListGetV10DataAudiencesListAudienceType  QianchuanAudienceListGetV10DataAudiencesListAudienceType = "BASIC"
	EXTEND_QianchuanAudienceListGetV10DataAudiencesListAudienceType QianchuanAudienceListGetV10DataAudiencesListAudienceType = "EXTEND"
	UPLOAD_QianchuanAudienceListGetV10DataAudiencesListAudienceType QianchuanAudienceListGetV10DataAudiencesListAudienceType = "UPLOAD"
)

// All allowed values of QianchuanAudienceListGetV10DataAudiencesListAudienceType enum
var AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceTypeEnumValues = []QianchuanAudienceListGetV10DataAudiencesListAudienceType{
	"BASIC",
	"EXTEND",
	"UPLOAD",
}

// NewQianchuanAudienceListGetV10DataAudiencesListAudienceTypeFromValue returns a pointer to a valid QianchuanAudienceListGetV10DataAudiencesListAudienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQianchuanAudienceListGetV10DataAudiencesListAudienceTypeFromValue(v string) (*QianchuanAudienceListGetV10DataAudiencesListAudienceType, error) {
	ev := QianchuanAudienceListGetV10DataAudiencesListAudienceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QianchuanAudienceListGetV10DataAudiencesListAudienceType: valid values are %v", v, AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QianchuanAudienceListGetV10DataAudiencesListAudienceType) IsValid() bool {
	for _, existing := range AllowedQianchuanAudienceListGetV10DataAudiencesListAudienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to qianchuan_audience_list_get_v1.0_data_audiences_list_audience_type value
func (v QianchuanAudienceListGetV10DataAudiencesListAudienceType) Ptr() *QianchuanAudienceListGetV10DataAudiencesListAudienceType {
	return &v
}
