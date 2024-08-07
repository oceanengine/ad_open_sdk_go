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

// PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType
type PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType string

// List of promotion_list_v3.0_data_list_promotion_materials_anchor_material_list_anchor_type
const (
	APP_GAME_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType             PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType             PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "APP_SHOP"
	INSURANCE_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType            PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType     PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType         PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType        PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType = "SHOPPING_CART"
)

// All allowed values of PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType enum
var AllowedPromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorTypeEnumValues = []PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewPromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorTypeFromValue returns a pointer to a valid PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorTypeFromValue(v string) (*PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType, error) {
	ev := PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType: valid values are %v", v, AllowedPromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_promotion_materials_anchor_material_list_anchor_type value
func (v PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType) Ptr() *PromotionListV30DataListPromotionMaterialsAnchorMaterialListAnchorType {
	return &v
}
