/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarInfoV2ResponseDataInfoListInner struct for StarInfoV2ResponseDataInfoListInner
type StarInfoV2ResponseDataInfoListInner struct {
	//
	CategoryId int64 `json:"category_id"`
	//
	CompanyId int64 `json:"company_id"`
	//
	CompanyName string `json:"company_name"`
	//
	CreateTime int64                                         `json:"create_time"`
	FirstInfo  StarInfoV2ResponseDataInfoListInnerFirstInfo  `json:"first_info"`
	SecondInfo StarInfoV2ResponseDataInfoListInnerSecondInfo `json:"second_info"`
	//
	StartId int64 `json:"start_id"`
	//
	StartName string                       `json:"start_name"`
	Status    StarInfoV2DataInfoListStatus `json:"status"`
}
