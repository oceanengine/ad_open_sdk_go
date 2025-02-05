/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormListV2DataListFormType
type ClueFormListV2DataListFormType string

// List of clue_form_list_v2_data_list_form_type
const (
	NATIVE_FORM_ClueFormListV2DataListFormType            ClueFormListV2DataListFormType = "NATIVE_FORM"
	ADVANCED_CREATIVE_FORM_ClueFormListV2DataListFormType ClueFormListV2DataListFormType = "ADVANCED_CREATIVE_FORM"
	NORMAL_FORM_ClueFormListV2DataListFormType            ClueFormListV2DataListFormType = "NORMAL_FORM"
)

// Ptr returns reference to clue_form_list_v2_data_list_form_type value
func (v ClueFormListV2DataListFormType) Ptr() *ClueFormListV2DataListFormType {
	return &v
}
