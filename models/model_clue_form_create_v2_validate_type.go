/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueFormCreateV2ValidateType
type ClueFormCreateV2ValidateType string

// List of clue_form_create_v2_validate_type
const (
	ALL_VERIFICATION_ClueFormCreateV2ValidateType  ClueFormCreateV2ValidateType = "ALL_VERIFICATION"
	AUTO_VERIFICATION_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "AUTO_VERIFICATION"
	CLUE_PRIORITY_ClueFormCreateV2ValidateType     ClueFormCreateV2ValidateType = "CLUE_PRIORITY"
	NONE_VERIFICATION_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "NONE_VERIFICATION"
	VALIDITY_PRIORITY_ClueFormCreateV2ValidateType ClueFormCreateV2ValidateType = "VALIDITY_PRIORITY"
)

// Ptr returns reference to clue_form_create_v2_validate_type value
func (v ClueFormCreateV2ValidateType) Ptr() *ClueFormCreateV2ValidateType {
	return &v
}
