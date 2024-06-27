/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType
type PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType string

// List of promotion_update_v3.0_promotion_materials_anchor_material_list_anchor_type
const (
	APP_GAME_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType             PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_GAME"
	APP_INTERNET_SERVICE_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_INTERNET_SERVICE"
	APP_SHOP_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType             PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "APP_SHOP"
	INSURANCE_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType            PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "INSURANCE"
	ONLINE_SUBSCRIBE_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType     PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "ONLINE_SUBSCRIBE"
	PRIVATE_CHAT_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType         PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "PRIVATE_CHAT"
	SHOPPING_CART_PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType        PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType = "SHOPPING_CART"
)

// All allowed values of PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType enum
var AllowedPromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues = []PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType{
	"APP_GAME",
	"APP_INTERNET_SERVICE",
	"APP_SHOP",
	"INSURANCE",
	"ONLINE_SUBSCRIBE",
	"PRIVATE_CHAT",
	"SHOPPING_CART",
}

// NewPromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorTypeFromValue returns a pointer to a valid PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorTypeFromValue(v string) (*PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType, error) {
	ev := PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType: valid values are %v", v, AllowedPromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_promotion_materials_anchor_material_list_anchor_type value
func (v PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType) Ptr() *PromotionUpdateV30PromotionMaterialsAnchorMaterialListAnchorType {
	return &v
}
