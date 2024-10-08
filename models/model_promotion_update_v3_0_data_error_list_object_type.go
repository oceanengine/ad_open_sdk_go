/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// PromotionUpdateV30DataErrorListObjectType
type PromotionUpdateV30DataErrorListObjectType string

// List of promotion_update_v3.0_data_error_list_object_type
const (
	BUDGET_PromotionUpdateV30DataErrorListObjectType             PromotionUpdateV30DataErrorListObjectType = "BUDGET"
	KEYWORDS_PromotionUpdateV30DataErrorListObjectType           PromotionUpdateV30DataErrorListObjectType = "KEYWORDS"
	PROJECT_SETTING_PromotionUpdateV30DataErrorListObjectType    PromotionUpdateV30DataErrorListObjectType = "PROJECT_SETTING"
	PROMOTION_BUDGET_PromotionUpdateV30DataErrorListObjectType   PromotionUpdateV30DataErrorListObjectType = "PROMOTION_BUDGET"
	PROMOTION_MATERIAL_PromotionUpdateV30DataErrorListObjectType PromotionUpdateV30DataErrorListObjectType = "PROMOTION_MATERIAL"
	PROMOTION_SETTING_PromotionUpdateV30DataErrorListObjectType  PromotionUpdateV30DataErrorListObjectType = "PROMOTION_SETTING"
)

// Ptr returns reference to promotion_update_v3.0_data_error_list_object_type value
func (v PromotionUpdateV30DataErrorListObjectType) Ptr() *PromotionUpdateV30DataErrorListObjectType {
	return &v
}
