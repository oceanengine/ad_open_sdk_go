/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormDetailV2DataFormValidateType
type ClueFormDetailV2DataFormValidateType string

// List of clue_form_detail_v2_data_form_validate_type
const (
	NONE_VERIFICATION_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "NONE_VERIFICATION"
	VALIDITY_PRIORITY_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "VALIDITY_PRIORITY"
	AUTO_VERIFICATION_ClueFormDetailV2DataFormValidateType ClueFormDetailV2DataFormValidateType = "AUTO_VERIFICATION"
	CLUE_PRIORITY_ClueFormDetailV2DataFormValidateType     ClueFormDetailV2DataFormValidateType = "CLUE_PRIORITY"
	ALL_VERIFICATION_ClueFormDetailV2DataFormValidateType  ClueFormDetailV2DataFormValidateType = "ALL_VERIFICATION"
)

// Ptr returns reference to clue_form_detail_v2_data_form_validate_type value
func (v ClueFormDetailV2DataFormValidateType) Ptr() *ClueFormDetailV2DataFormValidateType {
	return &v
}
