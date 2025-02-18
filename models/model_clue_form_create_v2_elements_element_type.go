/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2ElementsElementType
type ClueFormCreateV2ElementsElementType string

// List of clue_form_create_v2_elements_element_type
const (
	SEX_ClueFormCreateV2ElementsElementType        ClueFormCreateV2ElementsElementType = "SEX"
	RADIO_ClueFormCreateV2ElementsElementType      ClueFormCreateV2ElementsElementType = "RADIO"
	CITY_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "CITY"
	TEXT_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "TEXT"
	NAME_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "NAME"
	TEXTAREA_ClueFormCreateV2ElementsElementType   ClueFormCreateV2ElementsElementType = "TEXTAREA"
	EMAIL_ClueFormCreateV2ElementsElementType      ClueFormCreateV2ElementsElementType = "EMAIL"
	NUMBER_ClueFormCreateV2ElementsElementType     ClueFormCreateV2ElementsElementType = "NUMBER"
	SELECT_ClueFormCreateV2ElementsElementType     ClueFormCreateV2ElementsElementType = "SELECT"
	CHECKBOX_ClueFormCreateV2ElementsElementType   ClueFormCreateV2ElementsElementType = "CHECKBOX"
	CALCULATOR_ClueFormCreateV2ElementsElementType ClueFormCreateV2ElementsElementType = "CALCULATOR"
	TELEPHONE_ClueFormCreateV2ElementsElementType  ClueFormCreateV2ElementsElementType = "TELEPHONE"
	DATE_ClueFormCreateV2ElementsElementType       ClueFormCreateV2ElementsElementType = "DATE"
)

// Ptr returns reference to clue_form_create_v2_elements_element_type value
func (v ClueFormCreateV2ElementsElementType) Ptr() *ClueFormCreateV2ElementsElementType {
	return &v
}
