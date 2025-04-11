/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponDetailV2DataFormElementsElementType
type ClueCouponDetailV2DataFormElementsElementType string

// List of clue_coupon_detail_v2_data_form_elements_element_type
const (
	SELECT_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "SELECT"
	TELEPHONE_ClueCouponDetailV2DataFormElementsElementType  ClueCouponDetailV2DataFormElementsElementType = "TELEPHONE"
	EMAIL_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "EMAIL"
	DATE_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "DATE"
	NUMBER_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "NUMBER"
	CITY_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "CITY"
	TEXTAREA_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "TEXTAREA"
	SEX_ClueCouponDetailV2DataFormElementsElementType        ClueCouponDetailV2DataFormElementsElementType = "SEX"
	CHECKBOX_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "CHECKBOX"
	RADIO_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "RADIO"
	TEXT_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "TEXT"
	NAME_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "NAME"
	CALCULATOR_ClueCouponDetailV2DataFormElementsElementType ClueCouponDetailV2DataFormElementsElementType = "CALCULATOR"
)

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_element_type value
func (v ClueCouponDetailV2DataFormElementsElementType) Ptr() *ClueCouponDetailV2DataFormElementsElementType {
	return &v
}
