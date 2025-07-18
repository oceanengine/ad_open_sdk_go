/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormListV2DataListValidateType
type ClueFormListV2DataListValidateType string

// List of clue_form_list_v2_data_list_validate_type
const (
	CLUE_PRIORITY_ClueFormListV2DataListValidateType     ClueFormListV2DataListValidateType = "CLUE_PRIORITY"
	AUTO_VERIFICATION_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "AUTO_VERIFICATION"
	NONE_VERIFICATION_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "NONE_VERIFICATION"
	ALL_VERIFICATION_ClueFormListV2DataListValidateType  ClueFormListV2DataListValidateType = "ALL_VERIFICATION"
	VALIDITY_PRIORITY_ClueFormListV2DataListValidateType ClueFormListV2DataListValidateType = "VALIDITY_PRIORITY"
)

// Ptr returns reference to clue_form_list_v2_data_list_validate_type value
func (v ClueFormListV2DataListValidateType) Ptr() *ClueFormListV2DataListValidateType {
	return &v
}
