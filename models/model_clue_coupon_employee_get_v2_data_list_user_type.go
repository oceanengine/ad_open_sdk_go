/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeGetV2DataListUserType
type ClueCouponEmployeeGetV2DataListUserType string

// List of clue_coupon_employee_get_v2_data_list_user_type
const (
	TOUTIAO_ClueCouponEmployeeGetV2DataListUserType ClueCouponEmployeeGetV2DataListUserType = "TOUTIAO"
	DOUYIN_ClueCouponEmployeeGetV2DataListUserType  ClueCouponEmployeeGetV2DataListUserType = "DOUYIN"
)

// Ptr returns reference to clue_coupon_employee_get_v2_data_list_user_type value
func (v ClueCouponEmployeeGetV2DataListUserType) Ptr() *ClueCouponEmployeeGetV2DataListUserType {
	return &v
}
