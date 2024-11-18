/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueSmartphoneCreateV2ValidateType
type ClueSmartphoneCreateV2ValidateType string

// List of clue_smartphone_create_v2_validate_type
const (
	VALIDITY_PRIORITY_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "VALIDITY_PRIORITY"
	AUTO_VERIFICATION_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "AUTO_VERIFICATION"
	NONE_VERIFICATION_ClueSmartphoneCreateV2ValidateType ClueSmartphoneCreateV2ValidateType = "NONE_VERIFICATION"
	ALL_VERIFICATION_ClueSmartphoneCreateV2ValidateType  ClueSmartphoneCreateV2ValidateType = "ALL_VERIFICATION"
	CLUE_PRIORITY_ClueSmartphoneCreateV2ValidateType     ClueSmartphoneCreateV2ValidateType = "CLUE_PRIORITY"
)

// Ptr returns reference to clue_smartphone_create_v2_validate_type value
func (v ClueSmartphoneCreateV2ValidateType) Ptr() *ClueSmartphoneCreateV2ValidateType {
	return &v
}
