/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ClueFormCreateV2ElementsElementType
type ClueFormCreateV2ElementsElementType string

// List of clue_form_create_v2_elements_element_type
const (
	CITY_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "CITY"
	TEXT_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "TEXT"
	NAME_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "NAME"
	SELECT_ClueFormCreateV2ElementsElementType     ClueFormCreateV2ElementsElementType = "SELECT"
	CHECKBOX_ClueFormCreateV2ElementsElementType   ClueFormCreateV2ElementsElementType = "CHECKBOX"
	TELEPHONE_ClueFormCreateV2ElementsElementType  ClueFormCreateV2ElementsElementType = "TELEPHONE"
	RADIO_ClueFormCreateV2ElementsElementType      ClueFormCreateV2ElementsElementType = "RADIO"
	EMAIL_ClueFormCreateV2ElementsElementType      ClueFormCreateV2ElementsElementType = "EMAIL"
	CALCULATOR_ClueFormCreateV2ElementsElementType ClueFormCreateV2ElementsElementType = "CALCULATOR"
	TEXTAREA_ClueFormCreateV2ElementsElementType   ClueFormCreateV2ElementsElementType = "TEXTAREA"
	NUMBER_ClueFormCreateV2ElementsElementType     ClueFormCreateV2ElementsElementType = "NUMBER"
	SEX_ClueFormCreateV2ElementsElementType        ClueFormCreateV2ElementsElementType = "SEX"
	DATE_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "DATE"
)

// All allowed values of ClueFormCreateV2ElementsElementType enum
var AllowedClueFormCreateV2ElementsElementTypeEnumValues = []ClueFormCreateV2ElementsElementType{
	"CITY",
	"TEXT",
	"NAME",
	"SELECT",
	"CHECKBOX",
	"TELEPHONE",
	"RADIO",
	"EMAIL",
	"CALCULATOR",
	"TEXTAREA",
	"NUMBER",
	"SEX",
	"DATE",
}

// NewClueFormCreateV2ElementsElementTypeFromValue returns a pointer to a valid ClueFormCreateV2ElementsElementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClueFormCreateV2ElementsElementTypeFromValue(v string) (*ClueFormCreateV2ElementsElementType, error) {
	ev := ClueFormCreateV2ElementsElementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClueFormCreateV2ElementsElementType: valid values are %v", v, AllowedClueFormCreateV2ElementsElementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClueFormCreateV2ElementsElementType) IsValid() bool {
	for _, existing := range AllowedClueFormCreateV2ElementsElementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to clue_form_create_v2_elements_element_type value
func (v ClueFormCreateV2ElementsElementType) Ptr() *ClueFormCreateV2ElementsElementType {
	return &v
}
