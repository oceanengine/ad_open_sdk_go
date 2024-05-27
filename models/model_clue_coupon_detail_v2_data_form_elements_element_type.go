/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueCouponDetailV2DataFormElementsElementType
type ClueCouponDetailV2DataFormElementsElementType string

// List of clue_coupon_detail_v2_data_form_elements_element_type
const (
	EMAIL_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "EMAIL"
	CALCULATOR_ClueCouponDetailV2DataFormElementsElementType ClueCouponDetailV2DataFormElementsElementType = "CALCULATOR"
	TELEPHONE_ClueCouponDetailV2DataFormElementsElementType  ClueCouponDetailV2DataFormElementsElementType = "TELEPHONE"
	TEXTAREA_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "TEXTAREA"
	CHECKBOX_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "CHECKBOX"
	NUMBER_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "NUMBER"
	RADIO_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "RADIO"
	NAME_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "NAME"
	DATE_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "DATE"
	CITY_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "CITY"
	SELECT_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "SELECT"
	SEX_ClueCouponDetailV2DataFormElementsElementType        ClueCouponDetailV2DataFormElementsElementType = "SEX"
	TEXT_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "TEXT"
)

// All allowed values of ClueCouponDetailV2DataFormElementsElementType enum
var AllowedClueCouponDetailV2DataFormElementsElementTypeEnumValues = []ClueCouponDetailV2DataFormElementsElementType{
	"EMAIL",
	"CALCULATOR",
	"TELEPHONE",
	"TEXTAREA",
	"CHECKBOX",
	"NUMBER",
	"RADIO",
	"NAME",
	"DATE",
	"CITY",
	"SELECT",
	"SEX",
	"TEXT",
}

// NewClueCouponDetailV2DataFormElementsElementTypeFromValue returns a pointer to a valid ClueCouponDetailV2DataFormElementsElementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueCouponDetailV2DataFormElementsElementTypeFromValue(v string) (*ClueCouponDetailV2DataFormElementsElementType, error) {
	ev := ClueCouponDetailV2DataFormElementsElementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueCouponDetailV2DataFormElementsElementType: valid values are %v", v, AllowedClueCouponDetailV2DataFormElementsElementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueCouponDetailV2DataFormElementsElementType) IsValid() bool {
	for _, existing := range AllowedClueCouponDetailV2DataFormElementsElementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_element_type value
func (v ClueCouponDetailV2DataFormElementsElementType) Ptr() *ClueCouponDetailV2DataFormElementsElementType {
	return &v
}
