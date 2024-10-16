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
	NUMBER_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "NUMBER"
	RADIO_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "RADIO"
	CITY_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "CITY"
	SELECT_ClueCouponDetailV2DataFormElementsElementType     ClueCouponDetailV2DataFormElementsElementType = "SELECT"
	DATE_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "DATE"
	CALCULATOR_ClueCouponDetailV2DataFormElementsElementType ClueCouponDetailV2DataFormElementsElementType = "CALCULATOR"
	CHECKBOX_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "CHECKBOX"
	NAME_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "NAME"
	SEX_ClueCouponDetailV2DataFormElementsElementType        ClueCouponDetailV2DataFormElementsElementType = "SEX"
	TELEPHONE_ClueCouponDetailV2DataFormElementsElementType  ClueCouponDetailV2DataFormElementsElementType = "TELEPHONE"
	EMAIL_ClueCouponDetailV2DataFormElementsElementType      ClueCouponDetailV2DataFormElementsElementType = "EMAIL"
	TEXT_ClueCouponDetailV2DataFormElementsElementType       ClueCouponDetailV2DataFormElementsElementType = "TEXT"
	TEXTAREA_ClueCouponDetailV2DataFormElementsElementType   ClueCouponDetailV2DataFormElementsElementType = "TEXTAREA"
)

// Ptr returns reference to clue_coupon_detail_v2_data_form_elements_element_type value
func (v ClueCouponDetailV2DataFormElementsElementType) Ptr() *ClueCouponDetailV2DataFormElementsElementType {
	return &v
}
