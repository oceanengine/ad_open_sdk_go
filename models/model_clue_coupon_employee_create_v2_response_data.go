/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ClueCouponEmployeeCreateV2ResponseData
type ClueCouponEmployeeCreateV2ResponseData struct {
	//
	ExistedList []*ClueCouponEmployeeCreateV2ResponseDataExistedListInner `json:"existed_list,omitempty"`
	//
	SuccessList []*ClueCouponEmployeeCreateV2ResponseDataSuccessListInner `json:"success_list,omitempty"`
}
