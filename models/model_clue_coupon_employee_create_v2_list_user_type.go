/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeCreateV2ListUserType
type ClueCouponEmployeeCreateV2ListUserType string

// List of clue_coupon_employee_create_v2_list_user_type
const (
	TOUTIAO_ClueCouponEmployeeCreateV2ListUserType ClueCouponEmployeeCreateV2ListUserType = "TOUTIAO"
	DOUYIN_ClueCouponEmployeeCreateV2ListUserType  ClueCouponEmployeeCreateV2ListUserType = "DOUYIN"
)

// Ptr returns reference to clue_coupon_employee_create_v2_list_user_type value
func (v ClueCouponEmployeeCreateV2ListUserType) Ptr() *ClueCouponEmployeeCreateV2ListUserType {
	return &v
}
