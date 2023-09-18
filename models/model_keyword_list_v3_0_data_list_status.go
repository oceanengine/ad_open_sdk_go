/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// KeywordListV30DataListStatus
type KeywordListV30DataListStatus string

// List of keyword_list_v3.0_data_list_status
const (
	AUDITING_KeywordListV30DataListStatus       KeywordListV30DataListStatus = "AUDITING"
	AUDIT_ACCEPTED_KeywordListV30DataListStatus KeywordListV30DataListStatus = "AUDIT_ACCEPTED"
	AUDIT_REJECTED_KeywordListV30DataListStatus KeywordListV30DataListStatus = "AUDIT_REJECTED"
	DELETED_KeywordListV30DataListStatus        KeywordListV30DataListStatus = "DELETED"
)

// All allowed values of KeywordListV30DataListStatus enum
var AllowedKeywordListV30DataListStatusEnumValues = []KeywordListV30DataListStatus{
	"AUDITING",
	"AUDIT_ACCEPTED",
	"AUDIT_REJECTED",
	"DELETED",
}

// NewKeywordListV30DataListStatusFromValue returns a pointer to a valid KeywordListV30DataListStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewKeywordListV30DataListStatusFromValue(v string) (*KeywordListV30DataListStatus, error) {
	ev := KeywordListV30DataListStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for KeywordListV30DataListStatus: valid values are %v", v, AllowedKeywordListV30DataListStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v KeywordListV30DataListStatus) IsValid() bool {
	for _, existing := range AllowedKeywordListV30DataListStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to keyword_list_v3.0_data_list_status value
func (v KeywordListV30DataListStatus) Ptr() *KeywordListV30DataListStatus {
	return &v
}